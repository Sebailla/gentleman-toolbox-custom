package injector

import (
	"encoding/json"
	"testing"
)

func TestDebugNoClosingBrace(t *testing.T) {
	inj := NewJSONCInjector()
	base := []byte("{}")
	
	// Complex nested path with wildcard
	path := "permission.read.*.env"
	
	// Pass 1: Create it
	out1, err := inj.Inject(base, path, true)
	if err != nil {
		t.Fatalf("First inject failed: %v", err)
	}
	t.Logf("Result 1:\n%s", string(out1))
	
	// Pass 2: Idempotency
	out2, err := inj.Inject(out1, path, true)
	if err != nil {
		t.Fatalf("Second inject failed: %v", err)
	}
	
	if string(out1) != string(out2) {
		t.Logf("Result 2:\n%s", string(out2))
		t.Error("Not idempotent")
	}
}

func TestDebugVSCodeInvalidJSON(t *testing.T) {
	inj := NewJSONCInjector()
	base := []byte(`{
	  // User has comments and trailing commas in VS Code settings
	  "editor.formatOnSave": true,
	  "files.exclude": {
	    "**/.git": true,
	  },
	}
`)
	
	out, err := inj.Inject(base, "chat.tools.autoApprove", true)
	if err != nil {
		t.Fatalf("Inject failed: %v", err)
	}
	t.Logf("Result:\n%s", string(out))
	
	// Verify it's valid JSONC
	if err := inj.validate(out); err != nil {
		t.Fatalf("Result is invalid JSONC: %v", err)
	}
	
	var data map[string]any
	if err := json.Unmarshal(inj.Normalize(out), &data); err != nil {
		t.Fatalf("Unmarshal failed: %v\nNormalized:\n%s", err, string(inj.Normalize(out)))
	}
}
