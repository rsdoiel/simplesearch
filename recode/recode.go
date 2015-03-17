package recode

/**
 * recode.go - convert Latin 1 (iso-8859-1) to UTF-8
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 all rights reserved
 * Released under the Simplified BSD License.
 */
import (
	"errors"
)

func Latin1ToUtf8(in []byte) (string, error) {
	buf := make([]rune, len(in))
	for i, b := range in {
		buf[i] = rune(b)
	}
	return string(buf), errors.New("Latin1ToUTf8() not implemented.")
}
