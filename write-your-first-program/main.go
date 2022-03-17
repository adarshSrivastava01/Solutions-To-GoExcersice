// Create an OSX executable: GOOS=darwin GOARCH=386 go build

// Create a Windows executable: GOOS=windows GOARCH=386 go build

// Create a Linux executable: GOOS=linux GOARCH=arm GOARM=7 go build

package main

import "fmt"

func main() {
	fmt.Println("Adarsh")
	fmt.Println("Aashish")

	fmt.Print("Only Print", "Hey", "Hello")
	fmt.Println("Hey", "Hello", "Bye")
	// Println() replaces comma with "<space>" while Print() doesn't
}

// if you want to check your $GOPATH run <go env GOPATH>
// It's a file for Go runtime
// Stores Go source code files and compiled packages
// It's a path for gophers to follow

// You should save your files under
// Under $GOPATH/src
