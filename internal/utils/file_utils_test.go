package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalkExcludingFilesAndDirs(t *testing.T) {
	
	// arrange
	dir := t.TempDir()
	os.Mkdir(filepath.Join(dir, "sub"), 0755)
	os.Mkdir(filepath.Join(dir, "excludeDir"), 0755)
	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("a"), 0644)
	os.WriteFile(filepath.Join(dir, "file2.txt"), []byte("b"), 0644)
	os.WriteFile(filepath.Join(dir, "file2.html"), []byte("b"), 0644)
	os.WriteFile(filepath.Join(dir, "excludeMe.txt"), []byte("c"), 0644)
	os.WriteFile(filepath.Join(dir, "sub", "file3.txt"), []byte("d"), 0644)
	os.WriteFile(filepath.Join(dir, "excludeDir", "file4.txt"), []byte("e"), 0644)

	excludes := []string{"excludeMe.txt", "excludeDir", "*.html"}
	var visited []string

	// act
	WalkExcludingFilesAndDirs(dir, excludes, func(path string) error {
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			visited = append(visited, filepath.Base(path))
		}
		return nil
	})

	// assert
	assert.ElementsMatch(t, visited, []string{"file1.txt", "file2.txt", "file3.txt"})
}

func TestReadFileByLine(t *testing.T) {
	// arrange
	file := t.TempDir() + "/test.txt"
	os.WriteFile(file, []byte("line1\nline2\nline3\n"), 0644)
	var lines []string
	
	// act
	ReadFileByLine(file, func(line string, lineNum int) error {
		lines = append(lines, line)
		return nil
	})

	assert.Equal(t, lines, []string{"line1", "line2", "line3"})
}