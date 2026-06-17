# OpenRouter Go SDK Examples

This directory contains runnable examples for the OpenRouter Go SDK.

## Prerequisites

- Go 1.25 or higher
- An OpenRouter API key from [openrouter.ai/settings/keys](https://openrouter.ai/settings/keys)

## Setup

Export your API key:

```bash
export OPENROUTER_API_KEY="your-api-key"
```

## Running the chat example

```bash
cd examples/chat
go run .
```

To pin the SDK version explicitly:

```bash
go get github.com/OpenRouterTeam/go-sdk@v0.5.0
go run .
```

## Creating new examples

Add a new subdirectory under `examples/` with its own `go.mod`. Examples in this directory are not overwritten by Speakeasy generation.
