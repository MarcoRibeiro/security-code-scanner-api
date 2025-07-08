package useCases

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

type mockScanRunner struct {
	called   bool
	lastScan *domain.Scan
	lastAnalyzers []domain.Analyzer
}

func (m *mockScanRunner) RunScan(scan *domain.Scan, analyzers []domain.Analyzer) {
	m.called = true
	m.lastScan = scan
	m.lastAnalyzers = analyzers
}

func TestScanRunnerUrl_RunScan_LocalPath(t *testing.T) {
	// arrange
	dir := t.TempDir()
	file := filepath.Join(dir, "file.go")
	os.WriteFile(file, []byte("test"), 0644)

	scan := &domain.Scan{Path: dir}
	analyzers := []domain.Analyzer{}
	mockRunner := &mockScanRunner{}
	scanRunnerUrl := NewScanRunnerUrl(mockRunner)

	// act
	scanRunnerUrl.RunScan(scan, analyzers)

	// assert
	assert.True(t, mockRunner.called, "expected RunScan to be called")
	assert.Equal(t, dir, scan.Path, "expected scan.Path not be updated")
}
