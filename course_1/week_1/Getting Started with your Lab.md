# Getting Started with your Lab Sandbox
VSCode Sandbox comes with Go version 1.12.4 preinstalled.

## What extensions are already installed in this environment?
  - Gocode - a helper tool which is intended to be integrated with your source code editor
  - Gopkgs - a tool that provides list of available Go packages that can be imported. This is an alternative to `go list all`
  - GoOutline - a utility for extracting a JSON representation of the declarations in a Go source file. 
  - GoSymbols - a utility for extracting a JSON representation of the package symbols from a go source tree.
  - Guru - a tool for answering questions about Go source code.
  - GoRename - a command performs precise type-safe renaming of identifiers in Go source code.
  - Delve - a debugger for the Go programming language
  - GoCode - an autocompletion daemon for the Go programming language
  - GoReturns - adds zero-value return values to incomplete Go return statements, to save you time when writing Go.
  - GoLint - lints the Go source files named on its command line.
  
## How can I compile Go Code in my Terminal?
Open the terminal and navigate to the directory containing your file. For example, if your code file is HelloWorld.go, run the following command to compile.
```sh
go build HelloWorld.go
```
This will produce and output HelloWorld which is an executable. To run the file, execute the following command.

```sh
./HelloWorld
```

**To compile and run this file :**
1) Click on the ▶️ in the top right of the IDE to compile and run the go code. 

  More information on compiling code in VSCode can be found here:  https://code.visualstudio.com/docs/languages/go

## How can I Debug the code?
You can set break points by clicking to the left of line number where you want to break and investigate.
You can use the debugger in VSCode IDE using the icon on the left panel.