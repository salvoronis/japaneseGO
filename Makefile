DESTDIR=./bin
DEP1=github.com/gotk3/gotk3/gtk
APPNAME=kanjitest

all: build

run: actions.go myshit.go
	go run actions.go myshit.go

build: dep
	go build -o $(DESTDIR)/$(APPNAME) .

dep:
	go get $(DEP1)

clean:
	rm $(DESTDIR)/$(APPNAME)