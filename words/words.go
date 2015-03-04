package words

/**
 * words.go - ulities to convert an HTML files into an inverted word list.
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * @license Released under the Simplified BSD License.
 */

import (
	"errors"
	"fmt"
)

// A slice of ints containing pos indexes from another array or slice.
type IndexList []int

// Words contains the data structures for building a file list and inverted word index.
type Words struct {
	Files []string
	Words map[string]IndexList
}

// StripTags removes HTML markup returning only CData
func StripTags(html string) (string, error) {
	var (
		outSlice []byte
		outError error
	)

	byteSlice := []byte(html)
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
	return nil, errors.New("WordList() not implemented")
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
	// Only add unique pathname
	if i := indexOf(pathname, w.Files); i == -1 {
		w.Files = append(w.Files, pathname)
	}
	// Confirm position of pathname in the list
	i := indexOf(pathname, w.Files)
	if i == -1 {
		return errors.New(fmt.Sprintf("Could not update Words for %s", pathname))
	}
	for _, word := range words {
		// Create a slot for the map if needed.
		if w.Words[word] == nil {
			w.Words[word] = make(IndexList)
		}
		// Append index to word list
		w.Words[word] = append(w.Words[word], i)
		// Confirm we still have our file index list for word.
		if w.Words[word] == nil {
			return errors.New(fmt.Sprintf("Could not add %s to words %v", word, w.Words))
		}
	}
	return nil
}

// ToJSON - render the Words data structure a JSON
func (w *Words) ToJSON() (string, error) {
	return "", errors.New("ToJSON() not implemented")
}
