package recode

/**
 * recode_Test.go - test the conversion routines provided by recode.go.
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved
 * Released under the Simplified BSD License
 */
import (
	"testing"
)

func TestLatin1ToUtf8(t *testing.T) {
	expected := "This should convert to/from iso-8859-1 to UTF-8"
	lat1 := make([]rune, len(expected))
	for i, b := range expected {
		lat1[i] = rune(b)
	}
	result, err := Latin1ToUtf8(lat1)
	if result != expected {
		t.Error(err)
	}
}
