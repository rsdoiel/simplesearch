/**
 * makewordlist.go - scan a static website and generate a files and inverted word
 * list in JSON.
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package main

import (
	"../../filelist"
	"../../words"
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

func usage(msg string, errorCode int) {
	prog := filepath.Base(os.Args[0])
	fmt.Printf(usageText, prog, prog)
	fmt.Println(msg)
	os.Exit(errorCode)
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
	w := words.New()
	for i, fname := range dirContents {
		log.Printf("processing %d %s\n", i, fname)
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatal(err)
		}
		src := words.Flatten(data)
		wordList := words.WordList(src)
		if w.MergeWords(fname, wordList) == false {
			log.Fatal(fmt.Sprintf("Could not add words for %s <-- %s\n", fname, wordList))
		}

	}
	fileList, invertedWordList, err := w.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("files.json", []byte(fileList), 0664)
	ioutil.WriteFile("wordlist.json", []byte(invertedWordList), 0664)
}
