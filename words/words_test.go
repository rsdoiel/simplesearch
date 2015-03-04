/**
 * words_test.go - tests for words.go
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package words

import (
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
	w := new(Words)
	if w == nil {
		t.Error("Could not create Words struct.")
	}
}

func TestStripTags(t *testing.T) {
	expectedText := "This is a paragraph"
	plainText, err := StripTags("<p>This is a paragraph</p>")
	if err != nil {
		t.Error(err)
	}
	if expectedText != plainText {
		t.Errorf("[%s] != [%s]\n", expectedText, plainText)
	}
}

func TestsWordList(t *testing.T) {
	expectedWords := []string{"This", "si", "a", "test", "A", "title", "and", "paragraph"}

	w, err := WordList("<body><header>This si a test</header><h1>A title</h1><p>and a paragraph</p></body>")
	if err != nil {
		t.Errorf("WordList() returned error %v\n", err)
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
