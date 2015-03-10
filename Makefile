
test:
	gofmt -w words/words.go
	golint words/words.go
	gofmt -w words/words_test.go
	golint words/words_test.go
	go test -v ./words
	gofmt -w filelist/filelist.go
	golint filelist/filelist.go
	gofmt -w filelist/filelist_test.go
	golint filelist/filelist_test.go
	go test -v ./filelist
	gofmt -w makewordlist.go
	golint makewordlist.go

build: makewordlist.go words/words.go filelist/filelist.go
	go build makewordlist.go

clean: makewordlist
	rm makewordlist

install: makewordlist.go
	go install makewordlist.go
