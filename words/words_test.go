/**
 * words_test.go - tests for words.go
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package words

import (
	"bytes"
	"strings"
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
	expectedText := []byte("This is a paragraph")
	plainText := StripTags([]byte("<p>This is a paragraph</p>"))
	if bytes.Equal(expectedText, plainText) != true {
		t.Errorf("[%v] != [%v]\n", expectedText, plainText)
	}
	expectedText = []byte("This is a test, but without the script element.")
	plainText = StripTags([]byte("<body><p>This is a test<script>console.log('Should not include');</script>, but without the script element.</body>"))
	if len(expectedText) != len(plainText) {
		t.Errorf("different lengths %d <> %d\n", len(expectedText), len(plainText))
	}
	if bytes.Equal(expectedText, plainText) == false {
		t.Errorf("[%v] != [%v]\n", expectedText, plainText)
	}
}

func TestsWordList(t *testing.T) {
	src := []byte("<body><header>This si a test</header><h1>A title</h1><p>and a paragraph</p></body>")
	expectedWords := [][]byte{[]byte("This"), []byte("si"), []byte("a"), []byte("test"), []byte("A"), []byte("title"), []byte("and"), []byte("paragraph")}
	s := StripTags(src)
	w := WordList(s)
	if w == nil {
		t.Errorf("WordList() returned empty list")
	}
	if len(w) != len(expectedWords) {
		t.Errorf("len(w) %d does not equal len(expectedWords) %d\n", len(w), len(expectedWords))
	}
	for _, item := range expectedWords {
		if indexOf(w, string(item[:])) == -1 {
			t.Error("Could not find %v in %v\n", item, w)
		}
	}
}

func TestMergeWords(t *testing.T) {
	w := New()
	if w.MergeWords("test-01.html", []string{"This", "iS", "a", "TEST"}) == false {
		t.Errorf("MergeWords returned false %v\n", w)
	}
	if len(w.Files) != 1 {
		t.Errorf("Should have a single file in .Files: %v\n", w)
	}
	if containsString(w.Files, "test-01.html") == false {
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

	if w.MergeWords("test-02.html", []string{"This", "is", "not", "a", "unique", "test"}) == false {
		t.Errorf("MergeWords() returned false %v\n", w)
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

	if w.MergeWords("test-03.html", []string{"THIS", "SPACE", "", "?", "!", "A", "\n\r"}) == false {
		t.Errorf("MergeWords() returned false %v\n", w)
	}
	if w.Words[""] != nil {
		t.Errorf("MergeWords() is storing empty string: %v\n", w)
	}
	if w.Words["?"] != nil {
		t.Errorf("MergeWords() is storing question mark: %v\n", w)
	}
	if w.Words["!"] != nil {
		t.Errorf("MergeWords() is storing bang: %v\n", w)
	}
	if w.Words["\n\r"] != nil {
		t.Errorf("MergeWords() is storing line feed/carriage return: %v\n", w)
	}

}

func TestToJSON(t *testing.T) {
	w := New()
	if w.MergeWords("test-01.html", []string{"This", "iS", "a", "TEST"}) == false {
		t.Errorf("MergeWords returned error %v\n", w)
	}
	if w.MergeWords("test-02.html", []string{"This", "is", "not", "a", "unique", "test"}) == false {
		t.Errorf("MergeWords() returned error %v\n", w)
	}
	fileList, wordList, err := w.ToJSON()
	if err != nil {
		t.Errorf("ToJSON() failed: %v, %v\n", w, err)
	}
	if strings.Contains(fileList, "test-01.html") == false || strings.Contains(fileList, "test-02.html") == false {
		t.Errorf("Missing filenames in list: %v -> %s\n", w.Files, fileList)
	}
	if strings.Contains(wordList, "\"this\":[0,1]") == false {
		t.Errorf("Missing 'This' values in word list: %v -> %s\n", w.Words, wordList)
	}
}
