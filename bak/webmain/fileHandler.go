package main

import (
	"html/template"
	"net/http"
	"netfiles/fileutil"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	filenames, _ := fileutil.GetCurrentDirFiles(browsePath)
	// t, err := template.ParseFiles("template/html/browse.html")
	t := template.Must(template.ParseFiles("template/html/layout.html", "template/html/browse.html"))
	t.ExecuteTemplate(w, "layout", filenames)
}
