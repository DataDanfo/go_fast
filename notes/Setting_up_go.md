# Setting up your Go envirnment
Just like all other programming languages, Go also require setting up and environment where you can run your program without having to many difficulties to face.
If you havn't already installed Go on your machine or whatever you'll be using to run Go on, don't worry, this tutorial will guide you through each and every steps you'll needing to install Go and its supporting tools.
### Installing Go 
To install Go on your machine. Firstly, you'll need to download and install the Go development tool from the official [Go website](https://go.dev/dl). You'll make sure you're download the most recent version of Go, to get all the recents update and updated packages.
Choose the download for your platform and install. The .pkg installer for developers with Mac and the .msi installer for developers with Windows which will automatically install Go to the correct location. Ensure there are not older versions of Go installed

*If your a Mac developer, your can using [Homebrew]() with the command ``brew install go``. For windows developers who use [Chocolatel]() can install Go with the command ``choco install golang``*

 Most of Linux and BSD installers are gzipped TAR files and expand to a directory named 'go'. Copy this directory to /usr/local and add /usr/local/go/bin to your $PATH so that the go command is accessible:


``$ tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz``
``$ echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bash_profile``
``$ source $HOME/.bash_profile``
You might need root permissions to write to /usr/local. 
If the ``tar`` command fails, rerun it with ``sudo tar ``
``-C /usr/local -xzf ``
``go1.20.5.linux-amd64.tar.gz``

Go programs compiled to a single native binary and don not require any additional software to be installed in order to run them, compared to other languages like Java, Python, and JavaScript, which require you to install a virtual machine to run your program.  Using a single native binary makes it a lot easier to distribute programs written in Go.
You can check if the environment is setup correctly by opening your ternimal or command prompt and typing
`` $ go version ``
if everything is done correctly, you should see something like this:
``go version go1.20.5. darwin/arm64``
This tells yuo that the Go version installed is 1.20.5 on macOS. (*Darwin is the operating system at the heart of macOS, and arm64 is the name for the 64-bit chips based on ARM's design.*)
On x64 Linux, you should see soomethings like:
``go version go1.20.5 linux/arm64``
## Troubleshooting your Go installation 
When checking to see if everything was correctly installed and you're getting an error message, it's likely that you don't have Go in your executable path or you have a different program named Go in your path. If this is the case, on macOS and othe Unix-like systems,use ``which go`` to see the Go command being executed, if any. if nothing was retured, you'll need to fix your executable path. If you're on Linux or BSD, its possible the you installed the 64-bit Go development tools on a 32-bit system or the development tools for the wrong chip architecture.
## Go Tooling
In Go, all the devlopments tools are accessed through their Go counterpart. Like the ``go version`` command, there are others.
-Complier (``go build``)
-code formatter (``go fmt``)
-Dependency manager (``go mod``)
-Test runner (``go test``)
-A tool to scan common coding mistakes (``go vet``) etc.

## Your first Go program
Just like every other programming language, our first program is always a 'Hello World!' program.
*a program that simply outputs Hello World.* Let write one using Go.

### Making a Go module
Firstly,  we'll need to create a directory that will store or hold our program.

