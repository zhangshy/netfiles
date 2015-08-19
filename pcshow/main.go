package main

import "log"
import "net/http"

func main() {
	port := "80"
	http.Handle("/", http.FileServer(http.Dir("./template/html")))
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
