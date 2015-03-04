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
)

// Words contains the data structures for building a file list and inverted word index.
type Words struct {
	Files []string
	Words map[string][]int
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
	w.Files = append(w.Files, pathname)
	//w.Words = append(w.Words, words)
	if hasString(pathname, w.Files) == false {
		return errors.New("Did not add filename " + pathname)
	}
	return nil
}

// ToJSON - render the Words data structure a JSON
func (w *Words) ToJSON() (string, error) {
	return "", errors.New("ToJSON() not implemented")
}
