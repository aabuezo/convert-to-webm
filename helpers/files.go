package helpers

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func splitPath(path string) (dir string, suffix string, ext string) {
	// check if file exists
	if _, err := os.Stat(path); err != nil {
		log.Fatalf("Error: file %v does not exist.", path)
		os.Exit(1)
	}

	dir, filename := filepath.Split(path)
	ext = filepath.Ext(filename)
	suffix = strings.TrimSuffix(filename, ext)
	// log.Println("dir:", dir)
	// log.Println("filename:", filename)
	// log.Println("suffix:", suffix)
	// log.Println("ext:", ext)

	return
}
