package main

import (
	"html/template"
	"log"
	"net/http"
	"netfiles/fileutil"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	filenames, _ := fileutil.GetCurrentDirFiles("D:\\lqbz\\test\\go")
	t, err := template.ParseFiles("template/html/browse.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, filenames)
}
