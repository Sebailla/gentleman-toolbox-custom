package filemerge

import (
	"strings"
	"testing"
)

func TestMergeJSONCPreservingKeepsLineComments(t *testing.T) {
	base := []byte(`{
  // User's personal model preference
  "model": "gpt-4",
  "temperature": 0.7
}`)
	overlay := []byte(`{"model": "minimax-m2.7"}`)

	merged, err := MergeJSONCPreserving(base, overlay)
	if err != nil {
		t.Fatalf("MergeJSONCPreserving() error = %v", err)
	}

	result := string(merged)

	if !strings.Contains(result, "// User's personal model preference") {
		t.Fatal("line comment was destroyed during merge")
	}

	if !strings.Contains(result, `"minimax-m2.7"`) {
		t.Fatalf("overlay value not applied, got:\n%s", result)
	}
}

func TestMergeJSONCPreservingKeepsBlockComments(t *testing.T) {
	base := []byte(`{
  /* OpenCode settings
     DO NOT EDIT MANUALLY */
  "theme": "dark"
}`)
	overlay := []byte(`{"theme": "gentleman"}`)

	merged, err := MergeJSONCPreserving(base, overlay)
	if err != nil {
		t.Fatalf("MergeJSONCPreserving() error = %v", err)
	}

	result := string(merged)

	if !strings.Contains(result, "/* OpenCode settings") {
		t.Fatal("block comment was destroyed")
	}
}

func TestMergeJSONCPreservingAddsNewKeys(t *testing.T) {
	base := []byte(`{
  // My config
  "model": "gpt-4"
}`)
	overlay := []byte(`{"mcpServers": {"context7": {"command": "npx"}}}`)

	merged, err := MergeJSONCPreserving(base, overlay)
	if err != nil {
		t.Fatalf("MergeJSONCPreserving() error = %v", err)
	}

	result := string(merged)

	if !strings.Contains(result, "// My config") {
		t.Fatal("comment lost during new key addition")
	}

	if !strings.Contains(result, `"mcpServers"`) {
		t.Fatalf("new key not added, got:\n%s", result)
	}

	if !strings.Contains(result, `"gpt-4"`) {
		t.Fatal("existing value lost")
	}
}

func TestMergeJSONCPreservingHandlesEmptyBase(t *testing.T) {
	base := []byte(``)
	overlay := []byte(`{"model": "minimax-m2.7"}`)

	merged, err := MergeJSONCPreserving(base, overlay)
	if err != nil {
		t.Fatalf("MergeJSONCPreserving() error = %v", err)
	}

	result := string(merged)
	if !strings.Contains(result, `"minimax-m2.7"`) {
		t.Fatalf("value not set on empty base, got:\n%s", result)
	}
}

func TestMergeJSONCPreservingHandlesNilBase(t *testing.T) {
	overlay := []byte(`{"model": "minimax-m2.7"}`)

	merged, err := MergeJSONCPreserving(nil, overlay)
	if err != nil {
		t.Fatalf("MergeJSONCPreserving() error = %v", err)
	}

	result := string(merged)
	if !strings.Contains(result, `"minimax-m2.7"`) {
		t.Fatalf("value not set on nil base, got:\n%s", result)
	}
}
