package main

/**
 * lat1toutf8.go - convert a file encoded with Latin 1 (iso8859-1) to UTF-8.
 */
import (
	"../../recode"
	"fmt"
	"os"
	"strings"
)

func usage(msg string, exitCode int) {
	txt := `
	USAGE: recode LATIN_1_FILENAME UTF_8_FILENAME

	Options:

	-h, --help	This help message.

`
	fmt.Println(txt)
	if msg != "" {
		fmt.Println(msg)
	}
	os.Exit(exitCode)
}

func askForHelp(args []string) bool {
	var s string
	for _, arg := range args {
		s = strings.ToLower(arg)
		if s == "-h" || s == "--help" {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if askForHelp(args) == true {
		usage("", 0)
	}
	if len(args) != 2 {
		usage("Wrong number of arguments.", 1)
	}
	recode.Latin1ToUtf8(args[0], args[1])
}
