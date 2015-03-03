package main

import (
	"log"
	"os"
)

const maxFileNames = 100

type DirectoryListing struct {
	RootPathname string
	Root         *os.FileInfo
	Files        []os.FileInfo
	SubFolders   []os.FileInfo
	TotalEntries int
}

func GetDirectoryListing(pathname string) (DirectoryListing, error) {
	dir, err := os.Open(pathname)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
	defer dir.Close()

	listing := new(DirectoryListing)
	listing.RootPathname = pathname
	listing.Root = dir
	dirContents, err := dir.Readdir(maxFileNames)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
	log.Printf("DEBUG dirContents: %v\n", dirContents)

	allDone := false
	for allDone == false {
		for i := 0; i < len(dirContents); i++ {
			log.Printf("DEBUG %v\n", dirContents[i])
		}
		allDone = true
	}

	return nil, errors.New("ListDirectory() not fully implemented")
}

func main() {
	rootDir := "./demo/site-01"
	log.Println("DEBUG " + rootDir)
	for i := 0; i < len(dirContents); i++ {
		if dirContents[i].IsDir() == true {
			log.Printf("DEBUG %s is a directory\n", rootDir)
		}
	}
}
