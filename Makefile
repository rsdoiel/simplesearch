
test:
	go test -v ./words
	go test -v ./filelist
	go test -v
	golint makewordlist.go
	golint words/words.go
	golint words/words_test.go
	golint filelist/filelist.go
	golint filelist/filelist_test.go

makewordlist: makewordlist.go words/words.go filelist/filelist.go
	
build:
	go build makewordlist.go

clean:
	rm makewordlist
