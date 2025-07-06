package util

import (
	"os"
	"path/filepath"
)

func Walk(dir string, ignore []string, fn func(string) error) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			for _, pattern := range ignore {
				if matched, _ := filepath.Match(pattern, info.Name()); matched {
					return filepath.SkipDir
				}
			}
		}
		return fn(path)
	})
}

func IsIgnored(path string, ignore []string) bool {
	for _, pattern := range ignore {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}