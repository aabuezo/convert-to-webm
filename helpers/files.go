package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func splitPath(path string) (dir string, suffix string, ext string) {
	// splits the command arguments into dir, filename and extension
	dir, filename := filepath.Split(path)
	ext = filepath.Ext(filename)
	suffix = strings.TrimSuffix(filename, ext)
	return
}

func isFileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}
