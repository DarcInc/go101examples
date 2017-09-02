package main

import (
	"archive/zip"
	"bytes"
)

func main() {
	buffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buffer)
	defer zipWriter.Close()

	dataWriter, _ := zipWriter.Create("myfile.txt")

	dataWriter.Write([]byte("The end of the world as we know it."))

}
