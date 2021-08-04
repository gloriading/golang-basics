# Golang Learning

## Official website starters - Part 1

> https://golang.org/doc/tutorial/getting-started

* `go mod init example.com/hello`
* `touch hello.go`
* install VSCode Go extension & install tools suggested from VSCode prompt

  ```
  // from VOCode OUTPUT:
  Tools environment: GOPATH=/Users/gloriading/go
  Installing 10 tools at /Users/gloriading/go/bin in module mode.
    gopkgs
    go-outline
    gotests
    gomodifytags
    impl
    goplay
    dlv
    dlv-dap
    staticcheck
    gopls

  Installing github.com/uudashr/gopkgs/v2/cmd/gopkgs (/Users/gloriading/go/bin/gopkgs) SUCCEEDED
  Installing github.com/ramya-rao-a/go-outline (/Users/gloriading/go/bin/go-outline) SUCCEEDED
  Installing github.com/ramya-rao-a/go-outline (/Users/gloriading/go/bin/go-outline) SUCCEEDED
  Installing github.com/cweill/gotests/gotests (/Users/gloriading/go/bin/gotests) SUCCEEDED
  Installing github.com/cweill/gotests/gotests (/Users/gloriading/go/bin/gotests) SUCCEEDED
  Installing github.com/fatih/gomodifytags (/Users/gloriading/go/bin/gomodifytags) SUCCEEDED
  Installing github.com/fatih/gomodifytags (/Users/gloriading/go/bin/gomodifytags) SUCCEEDED
  Installing github.com/josharian/impl (/Users/gloriading/go/bin/impl) SUCCEEDED
  Installing github.com/josharian/impl (/Users/gloriading/go/bin/impl) SUCCEEDED
  Installing github.com/haya14busa/goplay/cmd/goplay (/Users/gloriading/go/bin/goplay) SUCCEEDED
  Installing github.com/haya14busa/goplay/cmd/goplay (/Users/gloriading/go/bin/goplay) SUCCEEDED
  Installing github.com/go-delve/delve/cmd/dlv (/Users/gloriading/go/bin/dlv) SUCCEEDED
  Installing github.com/go-delve/delve/cmd/dlv (/Users/gloriading/go/bin/dlv) SUCCEEDED
  Installing github.com/go-delve/delve/cmd/dlv@master (/Users/gloriading/go/bin/dlv-dap) SUCCEEDED
  Installing github.com/go-delve/delve/cmd/dlv@master (/Users/gloriading/go/bin/dlv-dap) SUCCEEDED
  Installing honnef.co/go/tools/cmd/staticcheck (/Users/gloriading/go/bin/staticcheck) SUCCEEDED
  Installing honnef.co/go/tools/cmd/staticcheck (/Users/gloriading/go/bin/staticcheck) SUCCEEDED
  Installing golang.org/x/tools/gopls (/Users/gloriading/go/bin/gopls) SUCCEEDED

  All tools successfully installed. You are ready to Go :).
  Installing golang.org/x/tools/gopls (/Users/gloriading/go/bin/gopls) SUCCEEDED

  All tools successfully installed. You are ready to Go :).
  ```

* To run .go file: `go run .`
* More commands, use `go help`

* Find a module in `https://pkg.go.dev/`, include in .go file, then run `go mod tidy`

```
~/Project/golang/basics/hello$ go mod tidy
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c

// go.sum is automatically added -- this is for authenticating the module
```

