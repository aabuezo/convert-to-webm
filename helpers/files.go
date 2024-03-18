package helpers

import (
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
