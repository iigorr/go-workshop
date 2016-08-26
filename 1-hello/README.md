# 1 hello

## Setup

Please follow instructions below
(Official instructions are under https://golang.org/doc/install)

Linux/Mac:
* Download binaries from https://golang.org/dl/
* extract to /usr/lib/go
* add /usr/lib/go/bin to your $PATH
* `mkdir -p $HOME/work/src`
* `export GOPATH=$HOME/work`

Windows:
* Download msi from https://golang.org/dl/
* Install under C:\go (or anywhere else)
* You may need to add C:\go\bin to your %PATH%
* `mkdir C:\work\src`
* `set GOPATH=C:\work`

Debian (not recomended):
* `sudo apt-get install golang-go`
* `mkdir -p $HOME/work/src`
* `export GOPATH=$HOME/work`

## Test installation

```bash
go version
```

## Exercise 1.1

Write, compile & run a simple "Hello World!" program.

### Write and run
* Create a file named $GOPATH/src/hello/hello.go`
* Implement
* `cd $GOPATH/src/hello`
* `go run hello.go`

### Compile & Run
* `go build`
* `./hello`

### Play around with tooling

* Run `go` to list all go commands
* Run `go help` to get details
* e.g. make sure the formatting is correct (if you don't use an extension): `go fmt hello.go`
* take a look at get, build, doc, test, vet, install



## Exercise 1.2

Write, compile & run a "Hello World!" programm, that reads a name from the standard input and prints a friendly greeting.

e.g.
> % ./hello-stdin

> Name: Gopher_

> Hello, Gopher!


You will need to import the package "bufio" and "os", so you can create a reader:
`reader := bufio.NewReader(os.Stdin)`

see https://golang.org/pkg/bufio/


## Exercise 1.3
Write, compile & run a "Hello World!" programm, that writes "Hello World!" as ascii-banner.
Use github.com/getwe/figlet4go, or github.com/CrowdSurge/banner for example.

Make sure to run `go get` to fetch the library. e.g. `go get github.com/getwe/figlet4go`


## Exercise 1.4

Write, compile & run a simple "Hello World!" web application, that listens on port 8080 and responds with "Hello, {path}", where {path} ist the url path

e.g.
> % curl localhost:8080/go/workshop

> Hello, go/workshop


Check out http.ListenAndServe in "net/http" (https://golang.org/pkg/net/http/)