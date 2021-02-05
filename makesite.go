// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

package main

import (
	"fmt"
	"io/ioutil"
)

// Storing Data
type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

// func main() {
// 	t := template.Must(template.New("template.tmpl").ParseFiles("new.html"))
// 	err = t.Execute(os.Stdout, ToDo.User)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func readFile() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	fmt.Print(string(fileContents))
}

func writeFile() {
	bytesToWrite := []byte("hello\ngo\n")
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	// Creates new template
// 	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
// 	err = t.Execute(os.Stdout, todos)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// Build Program
// $ go build

// Run latest Build
// ./makesite or $ ./(name of go file)
