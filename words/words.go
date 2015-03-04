/**
 * words.go - ulities to convert an HTML files into an inverted word list.
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package words

import (
	"errors"
)

type Words struct {
	Files []string
	Words map[string][]int
}

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

func WordList(src string) ([]string, error) {
	return nil, errors.New("WordList() not implemented")
}

func (*Words) MergeWords(pathname string, words []string) error {
	return errors.New("MergePathWordList() not implemented")
}

func (*Words) ToJSON() (string, error) {
	return "", errors.New("ToJSON() not implemented")
}
