package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, err := template.ParseFiles("template/html/upload.html")
	if err != nil {
		log.Println(err)
	}
	if r.Method == "GET" {
		t.Execute(w, nil)
	} else {
		file, handle, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
		}
		f, err := os.OpenFile(filepath.Join("files", handle.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		defer file.Close()
		io.Copy(f, file)
		t.Execute(w, "上传成功")
	}
}
