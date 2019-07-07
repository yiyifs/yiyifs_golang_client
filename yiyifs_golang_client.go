package main

import (
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	fileName := "ubuntu-16.04.6-desktop-amd64.iso"
	go func() {
		defer w.Close()
		defer m.Close()
		part, err := m.CreateFormFile(fileName, fileName)
		if err != nil {
			return
		}
		file, err := os.Open("/Users/jac/Downloads/" + fileName)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()
	rep, err := http.Post("http://localhost:9007/api/up?fn="+fileName, m.FormDataContentType(), r)
	if err != nil {
		log.Println(err.Error())
	}
	res, _ := ioutil.ReadAll(rep.Body)
	log.Println("finish", string(res))
}
