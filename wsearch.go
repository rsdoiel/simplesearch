/**
 * html2wordlist.go - convert the body of a HTML file into a
 * plain text document without any HTML elements.
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package main

import (
	"errors"
	"log"
)

func findHTMLFiles(root string) ([]string, error) {
	return nil, errors.New("findHTMLFiles() not implemented")
}

func extractWords(fname string) ([]string, error) {
	return nil, errors.New("extractWords() not implemented")
}

func main() {
	log.Println("Proof of concept converting an HTML document into a Wordlist")
	fileList, err := findHTMLFiles("demo")
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
	log.Printf("DEBUG %v\n", fileList)
	for i := 0; i < len(fileList); i++ {
		log.Printf("DEBUG process file no %i: %s\n", i, fileList[i])
		words, err := extractWords(fileList[i])
		log.Printf("DEBUG words in file: %v : %v\n", words, err)
		// Push the new words onto the wordlist
	}
}
