package useCases

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
	"github.com/marrcoribeiro/security-scanner-api/internal/utils"
)

type ScanRunner struct{}

func NewScanRunner() *ScanRunner {
	return &ScanRunner{}
}

func (s *ScanRunner) RunScan(scan *domain.Scan, analyzers []domain.Analyzer) {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	excludes := []string{}

	if condition := scan.Configuration != nil && scan.Configuration.Exclude != nil; condition {
		excludes = scan.Configuration.Exclude
	}

	err := utils.WalkExcludingFilesAndDirs(scan.Path, excludes, func(path string) error {
		fileExt := strings.ToLower(filepath.Ext(path))
		supportedAnalyzers := filterSupportedAnalyzers(analyzers, fileExt)

		if len(supportedAnalyzers) == 0 {
			return nil
		}
		
		wg.Add(1)
		go func() {
			defer wg.Done()
			utils.ReadFileByLine(path, func(line string, lineNum int) error {
				runAnalyzers(scan, supportedAnalyzers, path, line, lineNum, &mu)
				return nil
			})
		}()
		
		return nil
	})

	wg.Wait()
	if err != nil {
		scan.Err = fmt.Sprintf("Error during scan: %v", err)
		scan.Done = false
		return
	}

	scan.Done = true
}

func runAnalyzers(scan *domain.Scan, analyzers []domain.Analyzer, path string, line string, lineNum int, mu *sync.Mutex) {
	for _, analyzer := range analyzers {
		match := analyzer.Analyze(line)
		
		if match {
			finding := domain.Finding{
				Rule:    analyzer.Name(),
				File:    path,
				Message: line,
				Line:    lineNum,
			}
			mu.Lock()
			scan.Findings = append(scan.Findings, finding)
			mu.Unlock()
		}
	}
}

func filterSupportedAnalyzers(analyzers []domain.Analyzer, fileExt string) []domain.Analyzer {
	supported := []domain.Analyzer{}
	for _, analyzer := range analyzers {
		if slices.Contains(analyzer.SupportedFileExtensions(), fileExt) {
			supported = append(supported, analyzer)
		}
	}
	return supported
}