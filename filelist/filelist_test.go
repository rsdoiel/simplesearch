/**
 * filelist_test.go - test the filelist.go library.
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved
 * Released under the Simplified BSD License
 */
package filelist

import (
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
	expectedDirContents := []string{"../filelist", "../filelist/filelist.go", "../filelist/filelist_test.go"}
	rootDir := "../filelist"

	dirContents, err := GetDirectoryListing(rootDir)
	if err != nil {
		t.Error("GetDirectoryListing() threw error: %v\n", err)
	}
	if len(dirContents) != len(expectedDirContents) {
		t.Error("unexpected directory results: %v\n", dirContents)
	}

	for _, target := range expectedDirContents {
		if indexOf(target, dirContents) == false {
			t.Error("Could not find %s in %v\n", target, dirContents)
		}
	}
}

func TestFindHTMLFiles(t *testing.T) {
	t.Error("TestFindHTMLFiles() not implemented.")
}
