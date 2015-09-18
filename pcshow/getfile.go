package main

import (
	// "bytes"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"netfiles/fileutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

type fileinfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ipInfo struct {
	Ip string
}

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
	log.Println(runtime.GOOS)
	ips_url := "http://localhost:" + port + "/ips"
	switch runtime.GOOS {
	case "windows":
		log.Println("this is windows platform")
		exec.Command("cmd", "/C start "+ips_url).Run()
	case "linux":
		log.Println("this is linux platform")
		exec.Command("xdg-open", ips_url).Run()
	}
	log.Println("open_url end")
}

func handleips(w http.ResponseWriter, r *http.Request) {
	ips, err := get_ips()
	if err != nil {
		log.Println(err)
		io.WriteString(w, "error")
		return
	}
	log.Println("first ips is:" + ips[0])
	t, err := template.ParseFiles("static/template/ips.html")
	if err != nil {
		log.Println(err)
		io.WriteString(w, "error")
		return
	}
	ipinfo := ipInfo{Ip: ips[0]}
	t.Execute(w, ipinfo)
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
