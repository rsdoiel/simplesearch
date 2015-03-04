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

func TestWords(t *testing.T) {
	w := new(Words)
	if w == nil {
		t.Error("Could not create Words struct.")
	}
	w.Words = make(map[string][]int, 100)
	w.Words["one"] = append(w.Words["one"], 1)
	if w.Words["one"][0] != 1 {
		t.Error("Could not append 1 to Words.")
	}

	w2 := New()
	if w2 == nil {
		t.Error(".New() failed to create a new Words structure")
	}
	w2.Files = append(w2.Files, "test-01.txt")
	w2.Words["Hello"] = append(w2.Words["Hello"], 0)
	w2.Words["World"] = append(w2.Words["World"], 0)

	if len(w2.Files) != 1 {
		t.Error(".Files should be length 1")
	}
	if len(w2.Words["Hello"]) != 1 {
		t.Error(".Words[\"Hello\"] should be length 1")
	}
	if len(w2.Words["World"]) != 1 {
		t.Error(".Words[\"World\"] should be length 1")
	}
	if w2.Words["Fred"] != nil {
		t.Error(".Words[\"Fred\"] should be nil")
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
	w := New()
	if err := w.MergeWords("test-01.html", []string{"This", "iS", "a", "TEST"}); err != nil {
		t.Errorf("MergeWords returned error %v\n", err)
	}
	if len(w.Files) != 1 {
		t.Errorf("Should have a single file in .Files: %v\n", w)
	}
	if hasString("test-01.html", w.Files) == false {
		t.Errorf("MergeWords() failed did not add words for test.html: %v\n", w)
	}
	if w.Words["this"] == nil && w.Words["this"][0] == 0 {
		t.Errorf("Word 'this' should be in w.Words: %v\n", w.Words)
	}
	if w.Words["is"] == nil && w.Words["is"][0] == 0 {
		t.Errorf("Word 'is' should be in w.Words: %v\n", w.Words)
	}
	if w.Words["a"] == nil && w.Words["a"][0] == 0 {
		t.Errorf("Word 'a' should be in w.Words: %v\n", w.Words)
	}
	if w.Words["test"] == nil && w.Words["test"][0] == 0 {
		t.Errorf("Word 'test' should be in w.Words: %v\n", w.Words)
	}

	if err := w.MergeWords("test-02.html", []string{"This", "is", "not", "a", "unique", "test"}); err != nil {
		t.Errorf("MergeWords() returned error %v\n", err)
	}
	if len(w.Files) != 2 {
		t.Errorf("Should have two files in struct: %v\n", w)
	}
	if len(w.Words["this"]) != 2 {
		t.Errorf("Should have two entries for 'this' %v\n", w.Words)
	}
	if len(w.Words["not"]) != 1 {
		t.Errorf("Should have one entry for 'not' %v\n", w.Words)
	}
}

func TestToJSON(t *testing.T) {
	t.Error("ToJSON() not implemented")
}
