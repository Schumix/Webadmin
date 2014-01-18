CC=go build
CFLAGS=-o bin/Webadmin

all: webadmin

webadmin:
	$(CC) $(CFLAGS) src/*.go
	cp configs/config.json bin/config.json
	cp sql/Schumix.db3 bin/Schumix.db3

dep:
	# todo: check for $GOPATH, if not exists create it
	go get github.com/Schumix/go-sqlite3
	go get github.com/Schumix/gosemver
	go get github.com/hoisie/web
	go get github.com/mattn/go-session-manager

rebuild: clean all

clean:
	rm -rf bin
