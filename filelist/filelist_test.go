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

func indexOf(target string, arrayOfStrings []string) int {
	for i, str := range arrayOfStrings {
		if target == str {
			return i
		}
	}
	return -1
}

func TestGetDirectoryListing(t *testing.T) {
	expectedDirContents := []string{"../filelist", "../filelist/filelist.go", "../filelist/filelist_test.go"}
	rootDir := "../filelist"

	dirContents, err := GetDirectoryListing(rootDir)
	if err != nil {
		t.Error("GetDirectoryListing() returned an error: %v\n", err)
	}
	if len(dirContents) != len(expectedDirContents) {
		t.Error("unexpected directory results: %v\n", dirContents)
	}

	for _, target := range expectedDirContents {
		if indexOf(target, dirContents) == -1 {
			t.Error("Could not find %s in %v\n", target, dirContents)
		}
	}
}

func TestFindHTMLFiles(t *testing.T) {
	rootDir := "../test-data"
	expectedDirContents := []string{"../test-data/index.html", "../test-data/persona-demo.html"}
	dirContents, err := FindHTMLFiles(rootDir)
	if err != nil {
		t.Error("FindHTMLFiles() returned an error: %v\n", err)
	}
	if len(expectedDirContents) != len(dirContents) {
		t.Error("Expected %d found %d files.\n", len(expectedDirContents), len(dirContents))
	}
	for _, target := range expectedDirContents {
		if indexOf(target, dirContents) == -1 {
			t.Error("Cound not find %s in %v\n", target, dirContents)
		}
	}
}
