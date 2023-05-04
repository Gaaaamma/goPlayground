# GO palyground
Learning GoLang

## Create a Go module
1. Start your module using the go mod init command
2. Make the module greetings.go with a certain package name
```shell
# At the directory of module
$ go mod init {module_path}

# Now we can make the module.go and edit it
$ touch greetings.go 
```
```go
package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}
```
Create a Go module: https://go.dev/doc/tutorial/create-module 

## Use the module
1. Under the same parent directory we can make another directory
``` shell
# <home>/
# |-- greetings/
# |-- hello/
 $ mkdir hello
 $ cd hello  # Now under hello directory
```
2. Enable dependency tracking for the code you're about to write.
```shell
$ go mod init {module_path}
```
3. Implement the code that using module created before
```go
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```
4. You might notice that we can't find example.com/greetings. That is because you haven't published the module yet, you need to adapt the example.com/hello module so it can **find the example.com/greetings code on your local file system.**
```shell
# Specifies that example.com/greetings should be replaced with ../greetings for the purpose of locating the dependency
$ go mod edit -replace example.com/greetings=../greetings
# Run go mod tidy command to synchronize the example.com/hello module's dependencies
$ go mod tidy
```
5. Now it works
```shell
$ go run .
```
Use the module : https://go.dev/doc/tutorial/call-module-code <br>
### NOTICE: YOU MIGHT GET SOME ERROR MESSAGE IN VSCODE gopls
This is because there are multiple modules in your workspace (Though we can build and run the code)
```
gopls requires a module at the root of your workspace.
You can work with multiple modules by opening each one as a workspace folder.
Improvements to this workflow will be coming soon (https://github.com/golang/go/issues/32394),
and you can learn more here: https://github.com/golang/go/issues/36899.
```
- The setting in setting.json is deprecated now.
```
"gopls": {
    "experimentalWorkspaceModule": true,
}
```
- Go workspaces is preferred now(Go 1.18+)
### gopls Multiple-modules
Gopls has several alternatives for working on multiple modules simultaneously, described below. Starting with Go 1.18, Go workspaces are the preferred solution.<br>
- The easiest way to work on multiple modules in Go 1.18+ is to create a **go.work** file containing the modules you wish to work on
- Then set your workspace root to the directory containing the **go.work** file
```shell
cd $WORK
go work init # Create a go.work file using go work init
go work use ./tools/ ./tools/gopls/ # go work use MODULE_DIRECTORIES to add directories containing go.mod files to the workspace
```
Now, the problem might be solved<br>
Ref: https://github.com/golang/tools/blob/master/gopls/doc/workspace.md