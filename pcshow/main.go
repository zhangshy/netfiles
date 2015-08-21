package main

import "log"
import "net/http"

func main() {
	port := "80"
	http.Handle("/", http.FileServer(http.Dir("./static/html")))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	http.HandleFunc("/getfiles", handlegetfiles)
	http.HandleFunc("/upload", uploadHandler)
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
