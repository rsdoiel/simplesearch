/**
 * wsearch.go - convert static website into a JSON list of files
 * and JSON inverted word list.
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package main

import (
	"./filelist"
	"./words"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const usageText = `
 USAGE: %s DIRECTORY_TO_SCAN

 %s scans a directory and generates files.json and words.json.

 OPTIONS:

     -h, --help This help message.

`

func usage(msg string, error_code int) {
	prog := filepath.Base(os.Args[0])
	fmt.Printf(usageText, prog, prog)
	fmt.Println(msg)
	os.Exit(error_code)
}

func containsString(l []string, target string) bool {
	for _, item := range l {
		if target == item {
			return true
		}
	}
	return false
}

func main() {
	if containsString(os.Args, "-h") == true || containsString(os.Args, "--help") {
		usage("", 0)
	}
	if len(os.Args) != 2 {
		usage("Your are missing the directory to scan.", 1)
	}
	rootPath := os.Args[1]
	dirContents, err := filelist.FindHTMLFiles(rootPath)
	if err != nil {
		usage(fmt.Sprintf("%v\n", err), 1)
	}
	log.Printf("Files %v\n", dirContents)
	//w := words.New()
	for i, fname := range dirContents {
		log.Printf("processing %d %s\n", i, fname)
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatal(err)
		}
		startCut := 0
		endCut := len(data) - 1
		if bytes.Contains(bytes.ToLower(data), []byte("<body")) == true {
			startCut = bytes.Index(bytes.ToLower(data), []byte("<body"))
		}

		src, err := words.StripTags(string(data[startCut:endCut]))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("DEBUG %s\n", src)
	}
}
