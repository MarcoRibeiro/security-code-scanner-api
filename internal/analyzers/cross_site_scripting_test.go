package analyzers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrossSiteScriptingAnalyzer_SupportedFileExtensions(t *testing.T) {
	analyzer := NewCrossSiteScriptingAnalyzer()

	expectedExtensions := []string{".html", ".js"}
	if len(analyzer.SupportedFileExtensions()) != len(expectedExtensions) {
		t.Errorf("Expected %d supported file extensions, got %d", len(expectedExtensions), len(analyzer.SupportedFileExtensions()))
	}

	for i, ext := range expectedExtensions {
		if analyzer.SupportedFileExtensions()[i] != ext {
			t.Errorf("Expected supported file extension '%s', got '%s'", ext, analyzer.SupportedFileExtensions()[i])
		}
	}
}

func TestCrossSiteScriptingAnalyzer_Analyze(t *testing.T) {
	analyzer := NewCrossSiteScriptingAnalyzer()

	tests := []struct {
		name     string
		line     string
		expected     bool
	}{
		{
			name:     "Detects XSS with script tag",
			line:     `<script>alert('XSS');</script>`,
			expected:     true,
		},
		{
			name:     "Detects alert statement",
			line:     `Alert()`,
			expected:     true,
		},
		{
			name:     "Detects XSS with img tag",
			line:     `<img src="x" onerror="alert('XSS');">`,
			expected:     true,
		},
		{
			name:     "Ignores safe content",
			line:     `This is a safe string.`,
			expected:     false,
		},
	}

	for _, test := range tests {
		match := analyzer.Analyze(test.line)
		assert.Equal(t, test.expected, match, test.name)
	}
}