package main

import (
	"log"
	"os"
	"path/filepath"
)

func walkpath(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		log.Printf("WARNING: %v\n", err)
		return nil
	}
	log.Printf("DEBUG path: %s\n", path)
	return err
}

func getDirectoryListing(pathname string) error {
	dir, err := os.Open(pathname)
	if err == nil {
		defer dir.Close()
		filepath.Walk(pathname, walkpath)
	}
	log.Printf("WARNING: %v\n", err)

	return err
}

func main() {
	rootDir := "../"
	getDirectoryListing(rootDir)
}
