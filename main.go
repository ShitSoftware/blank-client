package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Invalid usage!\nYou need to specify filename.")
		return
	}
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Print("Can't open file:", filename)
		return
	}

	var b *bytes.Buffer
	writer := multipart.NewWriter(b)
	p, err := writer.CreateFormFile("file", filepath.Base(f.Name()))
	if err != nil {
		fmt.Print("Unknown error occurred")
		return
	}

	io.Copy(p, f)
	writer.Close()

	rq, err := http.NewRequest("POST", "https://blank.maxunof.me", b)
	if err != nil {
		fmt.Print("Unknown error occurred")
		return
	}
	rq.Header.Add("Content-Type", writer.FormDataContentType())

	fmt.Print("Uploading...")

	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		fmt.Print("\rCan't connect with web server")
		return
	}

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print("\rUnknown error occurred")
		return
	}

	fmt.Print("\r" + string(c))
}
