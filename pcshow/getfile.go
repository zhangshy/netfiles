package main

import (
	// "bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"netfiles/fileutil"
	"path/filepath"
)

type fileinfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func handlegetfiles(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	browsePath := r.Form["browsePath"][0]
	log.Println("browsePath:" + browsePath)
	fileinfos := make([]fileinfo, 0)
	filenames, _ := fileutil.GetCurrentDirFiles(browsePath)
	for _, filename := range filenames {
		fileinfos = append(fileinfos, fileinfo{Name: filename, Path: filepath.Join(browsePath, filename)})
	}
	fileinfoData, err := json.Marshal(fileinfos)
	if err != nil {
		log.Println("json init error")
	}
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Cache-Control", "no-cache")
	io.WriteString(w, string(fileinfoData))
}
