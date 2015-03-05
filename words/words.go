package words

/**
 * words.go - ulities to convert an HTML files into an inverted word list.
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * @license Released under the Simplified BSD License.
 */

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Words contains the data structures for building a file list and inverted word index.
type Words struct {
	Files []string
	Words map[string][]int
}

// New - initialize a new Words structure.
func New() *Words {
	w := new(Words)
	w.Words = make(map[string][]int, 100)
	return w
}

// StripTags removes HTML markup returning only CData
// consider replacing this with something based on golang.org/x/net/html parser.
func StripTags(html string) (string, error) {
	var (
		outSlice []byte
		outError error
	)

	byteSlice := []byte(strings.ToLower(html))
	inCData := false
	for _, c := range byteSlice {
		if c == '>' {
			inCData = true
		} else if c == '<' {
			inCData = false
		} else if inCData == true {
			outSlice = append(outSlice, c)
			if outSlice == nil {
				outError = errors.New("Cannot append element to output string")
			}
		}
	}
	return string(outSlice[:]), outError
}

// WordList scans HTML source and returns a list of words found.
func WordList(src string) ([]string, error) {
	words := strings.Split(strings.ToLower(src), " ")
	if len(words) == 0 {
		return nil, errors.New("No words found.")
	}
	// Trim leading/trailing puncuation and spaces.
	for i, item := range words {
		words[i] = strings.Trim(item, " .~!@#$%^&*()_+`-={}[];':\"<>?,./\n\r")
	}
	return words, nil
}

func indexOf(target string, l []string) int {
	for i, s := range l {
		if target == s {
			return i
		}
	}
	return -1
}

func hasString(target string, l []string) bool {
	for _, s := range l {
		if target == s {
			return true
		}
	}
	return false
}

// MergeWords - given a path and list of words update the Words datastructure
func (w *Words) MergeWords(pathname string, words []string) error {
	var (
		key string
		i   int
	)
	// Only add unique pathname
	if i = indexOf(pathname, w.Files); i == -1 {
		w.Files = append(w.Files, pathname)
	}
	// Confirm position of pathname in the list
	i = indexOf(pathname, w.Files)
	if i == -1 {
		return errors.New(fmt.Sprintf("Could not update Words for %s", pathname))
	}
	for _, word := range words {
		// Create a slot for the map if needed.
		key = strings.ToLower(word)
		if w.Words[key] == nil {
			w.Words[key] = make([]int, 1)
			w.Words[key][0] = i
		} else {
			// Append index to word list
			w.Words[key] = append(w.Words[key], i)
		}
		// Confirm we still have our file index list for word.
		if w.Words[key] == nil {
			return errors.New(fmt.Sprintf("Could not add %s to words %v", word, w.Words))
		}
	}
	return nil
}

// ToJSON - render the Words data structure a JSON
func (w *Words) ToJSON() (string, string, error) {
	fileList, err := json.Marshal(w.Files)
	if err != nil {
		return "", "", err
	}
	wordList, err := json.Marshal(w.Words)
	if err != nil {
		return string(fileList[:]), "", err
	}
	return string(fileList[:]), string(wordList[:]), nil
}
