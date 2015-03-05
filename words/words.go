package words

/**
 * words.go - ulities to convert an HTML files into an inverted word list.
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * @license Released under the Simplified BSD License.
 */

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/html"
	"io"
	"regexp"
	"strings"
)

const (
	delimitingCharacters = " ~!@#$%^&*()_+`-={}[];':\"<>?,./|\t\r\n"
	collapsingCharacters = "( |,|!|-|\t|\r|\n)+"
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

// Flatten removes HTML markup returning only CData as a space delimited list of words.
func Flatten(srcBytes []byte) []byte {
	var (
		outSlice [][]byte
	)

	z := html.NewTokenizer(bytes.NewReader(srcBytes))
	depth := 0
	inCData := true
	moreHTML := true
	cData := []byte("")
	for moreHTML == true {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				moreHTML = false
			}
		case html.TextToken:
			cData = z.Text()
			if depth > 0 && inCData == true {
				outSlice = append(outSlice, cData)
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if tt == html.StartTagToken {
				if bytes.Equal(tn, []byte("script")) == true || bytes.Equal(tn, []byte("head")) == true {
					inCData = false
				}
				depth++
			} else {
				if bytes.Equal(tn, []byte("script")) == true || bytes.Equal(tn, []byte("head")) == true {
					inCData = true
				}
				depth--
			}
		}
	}
	re := regexp.MustCompile(collapsingCharacters)
	return bytes.Trim(re.ReplaceAll(bytes.Join(outSlice, []byte("")), []byte(" ")), delimitingCharacters)
}

// WordList scans HTML source and returns a list of words found.
func WordList(srcBytes []byte) [][]byte {
	return bytes.Split(srcBytes, []byte(" "))
}

func indexOfString(l []string, target string) int {
	for i, s := range l {
		if target == s {
			return i
		}
	}
	return -1
}

func indexOf(l [][]byte, target string) int {
	t := []byte(target)
	for i, s := range l {
		if bytes.Equal(t, s) == true {
			return i
		}
	}
	return -1
}

func containsString(l []string, target string) bool {
	for _, s := range l {
		if target == s {
			return true
		}
	}
	return false
}

func containsInt(l []int, i int) bool {
	for _, j := range l {
		if i == j {
			return true
		}
	}
	return false
}

// MergeWords - given a path and list of words update the Words datastructure
func (w *Words) MergeWords(pathname string, words [][]byte) bool {
	var (
		key string
		i   int
	)
	// Only add unique pathname
	if i = indexOfString(w.Files, pathname); i == -1 {
		w.Files = append(w.Files, pathname)
	}
	// Confirm position of pathname in the list
	i = indexOfString(w.Files, pathname)
	if i == -1 {
		return false
	}
	for _, word := range words {
		// Create a slot for the map if needed.
		if key = strings.Trim(strings.ToLower(string(word[:])), delimitingCharacters); key != "" {
			if w.Words[key] == nil {
				w.Words[key] = make([]int, 1)
				w.Words[key][0] = i
			} else if containsInt(w.Words[key], i) == false {
				// Append index to word list
				w.Words[key] = append(w.Words[key], i)
			}
			// Confirm we still have our file index list for word.
			if w.Words[key] == nil {
				return false
			}
		}
	}
	return true
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
