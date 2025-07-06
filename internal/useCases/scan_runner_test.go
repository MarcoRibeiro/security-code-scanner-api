package useCases

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
)

type mockAnalyzer struct {
	name       		string
	extensions 		[]string
	matchLine  		string
}

func (m *mockAnalyzer) Name() string { return m.name }
func (m *mockAnalyzer) SupportedFileExtensions() []string { return m.extensions }
func (m *mockAnalyzer) Analyze(data string) (bool, error) {
	return data == m.matchLine, nil
}

func TestRunScan(t *testing.T) {
	dir := t.TempDir()
	file1 := filepath.Join(dir, "file1.go")
	file2 := filepath.Join(dir, "file2.go")
	file3 := filepath.Join(dir, "ignore.txt")

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
		Exclude: []string{"ignore.txt", "ignoredFolder"},
	}
	scan := &domain.Scan{
		Path: dir,
		Configuration: configuration,
	}
	analyzers := []domain.Analyzer{
		&mockAnalyzer{name: "Mock", extensions: []string{".go"}, matchLine: "matchme"},
	}

	RunScan(scan, analyzers)

	if len(scan.Findings) != 3 {
		t.Errorf("expected 3 findings, got %d", len(scan.Findings))
	}
	for _, f := range scan.Findings {
		if f.Rule != "Mock" {
			t.Errorf("unexpected rule: %s", f.Rule)
		}
		if f.Message != "matchme" {
			t.Errorf("unexpected message: %s", f.Message)
		}
		if strings.Contains(f.File, "ignoredFolder") {
			t.Errorf("should not scan files in ignoredFolder, but found: %s", f.File)
		}
	}
	if !scan.Done {
		t.Error("scan.Done should be true")
	}
	if scan.Err != "" {
		t.Errorf("unexpected scan.Err: %s", scan.Err)
	}
}
