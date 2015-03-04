/**
 * filelist.go - Scan a path an return qa list of files found.
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package filelist

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetDirectoryListing(pathname string) ([]string, error) {
	var dirContents []string

	dir, err := os.Open(pathname)
	if err == nil {
		defer dir.Close()
		filepath.Walk(pathname, func(filename string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			dirContents = append(dirContents, filename)
			return nil
		})
	}
	return dirContents, err
}

func FindHTMLFiles(pathname string) ([]string, error) {
	var dirContents []string

	dir, err := os.Open(pathname)
	if err == nil {
		defer dir.Close()
		filepath.Walk(pathname, func(filename string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			ext := strings.ToLower(path.Ext(filename))
			if ext == ".html" || ext == ".htm" {
				dirContents = append(dirContents, filename)
			}
			return nil
		})
	}
	return dirContents, err
}
