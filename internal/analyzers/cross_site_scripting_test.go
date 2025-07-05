package analyzers

import (
	"testing"
)

func TestCrossSiteScriptingAnalyzer(t *testing.T) {
	analyzer := NewCrossSiteScriptingAnalyzer()

	// Test supported file extensions
	expectedExtensions := []string{".html", ".js"}
	if len(analyzer.SupportedFileExtensions()) != len(expectedExtensions) {
		t.Errorf("Expected %d supported file extensions, got %d", len(expectedExtensions), len(analyzer.SupportedFileExtensions()))
	}

	for i, ext := range expectedExtensions {
		if analyzer.SupportedFileExtensions()[i] != ext {
			t.Errorf("Expected supported file extension '%s', got '%s'", ext, analyzer.SupportedFileExtensions()[i])
		}
	}

	// Test analysis
	data := `<script>alert('XSS');</script>`
	matches, err := analyzer.Analyze(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !matches {
		t.Error("Expected to find a cross-site scripting vulnerability")
	}
}


func TestCrossSiteScriptingAnalyzer_Analyze(t *testing.T) {
	analyzer := NewCrossSiteScriptingAnalyzer()

	tests := []struct {
		name     string
		line     string
		want     bool
	}{
		{
			name:     "Detects XSS with script tag",
			line:     `<script>alert('XSS');</script>`,
			want:     true,
		},
		{
			name:     "Detects alert statement",
			line:     `Alert()`,
			want:     true,
		},
		{
			name:     "Detects XSS with img tag",
			line:     `<img src="x" onerror="alert('XSS');">`,
			want:     true,
		},
		{
			name:     "Ignores safe content",
			line:     `This is a safe string.`,
			want:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			found, err := analyzer.Analyze(test.line)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if found != test.want {
				t.Errorf("Expected %v, got %v", test.want, found)
			}
		})
	}
}
