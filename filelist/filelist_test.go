package filelist_test

import (
	"../filelist"
	"log"
)

func main() {
	log.Fatalf("filelist_test.go not implemented")
	rootDir := "../filelist"
	dirContents, err := filelist.GetDirectoryListing(rootDir)
	if err != nil {
		log.Fatalf("GetDirectoryListing() threw err %v\n", err)
	}
	log.Printf("dirContents: %v\n", dirContents)
}
