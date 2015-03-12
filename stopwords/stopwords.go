package stopwords

/**
 * stopwords.go - A collection of stopwords based on lanaguages.
 * This is a naive proof of concept and if this moves beyond
 * this then a proper approach to internationalization should be
 * taken.
 */
import (
	"errors"
)

// StopWordList is the datastructure for holding a list of StopWords.
type StopWordList []string

// New creates a Stopword list structure that is suitable for use with
// StopWords methods.
func New([]string) (StopWordList, error) {
	return nil, errors.New("New() StopWordList not implemented")
}

// IsStopWordString validates if a given string is a stop word or not.
func (stopwords StopWordList) IsStopWordString(s string) bool {
	for _, w := range stopwords {
		if w == s {
			return true
		}
	}
	return false
}

// IsStopWord validates a byte array determining if it represents a stop word or not.
func (stopwords *StopWordList) IsStopWord(b []byte) bool {
	return stopwords.IsStopWordString(string(b))
}
