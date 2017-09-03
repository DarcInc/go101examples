/*
BSD 2-Clause License

Copyright (c) 2017, Darc Inc
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func createZip() {
	zipFile, _ := os.OpenFile("lorem.zip", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	inputFile, _ := os.Open("lorem.txt")
	defer inputFile.Close()

	dataWriter, _ := zipWriter.Create("lorem.txt")
	io.Copy(dataWriter, inputFile)

	zipWriter.Flush()
}

func unpackZip() {
	zipFile, _ := os.Open("lorem.zip")
	defer zipFile.Close()

	fileinfo, _ := os.Stat("lorem.zip")
	zipReader, _ := zip.NewReader(zipFile, fileinfo.Size())

	for _, z := range zipReader.File {
		file, _ := z.Open()
		defer file.Close()

		io.Copy(os.Stdout, file)
	}
}

func fileSizes() {
	origInfo, _ := os.Stat("lorem.txt")
	zipInfo, _ := os.Stat("lorem.zip")

	fmt.Println()
	fmt.Println()
	fmt.Printf("Original %d bytes but zipped %d bytes\n", origInfo.Size(), zipInfo.Size())
}

func main() {
	createZip()
	unpackZip()
	fileSizes()
}
