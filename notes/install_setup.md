# Setups for Golang

Go is a tool for managing Go source code.
when you install go, it comes automatically in your PATH

```bash
flag provided but not defined: -version
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.



```
## Go help
```bash

go help <topic>" // prints help for more information about that topic.

```

## Go Version

 go --version

## Go Module

go module is collection of go packages that can be installed together.

```bash

go mod init  somename_of_module
```

```go

module github.com/DataDanfo/go_fast/variables

go 1.21.5




```

## Go Package

go package is a collection of go files that can be compiled together

# filename main.go

```go

package main

import "fmt"  // Standard Library Package


func main(){
    
    fmt.Println("hello world")
}


```

# go build

```bash
go build main.go // compiles the project to bytecode
// go build -o execulatled_file main.go // creates an executable file for you
// go build .


```

## go run

```bash
go run . // runs the project no need to build

./ executled_file // runs the executable file if already build


```
