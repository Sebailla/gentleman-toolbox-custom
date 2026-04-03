package injector

import (
	"strings"
	"testing"
)

func TestInjectPreservesLineComments(t *testing.T) {
	input := []byte(`{
  // This is a user comment about their setup
  "model": "gpt-4",
  "temperature": 0.7
}`)

	injector := NewJSONCInjector()
	got, err := injector.Inject(input, "model", "minimax-m2.7")
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	result := string(got)

	if !strings.Contains(result, "// This is a user comment about their setup") {
		t.Fatal("line comment was destroyed during injection")
	}

	if !strings.Contains(result, `"minimax-m2.7"`) {
		t.Fatalf("expected model to be minimax-m2.7, got:\n%s", result)
	}
}

func TestInjectPreservesBlockComments(t *testing.T) {
	input := []byte(`{
  /* Multi-line comment
     from the developer */
  "theme": "dark",
  "fontSize": 14
}`)

	injector := NewJSONCInjector()
	got, err := injector.Inject(input, "theme", "gentleman")
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	result := string(got)

	if !strings.Contains(result, "/* Multi-line comment") {
		t.Fatal("block comment was destroyed during injection")
	}

	if !strings.Contains(result, `"gentleman"`) {
		t.Fatalf("expected theme to be gentleman, got:\n%s", result)
	}
}

func TestInjectAddsNewKeyPreservingStructure(t *testing.T) {
	input := []byte(`{
  // User preferences
  "model": "gpt-4"
}`)

	injector := NewJSONCInjector()
	got, err := injector.Inject(input, "apiKey", "sk-xxx")
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	result := string(got)

	if !strings.Contains(result, "// User preferences") {
		t.Fatal("comment was destroyed when adding new key")
	}

	if !strings.Contains(result, `"apiKey"`) {
		t.Fatalf("new key was not added, got:\n%s", result)
	}

	if !strings.Contains(result, `"sk-xxx"`) {
		t.Fatalf("new value was not set, got:\n%s", result)
	}

	if !strings.Contains(result, `"gpt-4"`) {
		t.Fatal("existing value was lost when adding new key")
	}
}

func TestInjectPreservesTrailingCommas(t *testing.T) {
	input := []byte(`{
  "plugins": ["a", "b"],
  "settings": {
    "x": true,
  },
}`)

	injector := NewJSONCInjector()
	got, err := injector.Inject(input, "extra", 42)
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	result := string(got)

	if !strings.Contains(result, `"extra"`) {
		t.Fatalf("new key was not added, got:\n%s", result)
	}
}

func TestInjectNestedKeyPath(t *testing.T) {
	input := []byte(`{
  // OpenCode config
  "mcpServers": {
    "engram": {
      "command": "engram"
    }
  }
}`)

	injector := NewJSONCInjector()
	got, err := injector.Inject(input, "mcpServers.engram.command", "new-engram")
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	result := string(got)

	if !strings.Contains(result, "// OpenCode config") {
		t.Fatal("comment was destroyed during nested injection")
	}

	if !strings.Contains(result, `"new-engram"`) {
		t.Fatalf("nested value was not updated, got:\n%s", result)
	}
}

func TestInjectReturnsErrorForInvalidJSON(t *testing.T) {
	input := []byte(`{this is not json at all`)

	injector := NewJSONCInjector()
	_, err := injector.Inject(input, "key", "value")
	if err == nil {
		t.Fatal("expected error for invalid JSON input")
	}
}

func TestInjectEmptyDocument(t *testing.T) {
	input := []byte(`{}`)

	injector := NewJSONCInjector()
	got, err := injector.Inject(input, "model", "minimax-m2.7")
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	result := string(got)

	if !strings.Contains(result, `"minimax-m2.7"`) {
		t.Fatalf("expected model to be set in empty doc, got:\n%s", result)
	}
}

func TestInjectObjectValue(t *testing.T) {
	input := []byte(`{
  // My config
  "settings": {}
}`)

	injector := NewJSONCInjector()
	got, err := injector.Inject(input, "server", map[string]any{
		"host": "localhost",
		"port": 8080,
	})
	if err != nil {
		t.Fatalf("Inject() error = %v", err)
	}

	result := string(got)

	if !strings.Contains(result, "// My config") {
		t.Fatal("comment was destroyed during object injection")
	}

	if !strings.Contains(result, `"server"`) {
		t.Fatalf("object key was not added, got:\n%s", result)
	}
}
