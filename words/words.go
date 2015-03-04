/**
 * words.go - ulities to convert an HTML files into an inverted word list.
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package words

type Words struct {
	Files []string
	Words map[string][]int
}

func StripTags(html []byte) ([]byte, error) {
	inCData := false
	return nil, errors.New("StripTags() not implemented")
}

func WordList(src []byte) ([]string, error) {
	var wordList []string

	return nil, errors.New("WordList() not implemented")
}

func (*Words) MergeWords(pathname string, words []string) error {
	return errors.New("MergePathWordList() not implemented")
}

func (*Words) ToJSON() (string, error) {
	return "", errors.New("ToJSON() not implemented")
}
