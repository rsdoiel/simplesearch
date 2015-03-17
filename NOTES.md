
# Character encodings

Parsing the HTML correctly to split into words requires determining how the
file is encoded. If the file is fetched from a URL that information may be
supplied as part of the HTTP headers but sometimes you have a file on disc
without clear indications of the encoding type. In the best libraries for
making a best guess on encoding have been part of the the stand libraries
supplied with the programming language.  Go, as a relatively young, language
has a few options.

+ [golang.org/x/net/html/charset](https://godoc.org/golang.org/x/net/html/charset) - A library to guess encoding type
    + returns mime type, string name and certainty as bool
    + maintained by Go Authors
    + go get golang.org/x/net/html/charsets
+ [net/http/sniff](https://golang.org/src/net/http/sniff.go) - A library to guess encoding type based on analyzing some of the text in a file.
    + returns mime type
    + analsys based on first 512 bytes
    + maintained by Go Authors
+ [github.com/saintfish/chardet](https://github.com/saintfish/chardet) - a library created about 3 yrs ago by saintfish derived from UCI
    + go get github.com/saintfish/chardet
+ [golang.org/x/text/encoding](https://godoc.org/golang.org/x/text/encoding) - Package encoding defines an interface for character encodings, such as Shift JIS and Windows 1252, that can convert to and from UTF-8.
    + go get golang.org/x/text/encoding


# Misc links

+ http://stackoverflow.com/questions/24555819/golang-persist-using-iso-8859-1-charset

