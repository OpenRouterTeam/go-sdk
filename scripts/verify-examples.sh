#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

MODULE="github.com/OpenRouterTeam/go-sdk"
PINNED_BUILD="${PINNED_BUILD:-auto}"
ALLOW_PENDING_RELEASE="${ALLOW_PENDING_RELEASE:-0}"

usage() {
	cat <<EOF
Usage: $(basename "$0")

Verify all examples compile and stay aligned with the SDK release version.

Environment:
  PINNED_BUILD=auto|1|0
    auto  build against the pinned module when examples match gen.lock (default)
    1     always build with GOWORK=off
    0     skip pinned-module builds
  ALLOW_PENDING_RELEASE=1
    allow examples to pin an older version than gen.lock's releaseVersion.
    For release (speakeasy-sdk-regen-*) PRs: the new version has no published
    tag yet, so example go.mod pins cannot be bumped until after the merge
    (the "Bump examples on release" job does it). Examples must still agree
    with each other and still compile against the workspace SDK.
EOF
}

read_lock_version() {
	if [[ ! -f .speakeasy/gen.lock ]]; then
		return 0
	fi

	awk '/^  releaseVersion: / { print "v" $2; exit }' .speakeasy/gen.lock
}

discover_examples() {
	local gomod dir
	shopt -s nullglob
	for gomod in examples/*/go.mod; do
		dir="$(dirname "$gomod")"
		printf '%s\n' "$dir"
	done
	shopt -u nullglob
}

example_version() {
	awk -v module="$MODULE" '$1 == "require" && $2 == module { print $3; exit }' "$1/go.mod"
}

# is_older_version A B — true when semver A < B (so A lags a pending release B).
is_older_version() {
	[[ "$1" != "$2" && "$(printf '%s\n%s\n' "$1" "$2" | sort -V | head -1)" == "$1" ]]
}

collect_examples() {
	local dir
	EXAMPLES=()
	while IFS= read -r dir; do
		[[ -n "$dir" ]] || continue
		EXAMPLES+=("$dir")
	done < <(discover_examples)
}

check_versions() {
	local expected="" lock_version="" dir version

	for dir in "${EXAMPLES[@]}"; do
		version="$(example_version "$dir")"
		if [[ -z "$version" ]]; then
			echo "missing ${MODULE} require in ${dir}/go.mod" >&2
			exit 1
		fi
		if [[ -z "$expected" ]]; then
			expected="$version"
		elif [[ "$version" != "$expected" ]]; then
			echo "version mismatch: ${dir} pins ${version}, expected ${expected}" >&2
			exit 1
		fi
	done

	lock_version="$(read_lock_version)"
	if [[ -n "$lock_version" && "$expected" != "$lock_version" ]]; then
		if [[ "$ALLOW_PENDING_RELEASE" == "1" ]] && is_older_version "$expected" "$lock_version"; then
			echo "examples pin ${expected}, gen.lock releaseVersion is ${lock_version} (pending release — allowed)"
			echo "the post-merge 'Bump examples on release' job will bump the pins"
			return 0
		fi
		echo "examples pin ${expected} but .speakeasy/gen.lock releaseVersion is ${lock_version}" >&2
		echo "run: scripts/bump-examples.sh ${lock_version}" >&2
		exit 1
	fi

	echo "all examples pin ${expected}"
	if [[ -n "$lock_version" ]]; then
		echo "matches .speakeasy/gen.lock releaseVersion (${lock_version})"
	fi
}

build_examples() {
	local mode="$1"
	local dir

	for dir in "${EXAMPLES[@]}"; do
		echo "building ${dir} (${mode})"
		(
			cd "$dir"
			if [[ "$mode" == "pinned module" ]]; then
				GOWORK=off go build -o /dev/null .
				GOWORK=off go vet ./...
			else
				go build -o /dev/null .
				go vet ./...
			fi
		)
	done
}

should_run_pinned_build() {
	local expected lock_version

	case "$PINNED_BUILD" in
	0) return 1 ;;
	1) return 0 ;;
	auto)
		lock_version="$(read_lock_version)"
		[[ -z "$lock_version" ]] && return 0
		expected="$(example_version "${EXAMPLES[0]}")"
		[[ "$expected" == "$lock_version" ]]
		;;
	*)
		echo "invalid PINNED_BUILD=${PINNED_BUILD}" >&2
		exit 1
		;;
	esac
}

main() {
	if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
		usage
		exit 0
	fi

	collect_examples
	if ((${#EXAMPLES[@]} == 0)); then
		echo "no examples found under examples/" >&2
		exit 1
	fi

	check_versions
	build_examples "workspace SDK"

	if should_run_pinned_build; then
		build_examples "pinned module"
	else
		echo "skipping pinned-module build (examples are behind gen.lock or PINNED_BUILD=0)"
	fi

	echo "examples verified"
}

main "$@"
