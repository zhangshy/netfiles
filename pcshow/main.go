package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

// 全局变量不能使用:=
var uploadFilePath string

func get_ips() ([]string, error) {
	ips := make([]string, 0)
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return ips, nil
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return ips, nil
		}
		for _, addr := range addrs {
			ips = append(ips, addr.String())
		}
	}
	return ips, nil
}

func main() {
	port := "80"
	uploadFilePath = "./files"
	if os.MkdirAll(uploadFilePath, 0766) != nil {
		log.Println("create " + uploadFilePath + " error!")
		return
	}
	ips, _ := get_ips()
	log.Println("first ip is:" + ips[0])
	switch runtime.GOOS {
	case "windows":
		log.Println("this is windows")
		exec.Command("cmd", "/C start http://localhost").Run()

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
