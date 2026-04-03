package injector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// ConfigInjector alters configuration files ensuring immutability
// in zones not affected by the injection.
type ConfigInjector interface {
	Inject(source []byte, key string, newValue any) (modified []byte, err error)
}

// JSONCInjector implements ConfigInjector for JSONC files.
type JSONCInjector struct{}

// NewJSONCInjector returns a ready-to-use JSONC injector.
func NewJSONCInjector() *JSONCInjector {
	return &JSONCInjector{}
}

// Inject sets key to newValue inside a JSONC document, preserving
// every comment and formatting detail. It treats key as a LITERAL key,
// meaning dots in the key are NOT interpreted as nesting (e.g. "chat.tools").
func (j *JSONCInjector) Inject(source []byte, key string, newValue any) ([]byte, error) {
	if err := j.validate(source); err != nil {
		return nil, err
	}

	encoded, err := j.marshalFormatted(source, newValue)
	if err != nil {
		return nil, fmt.Errorf("marshal value for key %q: %w", key, err)
	}

	// Try to find and replace the existing key.
	// This works for any key, including those with dots like "chat.tools.autoApprove".
	replaced, found := j.replaceValue(source, key, string(encoded))
	if found {
		return replaced, nil
	}

	// Key doesn't exist — append before the closing brace.
	return j.appendKey(source, key, string(encoded))
}

func (j *JSONCInjector) marshalFormatted(source []byte, value any) ([]byte, error) {
	isCompact := !bytes.Contains(source, []byte("\n"))
	if isCompact {
		return json.Marshal(value)
	}
	switch value.(type) {
	case map[string]any, []any:
		return json.MarshalIndent(value, "", "  ")
	default:
		// For strings/bools/numbers, simple marshal is fine.
		return json.Marshal(value)
	}
}

func (j *JSONCInjector) validate(source []byte) error {
	cleaned := j.Normalize(source)
	var target any
	if err := json.Unmarshal(cleaned, &target); err != nil {
		return fmt.Errorf("invalid JSONC: %w", err)
	}
	return nil
}

func (j *JSONCInjector) replaceValue(source []byte, key string, encodedValue string) ([]byte, bool) {
	text := string(source)
	pos := j.findKey(text, key)
	if pos < 0 {
		return nil, false
	}

	searchKey := fmt.Sprintf(`"%s"`, key)
	afterKey := pos + len(searchKey)
	colonIdx := -1
	for i := afterKey; i < len(text); i++ {
		if text[i] == ':' {
			colonIdx = i
			break
		}
		if text[i] == ' ' || text[i] == '\t' || text[i] == '\n' || text[i] == '\r' {
			continue
		}
		if text[i] == '/' && i+1 < len(text) && (text[i+1] == '/' || text[i+1] == '*') {
			_, end := j.findCommentBounds(text, i)
			if end > i {
				i = end - 1
				continue
			}
		}
		break
	}

	if colonIdx < 0 {
		return nil, false
	}

	valueStart, valueEnd := j.findValueBounds(text, colonIdx+1)
	if valueStart < 0 {
		return nil, false
	}

	// Semantical equality check.
	currentVal := text[valueStart:valueEnd]
	if j.isSemanticallyEqual(currentVal, encodedValue) {
		return source, true
	}

	var result bytes.Buffer
	result.Write(source[:valueStart])
	result.WriteString(encodedValue)
	result.Write(source[valueEnd:])

	return result.Bytes(), true
}

func (j *JSONCInjector) isSemanticallyEqual(a, b string) bool {
	if a == b {
		return true
	}
	var va, vb any
	if err := json.Unmarshal(j.Normalize([]byte(a)), &va); err != nil {
		return false
	}
	if err := json.Unmarshal(j.Normalize([]byte(b)), &vb); err != nil {
		return false
	}
	aj, _ := json.Marshal(va)
	bj, _ := json.Marshal(vb)
	return bytes.Equal(aj, bj)
}

func (j *JSONCInjector) appendKey(source []byte, key string, encodedValue string) ([]byte, error) {
	text := string(source)
	lastBrace := strings.LastIndex(text, "}")
	if lastBrace < 0 {
		return nil, fmt.Errorf("no closing brace found")
	}

	indent := j.detectIndent(text)
	needsComma := j.hasExistingEntries(text, lastBrace) && !j.hasTrailingComma(text, lastBrace)
	isCompact := !bytes.Contains(source, []byte("\n"))

	var insertion strings.Builder
	if needsComma {
		insertion.WriteString(",")
	}
	if !isCompact {
		insertion.WriteString("\n")
		insertion.WriteString(indent)
	} else if needsComma {
		insertion.WriteString(" ")
	}
	insertion.WriteString(fmt.Sprintf(`"%s": %s`, key, encodedValue))

	var result bytes.Buffer
	result.Write(source[:lastBrace])
	result.WriteString(insertion.String())
	result.Write(source[lastBrace:])

	return result.Bytes(), nil
}

func (j *JSONCInjector) findValueBounds(text string, afterColon int) (int, int) {
	start := afterColon
	for start < len(text) {
		c := text[start]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			start++
			continue
		}
		if c == '/' && start+1 < len(text) && (text[start+1] == '/' || text[start+1] == '*') {
			_, end := j.findCommentBounds(text, start)
			if end > start {
				start = end
				continue
			}
		}
		break
	}

	if start >= len(text) {
		return -1, -1
	}

	ch := text[start]
	switch ch {
	case '"':
		return j.findStringEnd(text, start)
	case '{':
		return j.findBracketEnd(text, start, '{', '}')
	case '[':
		return j.findBracketEnd(text, start, '[', ']')
	default:
		end := start
		for end < len(text) {
			c := text[end]
			if c == ',' || c == '}' || c == ']' || c == '\n' || c == '\r' || c == ' ' || c == '\t' {
				break
			}
			end++
		}
		return start, end
	}
}

func (j *JSONCInjector) findStringEnd(text string, start int) (int, int) {
	escaped := false
	for i := start + 1; i < len(text); i++ {
		if escaped {
			escaped = false
			continue
		}
		if text[i] == '\\' {
			escaped = true
			continue
		}
		if text[i] == '"' {
			return start, i + 1
		}
	}
	return -1, -1
}

func (j *JSONCInjector) findBracketEnd(text string, start int, open byte, close byte) (int, int) {
	depth, inStr, escaped := 0, false, false
	for i := start; i < len(text); i++ {
		ch := text[i]
		if escaped {
			escaped = false
			continue
		}
		if inStr {
			if ch == '\\' {
				escaped = true
			} else if ch == '"' {
				inStr = false
			}
			continue
		}
		if ch == '"' {
			inStr = true
			continue
		}
		if ch == open {
			depth++
		} else if ch == close {
			depth--
			if depth == 0 {
				return start, i + 1
			}
		}
	}
	return -1, -1
}

func (j *JSONCInjector) isInsideComment(text string, pos int) bool {
	inString, escaped, inLineComment, inBlockComment := false, false, false, false
	for i := 0; i < pos; i++ {
		ch := text[i]
		if inLineComment {
			if ch == '\n' {
				inLineComment = false
			}
			continue
		}
		if inBlockComment {
			if ch == '*' && i+1 < len(text) && text[i+1] == '/' {
				inBlockComment = false
				i++
			}
			continue
		}
		if inString {
			if escaped {
				escaped = false
				continue
			}
			if ch == '\\' {
				escaped = true
				continue
			}
			if ch == '"' {
				inString = false
			}
			continue
		}
		if ch == '"' {
			inString = true
			continue
		}
		if ch == '/' && i+1 < len(text) {
			if text[i+1] == '/' {
				inLineComment = true
				i++
				continue
			}
			if text[i+1] == '*' {
				inBlockComment = true
				i++
				continue
			}
		}
	}
	return inLineComment || inBlockComment
}

func (j *JSONCInjector) detectIndent(text string) string {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		trimmed := strings.TrimLeft(line, " \t")
		if len(trimmed) > 0 && trimmed[0] == '"' {
			return line[:len(line)-len(trimmed)]
		}
	}
	return "  "
}

func (j *JSONCInjector) hasExistingEntries(text string, lastBrace int) bool {
	prefix := text[:lastBrace]
	lastOpen := strings.LastIndex(prefix, "{")
	if lastOpen < 0 {
		return false
	}
	content := []byte(prefix[lastOpen+1:])
	normalized := j.Normalize(content)
	return len(bytes.TrimSpace(normalized)) > 0
}

func (j *JSONCInjector) hasTrailingComma(text string, lastBrace int) bool {
	prefix := text[:lastBrace]
	normalized := j.Normalize([]byte(prefix))
	trimmed := bytes.TrimSpace(normalized)
	// Check for trailing comma OR if it's an empty object.
	return len(trimmed) > 0 && (trimmed[len(trimmed)-1] == ',' || trimmed[len(trimmed)-1] == '{')
}

func (j *JSONCInjector) findKey(text string, key string) int {
	searchKey := fmt.Sprintf(`"%s"`, key)
	idx := 0
	for {
		pos := strings.Index(text[idx:], searchKey)
		if pos < 0 {
			return -1
		}
		pos += idx
		if j.isInsideComment(text, pos) {
			idx = pos + len(searchKey)
			continue
		}
		beforePos := -1
		for i := pos - 1; i >= 0; i-- {
			c := text[i]
			if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
				continue
			}
			if c == '{' || c == ',' || c == '[' {
				beforePos = i
				break
			}
			break
		}
		if beforePos == -1 && strings.TrimSpace(text[:pos]) != "" {
			idx = pos + 1
			continue
		}
		afterKey := pos + len(searchKey)
		foundColon := false
		for i := afterKey; i < len(text); i++ {
			c := text[i]
			if c == ':' {
				foundColon = true
				break
			}
			if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
				continue
			}
			if c == '/' && i+1 < len(text) && (text[i+1] == '/' || text[i+1] == '*') {
				_, end := j.findCommentBounds(text, i)
				if end > i {
					i = end - 1
					continue
				}
			}
			break
		}
		if foundColon {
			return pos
		}
		idx = pos + 1
	}
}

func (j *JSONCInjector) findCommentBounds(text string, start int) (int, int) {
	if start+1 >= len(text) || text[start] != '/' {
		return -1, -1
	}
	if text[start+1] == '/' {
		end := strings.Index(text[start:], "\n")
		if end < 0 {
			return start, len(text)
		}
		return start, start + end + 1
	}
	if text[start+1] == '*' {
		end := strings.Index(text[start:], "*/")
		if end < 0 {
			return start, len(text)
		}
		return start, start + end + 2
	}
	return -1, -1
}

// Normalize strips comments and trailing commas.
func (j *JSONCInjector) Normalize(source []byte) []byte {
	out := make([]byte, 0, len(source))
	inStr, escaped, inLineComment, inBlockComment := false, false, false, false
	for i := 0; i < len(source); i++ {
		ch := source[i]
		if inLineComment {
			if ch == '\n' {
				inLineComment = false
				out = append(out, ch)
			}
			continue
		}
		if inBlockComment {
			if ch == '*' && i+1 < len(source) && source[i+1] == '/' {
				inBlockComment = false
				i++
			}
			continue
		}
		if inStr {
			out = append(out, ch)
			if escaped {
				escaped = false
				continue
			}
			if ch == '\\' {
				escaped = true
				continue
			}
			if ch == '"' {
				inStr = false
			}
			continue
		}
		if ch == '"' {
			inStr = true
			out = append(out, ch)
			continue
		}
		if ch == '/' && i+1 < len(source) {
			if source[i+1] == '/' {
				inLineComment = true
				i++
				continue
			}
			if source[i+1] == '*' {
				inBlockComment = true
				i++
				continue
			}
		}
		if ch == ',' {
			nextIdx := i + 1
			found := false
			for nextIdx < len(source) {
				next := source[nextIdx]
				if next == ' ' || next == '\t' || next == '\n' || next == '\r' {
					nextIdx++
					continue
				}
				if next == '}' || next == ']' {
					found = true
				}
				break
			}
			if found {
				continue
			}
		}
		out = append(out, ch)
	}
	return out
}
