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

// MergeWords - given a path and list of words update the Words datastructure
func (*Words) MergeWords(pathname string, words []string) error {
	return errors.New("MergePathWordList() not implemented")
}

// ToJSON - render the Words data structure a JSON
func (*Words) ToJSON() (string, error) {
	return "", errors.New("ToJSON() not implemented")
}
