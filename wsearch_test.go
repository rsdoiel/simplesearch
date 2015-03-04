package main_test

import (
	"./filelist"
	"./words"
	"testing"
)

func indexOf(target string, words []string) int {
	for i, s := range words {
		if target == s {
			return i
		}
	}
	return -1
}

func TestWords(t *testing.T) {
	w := new(words.Words)
	if w == nil {
		t.Error("Could not create words.Words struct.")
	}
}

func TestStripTags(t *testing.T) {
	expectedText := "This is a paragraph"
	plainText, err := words.StripTags("<p>This is a paragraph</p>")
	if err != nil {
		t.Error(err)
	}
	if expectedText != plainText {
		t.Errorf("[%s] != [%s]\n", expectedText, plainText)
	}
}

func TestsWordList(t *testing.T) {
	expectedWords := []string{"This", "si", "a", "test", "A", "title", "and", "paragraph"}

	w, err := words.WordList("<body><header>This si a test</header><h1>A title</h1><p>and a paragraph</p></body>")
	if err != nil {
		t.Errorf("words.WordList() returned error %v\n", err)
	}
	if len(w) != len(expectedWords) {
		t.Errorf("len(w) %d does not equal len(expectedWords) %d\n", len(w), len(expectedWords))
	}
	for _, s := range expectedWords {
		if indexOf(s, w) == -1 {
			t.Error("Could not find %s in %v\n", s, w)
		}
	}
}

func TestMergeWords(t *testing.T) {
	t.Error("MergePathWordList() not implemented")
}

func TestToJSON(t *testing.T) {
	t.Error("ToJSON() not implemented")
}

func TestGetDirectoryListing(t *testing.T) {
	expectedDirContents := []string{"filelist", "filelist/filelist.go", "filelist/filelist_test.go"}
	rootDir := "filelist"

	dirContents, err := filelist.GetDirectoryListing(rootDir)
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
	t.Error("TestFindHTMLFiles() not implemented.")
}

func TestCLI(t *testing.T) {
	t.Error("TestCLI() not implemented.")
}
