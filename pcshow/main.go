package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
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
		// log.Println(iface.Name)
		addrs, err := iface.Addrs()
		if err != nil {
			return ips, nil
		}
		for _, addr := range addrs {
			// log.Println(addr.Network(), addr.String())
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// 只保留ipv4地址
			if ip.To4() != nil && ip.String() != "127.0.0.1" {
				ips = append(ips, ip.String())
			}
		}
	}
	return ips, nil
}

// 使用time.Timer进行延时操作
// 启动默认应用打开url
func open_url(port string, timer *time.Timer) {
	<-timer.C
	ips, _ := get_ips()
	log.Println("first ips is:" + ips[0])
	log.Println(runtime.GOOS)
	ips_url := "http://localhost:" + port
	switch runtime.GOOS {
	case "windows":
		log.Println("this is windows platform")
		exec.Command("cmd", "/C start ", ips_url).Run()
	case "linux":
		log.Println("this is linux platform")
		exec.Command("xdg-open", ips_url).Run()
	}
	log.Println("open_url end")
}

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
	go open_url(port, time.NewTimer(time.Second*3))
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
