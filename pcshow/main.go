package main

import (
	"log"
	"net/http"
	"os"
)

// 全局变量不能使用:=
var uploadFilePath string

func main() {
	port := "80"
	uploadFilePath = "./files"
	if os.MkdirAll(uploadFilePath, 0766) != nil {
		log.Println("create " + uploadFilePath + " error!")
		return
	}
	http.Handle("/", http.FileServer(http.Dir("./static/html")))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	http.Handle("/test/", http.StripPrefix("/test/", http.FileServer(http.Dir("./static/test"))))
	http.HandleFunc("/getfiles", handlegetfiles)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/tree_file", treeFileHandler)
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
