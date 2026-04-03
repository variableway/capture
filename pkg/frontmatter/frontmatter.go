package frontmatter

import (
	"bytes"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

const delimiter = "---"

// Parse splits a Markdown file into YAML frontmatter and body.
func Parse(data []byte) (frontmatter map[string]interface{}, body string, err error) {
	content := strings.TrimSpace(string(data))

	if !strings.HasPrefix(content, delimiter) {
		return nil, content, fmt.Errorf("no frontmatter delimiter found")
	}

	end := strings.Index(content[len(delimiter):], delimiter)
	if end == -1 {
		return nil, content, fmt.Errorf("no closing frontmatter delimiter found")
	}

	fmStr := content[len(delimiter) : len(delimiter)+end]
	body = strings.TrimSpace(content[len(delimiter)+end+len(delimiter):])

	frontmatter = make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(fmStr), &frontmatter); err != nil {
		return nil, content, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	return frontmatter, body, nil
}

// Encode combines a frontmatter map and body into a single Markdown byte slice.
func Encode(frontmatter map[string]interface{}, body string) ([]byte, error) {
	var buf bytes.Buffer

	fmData, err := yaml.Marshal(frontmatter)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal frontmatter: %w", err)
	}

	buf.WriteString(delimiter + "\n")
	buf.Write(fmData)
	buf.WriteString(delimiter + "\n")
	if body != "" {
		buf.WriteString("\n" + body + "\n")
	}

	return buf.Bytes(), nil
}
