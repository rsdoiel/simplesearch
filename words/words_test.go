/**
 * words_test.go - tests for words.go
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package words_test

import (
	"../words"
	"testing"
)

func indexOf(target string, words []string) bool {
	for _, s := range words {
		if target == s {
			return true
		}
	}
	return false
}

func TestWords(t *testing.T) {
	w := new(words.Words)
	if w == nil {
		t.Error("Could not create words.Words element.")
	}
}

func TestStripTags(t *testing.T) {
	expectedText := "This is a paragraph"
	plainText, err := words.StripTags("<p>This is a paragraph</p>")
	if err != nil {
		t.Error(err)
	}
	if expectedText != plainText {
		t.Error(expectedText + " != " + plainText)
	}
}

func TestsWordList(t *testing.T) {
	expectedWords := []string{"This", "si", "a", "test", "A", "title", "and", "paragraph"}

	w, err := words.WordList("<body><header>This si a test</header><h1>A title</h1><p>and a paragraph</p></body>")
	if err != nil {
		t.Error(err)
	}
	if len(w) != len(expectedWords) {
		t.Error("len(w) does not equal len(expectedWords)")
	}
	for _, s := range expectedWords {
		if indexOf(s, w) == false {
			t.Error("Could not find " + s)
		}
	}
}

func TestMergeWords(t *testing.T) {
	t.Error("MergePathWordList() not implemented")
}

func TestToJSON(t *testing.T) {
	t.Error("ToJSON() not implemented")
}
