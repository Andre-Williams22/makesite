// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	// "string"
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

func main() {
	// t := template.Must(template.New("template.tmpl").ParseFiles("new.html"))
	// err = t.Execute(os.Stdout, ToDo.User)
	// if err != nil {
	// 	panic(err)
	// }

	fileFlag := flag.String("file", "first-post.txt", "define input text")
	dirFlag := flag.String("directory", "none", "generates all .txt files in directory")
	outputDirFlag := flag.String("output", "templates/", "Generator output directory")
	flag.Parse()

	if *dirFlag == "none" {
		runFile(*fileFlag, "txt_dir/")
	} else {
		runDir(*dirFlag, *outputDirFlag)
	}
}

func runFile(fileFlag, directory string) {

	var fileName string = fileFlag

	if fileName[strings.Index(fileFlag, "."):len(fileFlag)] != ".txt" {
		return
	}

	if strings.Contains(strings.ToLower(fileFlag), ".md") {

		var data string = readFile(directory + fileFlag)
		tmpl := renderTemplate("template.tmpl", data, fileName)
		output := blackfriday.Run(tmpl)
		ioutil.WriteFile(output, fileFlag)

		return
	}

	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".html"

	var data string = readFile(directory + fileFlag)
	renderTemplate("template.tmpl", data, fileName)
}

func readFile(file string) string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	fmt.Print(string(fileContents))
	return string(fileContents)
}

func writeFile(file string, content string) {
	bytesToWrite := []byte("hello\ngo\n")
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func saveFile() {
	var filename string
	var dir string

	flag.StringVar(&filename, "file", "", "File name")
	flag.StringVar(&dir, "dir", "", "Directory name")
	flag.Parse()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name()[len(f.Name())-4:] == ".txt" {
			content := readFile(dir + "/" + f.Name())
			writeFile(f.Name(), content)
		}
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

// go build -o main
// ./main
// go run main.go
