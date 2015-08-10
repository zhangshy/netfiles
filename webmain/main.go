package main

import (
	"log"
	"net/http"
)

func main() {
	//使用http.StripPrefix可启动简单文件服务器
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("D:\\lqbz\\test\\go"))))
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/browse", fileHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
