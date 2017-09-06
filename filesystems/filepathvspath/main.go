package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	path := filepath.Join("foo", "bar")
	fmt.Println(path)

	path, _ = filepath.Abs(path)
	fmt.Println(path)

	path = filepath.Join("foo/bar", "baz")
	fmt.Println(path)

	// Spin through the files, print the absolute path,
	// the file's directory, the base name, and the
	// file extension.
	files, _ := ioutil.ReadDir(".")
	for _, f := range files {
		path, _ := filepath.Abs(f.Name())
		fmt.Println(path)

		dir := filepath.Dir(path)
		fmt.Println(dir)

		base := filepath.Base(path)
		fmt.Println(base)

		ext := filepath.Ext(path)
		fmt.Println(ext)
	}

	matches, _ := filepath.Glob("*.go")
	fmt.Printf("%v\n", matches)
}
