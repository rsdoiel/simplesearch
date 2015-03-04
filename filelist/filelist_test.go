package filelist_test

import (
	"../filelist"
	"fmt"
	"testing"
)

func indexOf(target string, arrayOfStrings []string) bool {
	for _, str := range arrayOfStrings {
		if target == str {
			return true
		}
	}
	return false
}

func TestGetDirectoryListing(t *testing.T) {
	expectedDirContents := []string{
		"../filelist",
		"../filelist/filelist.go",
		"../filelist/filelist_test.go",
	}
	rootDir := "../filelist"
	dirContents, err := filelist.GetDirectoryListing(rootDir)
	if err != nil {
		t.Error("GetDirectoryListing() threw err %v\n", err)
	}
	fmt.Printf("DEBUG %v\n", dirContents)
	if len(dirContents) != 3 {
		t.Error("unexpected directory results: %v\n", dirContents)
	}

	for _, target := range expectedDirContents {
		if indexOf(target, dirContents) == false {
			t.Error("Could not find %s in %v\n", target, dirContents)
		}
	}
}
