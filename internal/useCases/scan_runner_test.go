package useCases

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

type mockAnalyzer struct {
	name       		string
	extensions 		[]string
	matchLine  		string
}

func (m *mockAnalyzer) Name() string { return m.name }
func (m *mockAnalyzer) SupportedFileExtensions() []string { return m.extensions }
func (m *mockAnalyzer) Analyze(data string) bool {
	return data == m.matchLine
}

func TestRunScan(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	file1 := filepath.Join(dir, "file1.go")
	file2 := filepath.Join(dir, "file2.go")
	file3 := filepath.Join(dir, "ignore.go")

	subdir := filepath.Join(dir, "subfolder")
	os.Mkdir(subdir, 0755)
	subfile := filepath.Join(subdir, "subfile.go")
	os.WriteFile(file1, []byte("safe\nmatchme\n"), 0644)
	os.WriteFile(file2, []byte("matchme\n"), 0644)
	os.WriteFile(file3, []byte("matchme\n"), 0644)
	os.WriteFile(subfile, []byte("matchme\n"), 0644)

	ignoredDir := filepath.Join(dir, "ignoredFolder")
	os.Mkdir(ignoredDir, 0755)
	ignoredFile := filepath.Join(ignoredDir, "shouldnotfind.go")
	os.WriteFile(ignoredFile, []byte("matchme\n"), 0644)

	configuration := &domain.Configuration{
		Exclude: []string{"ignore.go", "ignoredFolder"},
	}
	scan := &domain.Scan{
		Path: dir,
		Configuration: configuration,
	}
	analyzers := []domain.Analyzer{
		&mockAnalyzer{name: "really dangerous thing", extensions: []string{".go"}, matchLine: "matchme"},
	}

	scanner := NewScanRunner()

	// Act
	scanner.RunScan(scan, analyzers)

	// Assert
	assert.Equal(t, 3, len(scan.Findings), "expected 3 findings")
	assert.ElementsMatch(t, []domain.Finding{
		{Rule: "really dangerous thing", File: file1, Message: "matchme", Line: 2},
		{Rule: "really dangerous thing", File: file2, Message: "matchme", Line: 1},
		{Rule: "really dangerous thing", File: subfile, Message: "matchme", Line: 1},
	}, scan.Findings, "findings should match expected results")
	assert.Equal(t, true, scan.Done, "scan should be marked as done")
	assert.Empty(t, scan.Err, "scan should not have errors")
}
