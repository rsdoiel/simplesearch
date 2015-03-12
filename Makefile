
test: words/words.go filelist/filelist.go cmds/makewordlist/makewordlist.go
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
	gofmt -w cmds/makewordlist/makewordlist.go
	golint cmds/makewordlist/makewordlist.go

build: cmds/makewordlist/makewordlist.go words/words.go filelist/filelist.go
	go build cmds/makewordlist/makewordlist.go

clean:
	if [ -f makewordlist ]; then rm makewordlist; fi
	if [ -f files.json ]; then rm files.json; fi
	if [ -f wordlist.json ]; then rm wordlist.json; fi

install: cmds/makewordlist/makewordlist.go
	go install cmds/makewordlist/makewordlist.go

