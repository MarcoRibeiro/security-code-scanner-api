package useCases

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
)

type benchAnalyzer struct{}

func (b *benchAnalyzer) Name() string { return "BenchAnalyzer" }
func (b *benchAnalyzer) SupportedFileExtensions() []string { return []string{".go"} }
func (b *benchAnalyzer) Analyze(line string) bool { return len(line) > 0 }

func BenchmarkScanRunner_RunScan(b *testing.B) {
	dir := b.TempDir()
	// Create 100 files with 100 lines each
	for i := 0; i < 100; i++ {
		file := filepath.Join(dir, "file"+string(rune(i))+".go")
		f, _ := os.Create(file)
		for j := 0; j < 100; j++ {
			f.WriteString("line content\n")
		}
		f.Close()
	}

	analyzers := []domain.Analyzer{&benchAnalyzer{}}
	scanRunner := NewScanRunner()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scan := &domain.Scan{Path: dir}
		scanRunner.RunScan(scan, analyzers)
	}
}
