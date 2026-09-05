package components_test

import (
	"bytes"
	"encoding/json"
	"testing"

	openrouter "github.com/OpenRouterTeam/go-sdk"
	"github.com/OpenRouterTeam/go-sdk/models/components"
)

func TestResponsesRequestMarshalWithoutServiceTier(t *testing.T) {
	req := components.ResponsesRequest{
		Input: openrouter.Pointer(
			components.CreateInputsUnionStr("Hey there"),
		),
		Model:  openrouter.Pointer("openai/gpt-4o-mini"),
		Stream: openrouter.Bool(false),
	}

	b, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	if !json.Valid(b) {
		t.Fatalf("invalid JSON: %s", string(b))
	}

	if !bytes.Contains(b, []byte(`"service_tier":"auto"`)) {
		t.Fatalf("expected service_tier default to be marshaled as string auto, got: %s", string(b))
	}
}
