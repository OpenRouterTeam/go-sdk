# OpenRouter Go SDK Examples

Runnable examples for the OpenRouter Go SDK. Each subdirectory is a self-contained module with its own `go.mod`.

## Prerequisites

- Go 1.25 or higher
- An OpenRouter API key from [openrouter.ai/settings/keys](https://openrouter.ai/settings/keys)

## Setup

```bash
export OPENROUTER_API_KEY="your-api-key"
```

## Examples

| Directory | SDK client | Run |
|-----------|------------|-----|
| [`chat`](chat) | `client.Chat.Send` | `cd examples/chat && go run .` |
| [`chat-stream`](chat-stream) | streaming chat | `cd examples/chat-stream && go run .` |
| [`models`](models) | `client.Models.List` | `cd examples/models && go run .` |
| [`model`](model) | `client.Models.Get` | `cd examples/model && go run .` |
| [`embed`](embed) | `client.Embeddings.Generate` | `cd examples/embed && go run .` |
| [`rerank`](rerank) | `client.Rerank.Rerank` | `cd examples/rerank && go run .` |
| [`providers`](providers) | `client.Providers.List` | `cd examples/providers && go run .` |
| [`generation`](generation) | `client.Generations.GetGeneration` | see below |

Start with [`chat`](chat) if you are new to the SDK.

### Generation example

The generation example looks up metadata for a prior request. Run chat first, then pass the completion id:

```bash
cd examples/chat && go run .
export OPENROUTER_GENERATION_ID="gen-..."
cd ../generation && go run .
```

## SDK version

Each example pins a released SDK version in `go.mod`:

```go
require github.com/OpenRouterTeam/go-sdk v0.5.38
```

This should match the version in [README.md](../README.md) and `.speakeasy/gen.lock` `releaseVersion`. CI runs `scripts/bump-examples.sh` after releases to keep these in sync.

To use a different version:

```bash
go get github.com/OpenRouterTeam/go-sdk@v0.5.38
go run .
```

Do not add a `replace` directive when copying an example into your own project.

## Local SDK development

When working inside this repository, the root [`go.work`](../go.work) file wires examples to the local SDK checkout automatically. No per-example `replace` directives are needed.

If you clone the repo and examples fail to resolve the SDK module, run commands from anywhere inside the repository so Go picks up `go.work`.

## Creating new examples

Add a new subdirectory under `examples/` with its own `go.mod`. Examples in this directory are not overwritten by Speakeasy generation.

New examples should:

1. `require github.com/OpenRouterTeam/go-sdk` at a pinned release version
2. Be listed in the root `go.work` file for local development

## Release automation

When Speakeasy bumps `.speakeasy/gen.lock`, CI runs `scripts/bump-examples.sh` on `main` to sync example pins and docs with `releaseVersion`.

Locally:

```bash
scripts/bump-examples.sh        # reads releaseVersion from .speakeasy/gen.lock
scripts/bump-examples.sh v0.6.0 # or pass an explicit version
scripts/verify-examples.sh      # compile + version alignment checks
```

Pull requests that change SDK code or examples must pass `scripts/verify-examples.sh` in CI.
