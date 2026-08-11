#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

MODULE="github.com/OpenRouterTeam/go-sdk"
DOCS=(
	examples/README.md
	README.md
	OVERVIEW.md
	doc.go
)

usage() {
	cat <<EOF >&2
Usage: $(basename "$0") [--docs-only] [vX.Y.Z]

Bump pinned SDK versions in all examples/ modules and docs.

If VERSION is omitted, reads releaseVersion from .speakeasy/gen.lock.

Options:
  --docs-only      Only rewrite version strings in docs (README.md, OVERVIEW.md,
                   doc.go, examples/README.md, example_test.go). Skips the
                   examples/*/go.mod 'go get' bump, so it needs no network and
                   no published module — safe to run on the regen branch before
                   the release tag exists.

Environment:
  GO_GET_RETRIES   attempts to fetch the module from the proxy (default 20)
  GO_GET_INTERVAL  seconds to wait between attempts (default 30)
EOF
}

normalize_version() {
	local version="$1"
	if [[ "$version" != v* ]]; then
		version="v${version}"
	fi
	printf '%s' "$version"
}

read_lock_version() {
	if [[ ! -f .speakeasy/gen.lock ]]; then
		echo "missing .speakeasy/gen.lock; pass an explicit version" >&2
		exit 1
	fi

	local version
	version="$(awk '/^  releaseVersion: / { print $2; exit }' .speakeasy/gen.lock)"
	if [[ -z "$version" ]]; then
		echo "could not read releaseVersion from .speakeasy/gen.lock" >&2
		exit 1
	fi

	normalize_version "$version"
}

update_docs() {
	local version="$1"
	local doc

	# Pass the version through the environment rather than interpolating it into
	# the perl program text: the module path contains '@v', and a bare '@v' in a
	# double-quoted perl -e string is parsed as array interpolation (expands to
	# empty), so the pattern would never match. Escaping as '\@' + $ENV avoids it.
	for doc in "${DOCS[@]}"; do
		[[ -f "$doc" ]] || continue
		BUMP_VERSION="$version" perl -pi -e '
			my $v = $ENV{BUMP_VERSION};
			s#github\.com/OpenRouterTeam/go-sdk\@v[0-9]+(?:\.[0-9]+)*#github.com/OpenRouterTeam/go-sdk\@$v#g;
			s#(require github\.com/OpenRouterTeam/go-sdk )v[0-9]+(?:\.[0-9]+)*#$1$v#g;
		' "$doc"
	done

	# example_test.go asserts the SDK version via `// Output: X.Y.Z` (no `v`
	# prefix, no module path), so it needs a dedicated substitution.
	local bare="${version#v}"
	if [[ -f example_test.go ]]; then
		BUMP_BARE="$bare" perl -pi -e '
			my $b = $ENV{BUMP_BARE};
			s#(// Output: )[0-9]+(?:\.[0-9]+)*#$1$b#g;
		' example_test.go
	fi
}

# GOSUMDB=off: the checksum database (sum.golang.org) can lag the module proxy
# by 10+ minutes after a tag publishes, but `go get` computes and records the
# correct hash into go.sum directly from the proxy/tag without it. Waiting on
# sum.golang.org is the reason release-time bumps used to fail; bypass it here
# (downstream consumers still get sumdb-verified hashes once it backfills).
GO_GET_RETRIES="${GO_GET_RETRIES:-20}"
GO_GET_INTERVAL="${GO_GET_INTERVAL:-30}"

go_get_with_retry() {
	local dir="$1"
	local version="$2"
	local attempt

	for attempt in $(seq 1 "$GO_GET_RETRIES"); do
		if (cd "$dir" && GOWORK=off GOSUMDB=off go get "${MODULE}@${version}"); then
			return 0
		fi
		if (( attempt == GO_GET_RETRIES )); then
			echo "failed to fetch ${MODULE}@${version} for ${dir} after ${GO_GET_RETRIES} attempts" >&2
			return 1
		fi
		echo "waiting for ${MODULE}@${version} to appear on the module proxy (attempt ${attempt}/${GO_GET_RETRIES})..." >&2
		sleep "$GO_GET_INTERVAL"
	done
}

DOCS_ONLY=0
while [[ $# -gt 0 ]]; do
	case "$1" in
	-h | --help)
		usage
		exit 0
		;;
	--docs-only)
		DOCS_ONLY=1
		shift
		;;
	-*)
		echo "unknown option: $1" >&2
		usage
		exit 1
		;;
	*)
		break
		;;
	esac
done

VERSION="${1:-}"
if [[ -z "$VERSION" ]]; then
	VERSION="$(read_lock_version)"
else
	VERSION="$(normalize_version "$VERSION")"
fi

# Docs-only: rewrite version strings without touching example modules. No
# network and no published tag required, so this can run on the regen branch
# before the release is cut — keeping each release's pkg.go.dev README in sync.
if (( DOCS_ONLY )); then
	echo "bumping docs to ${MODULE}@${VERSION} (docs-only)"
	update_docs "$VERSION"
	echo "done."
	exit 0
fi

echo "bumping examples to ${MODULE}@${VERSION}"

shopt -s nullglob
example_mods=(examples/*/go.mod)
shopt -u nullglob

if ((${#example_mods[@]} == 0)); then
	echo "no examples found under examples/" >&2
	exit 1
fi

for gomod in "${example_mods[@]}"; do
	dir="$(dirname "$gomod")"
	echo "  ${dir}"
	go_get_with_retry "$dir" "$VERSION"
	(
		cd "$dir"
		GOWORK=off GOSUMDB=off go mod tidy
	)
done

update_docs "$VERSION"

if [[ -f go.work ]]; then
	go work sync
fi

echo "done. run scripts/verify-examples.sh to confirm."
