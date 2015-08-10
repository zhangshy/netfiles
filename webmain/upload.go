package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("template\\html\\upload.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, nil)
	} else {
		file, handle, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
		}
		f, err := os.OpenFile("files\\"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		defer file.Close()
		io.Copy(f, file)
		io.WriteString(w, "上传成功")
	}
}
