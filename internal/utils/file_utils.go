package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func WalkExcludingFilesAndDirs(dir string, excludes []string, fn func(string) error) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if shouldExclude(path, excludes) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		return fn(path)
	})
}

func ReadFileByLine(path string, fn func(line string, lineNum int) error) error {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		lineNum := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			fn(line, lineNum)
		}

		return nil
}

func shouldExclude(path string, excludes []string) bool {
	for _, pattern := range excludes {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}