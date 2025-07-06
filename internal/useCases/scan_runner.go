package useCases

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
)

func RunScan(scan *domain.Scan, analyzers []domain.Analyzer) {
	excludeMap := make(map[string]struct{})
	for _, ex := range scan.Configuration.Exclude {
		excludeMap[ex] = struct{}{}
	}

	filepath.Walk(scan.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			scan.Err = fmt.Sprintf("Error accessing path %s: %v", path, err)
			return err
		}

		base := filepath.Base(path)
		if _, found := excludeMap[base]; found {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			scan.Err = fmt.Sprintf("Error opening file %s: %v", path, err)
			return nil
		}
		defer file.Close()

		fileExt := strings.ToLower(filepath.Ext(path))
		lineNum := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			for _, analyzer := range analyzers {
				supported := slices.Contains(analyzer.SupportedFileExtensions(), fileExt)
				if !supported {
					continue
				}
				match, err := analyzer.Analyze(line)
				if err != nil {
					continue
				}
				if match {
					finding := domain.Finding{
						Rule:    analyzer.Name(),
						File:    path,
						Message: line,
						Line:    lineNum,
					}
					scan.Findings = append(scan.Findings, finding)
				}
			}
		}
		return nil
	})

	scan.Done = true
}