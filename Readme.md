# Schumix Webadmin [![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/Schumix/webadmin/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

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
* SQLite go binding:    `go get github.com/mattn/go-sqlite3`
* Backend web services: `go get github.com/hoisie/web`
* Session manager:      `go get github.com/mattn/go-session-manager`
* Go version comparer:	`go get github.com/Jackneill/gosemver`

### Configuration

The configuration file (`config.json`) can be found at the root folder.

* `"Timeout" : "30m"`<br>
The duration of time which the bot should try to reconnect to the server in case the connection is lost.
Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

### Running

You can run it by:<br>
`./run.sh`<br>
Access the site at `localhost:45987` with default port.<br>
You can change the site's default settings in `config.json`.
