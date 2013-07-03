# Schumix WebAdmin

### Required packages/programs
* `go` programming language<br>
	You can install it by: <br>
	Debian-based: `sudo apt-get install golang`<br>
	Archlinux: `sudo pacman -S go`<br>
* Setup `GOPATH`<br>
	`mkdir ~/.gopath`<br>
	`export GOPATH=~/.gopath`<br>
	You can change `~/.gopath` to whatever you want.<br>
* `github.com/mattn/go-sqlite3` SQLite go binding<br>
	`go get github.com/mattn/go-sqlite3`

After you can run it by:<br>
`go run main.go`