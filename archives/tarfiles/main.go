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
	"archive/tar"
	"io"
	"os"
)

/*
 * This function creates an archive and writes a file
 * into the archive.
 */
func createArchive() {
	someData := []byte("The quick brown fox jumps over the lazy dog")

	file, _ := os.OpenFile("temp.tar", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()

	writer := tar.NewWriter(file)
	header := &tar.Header{Name: "MyFile.txt", Mode: 0644}
	writer.WriteHeader(header)
	writer.Write(someData)
}

/*
 * This function opens an archive for reading, reads the
 * header to get the file information, then copys the
 * data to a new file.
 */
func readArchive() {
	file, _ := os.Open("temp.tar")
	defer file.Close()

	reader := tar.NewReader(file)
	header, _ := reader.Next()

	outfile, _ := os.OpenFile(header.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
	defer outfile.Close()

	io.Copy(outfile, reader)
}

func main() {
	createArchive()
	readArchive()
}
