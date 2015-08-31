package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("upload enter")

	r.ParseForm()
	file, handle, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	f, err := os.OpenFile(filepath.Join(uploadFilePath, handle.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	defer file.Close()
	io.Copy(f, file)
}
