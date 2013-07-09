# Schumix Webadmin

#### LICENCE: GNU LGPL 3

## Required packages/programs for compilation/running without building
* `go` programming language<br>
	You can install it by: <br>
	Debian-based: `sudo apt-get install golang`<br>
	Archlinux: `sudo pacman -S go`<br>
* Setup `GOPATH`<br>
	`mkdir ~/.gopath`<br>
	`export GOPATH=~/.gopath`<br>
	You can change `~/.gopath` to whatever you want.<br>
### Packages
* SQLite go binding: 	`go get github.com/mattn/go-sqlite3`
* Backend web services: `go get github.com/hoisie/web`
* Session manager: 	`go get github.com/mattn/go-session-manager`

After you can run it by:<br>
`go run main.go`<br>
Access the site at `localhost:45987` with default port.<br>
You can change the site's default settings in `config.json`.
