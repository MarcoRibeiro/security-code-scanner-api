package useCases

import (
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
)

type ScanRunnerUrl struct{
	ScanRunner domain.ScanRunner
}

func NewScanRunnerUrl(ScanRunner domain.ScanRunner) *ScanRunnerUrl {
	return &ScanRunnerUrl{ScanRunner: ScanRunner}
}

func (s *ScanRunnerUrl) RunScan(scan *domain.Scan, analyzers []domain.Analyzer) {
	var tmpDir string
	
	if strings.HasPrefix(scan.Path, "http") && strings.HasSuffix(scan.Path, ".git") || strings.Contains(scan.Path,"github.com") {
		tmpDir, _ = cloneGitRepo(scan.Path)
		scan.Path = tmpDir
	}

	s.ScanRunner.RunScan(scan, analyzers)

	if tmpDir != "" {
		defer os.RemoveAll(tmpDir)
	}
}


func cloneGitRepo(url string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "repo-*")
	if err != nil {
		return "", err
	}

	_, err = git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return "", err
	}

	return tmpDir, nil
}