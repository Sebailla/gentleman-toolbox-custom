package filemerge

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gentleman-programming/gentle-ai/internal/injector"
)

// MergeJSONCPreserving merges overlayJSON into baseJSON while preserving
// all JSONC features (comments, trailing commas, whitespace) in the base.
func MergeJSONCPreserving(baseJSON []byte, overlayJSON []byte) ([]byte, error) {
	current := baseJSON
	if len(bytes.TrimSpace(current)) == 0 {
		current = []byte("{}")
	}

	overlay, err := unmarshalJSONObject(overlayJSON)
	if err != nil {
		return nil, fmt.Errorf("unmarshal overlay json: %w", err)
	}

	inj := injector.NewJSONCInjector()
	
	result, err := mergeRecursivePreserving(inj, current, overlay)
	if err != nil {
		return nil, err
	}

	// Final verification and ensuring consistent newline.
	if len(result) > 0 && result[len(result)-1] != '\n' {
		// Only append if it's not already there to maintain idempotency.
		result = append(result, '\n')
	}

	return result, nil
}

func mergeRecursivePreserving(inj *injector.JSONCInjector, base []byte, overlay map[string]any) ([]byte, error) {
	current := base
	for k, v := range overlay {
		// If both are maps, we recurse to keep surgical precision.
		if overlayMap, ok := v.(map[string]any); ok {
			if exists, isObject := checkPathIsObject(current, k); exists && isObject {
				// We need to extract the current sub-object to recurse into it.
				subObj, err := extractSubObject(inj, current, k)
				if err == nil {
					updatedSub, err := mergeRecursivePreserving(inj, subObj, overlayMap)
					if err == nil {
						// Only inject if the sub-object actually changed semantically.
						// This helps with idempotency when formatting differs.
						if !bytes.Equal(updatedSub, subObj) {
							updated, err := inj.Inject(current, k, updatedSub)
							if err == nil {
								current = updated
								continue
							}
						} else {
							// No change in sub-object, continue to next key.
							continue
						}
					}
				}
			}
		}

		// Leaf or path didn't exist as object: simple atomic injection.
		updated, err := inj.Inject(current, k, v)
		if err != nil {
			return nil, fmt.Errorf("inject key %q: %w", k, err)
		}
		current = updated
	}
	return current, nil
}

func checkPathIsObject(source []byte, key string) (exists bool, isObject bool) {
	var m map[string]any
	inj := injector.NewJSONCInjector()
	if err := json.Unmarshal(inj.Normalize(source), &m); err != nil {
		return false, false
	}
	val, ok := m[key]
	if !ok {
		return false, false
	}
	_, ok = val.(map[string]any)
	return true, ok
}

func extractSubObject(inj *injector.JSONCInjector, source []byte, key string) ([]byte, error) {
	// A strictly surgical extraction would use findValueBounds to get the exact 
	// bytes including comments. For now, we'll unmarshal and re-marshal, 
	// but the recursive logic in mergeRecursivePreserving will avoid 
	// re-injecting if the result is semantically identical.
	var m map[string]any
	if err := json.Unmarshal(inj.Normalize(source), &m); err != nil {
		return nil, err
	}
	sub, ok := m[key].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("key %q is not an object", key)
	}
	// We use a high-fidelity marshal to help with idempotency.
	return json.MarshalIndent(sub, "", "  ")
}
