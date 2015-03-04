/**
 * exp-01.go - experiemental file scanner
 *
 * @author R. S. Doiel, <rsdoiel@usc.edu>
 * copyright (c) 2015 All rights reserved.
 * Released under the Simplified BSD License.
 */
package main

import (
	"../filelist"
	"log"
)

func main() {
	rootDir := "../"
	log.Println("List all files.")
	dirContents, err := filelist.GetDirectoryListing(rootDir)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
	for i := 0; i < len(dirContents); i++ {
		log.Printf("%v\n", dirContents[i])
	}
	log.Println("List only HTML files.")
	dirContents, err = filelist.FindHTMLFiles(rootDir)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
	for i := 0; i < len(dirContents); i++ {
		log.Printf("%v\n", dirContents[i])
	}

}
