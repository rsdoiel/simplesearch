
# simplesearch

Poof of concept command line tool to generate file and inverted word lists useful for naive search implementations.

## Basic idea

Simple search can be implemented with an inverted word list. For small websites evaluating an 
pre-calculated inverted word list for search terms is reasonable straight forward. This might be
useful where you have a dynamic language like PHP avialble but do not want to use a database to
back the website.  Likewise if the list is small enough you could even have the web browser caliculate
the search results.

## Proof of concept

I have written a tool, _makewordlist_, that will generate two JSON files, files.json and words.json, from
a provided path (e.g. /my/htdocs). You can then read these two files in PHP and find which files contain
the words you are searching for.

Commands to run--

```shell
    go build makewordlist.go
    ./makewordlist test-data
    php demo-simplesearch.php otto javascript persona fred Jane
```

This PHP search demo should return something like--

```
    These are the pages that have some of the words (i.e. otto, javascript, persona):
    Page: test-data/index.html
    Page: test-data/index.html
    Page: test-data/persona-demo.html
```



