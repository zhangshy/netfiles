package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("download enter")
	r.ParseForm()
	filename := r.Form["file"][0]
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
		return
	}
	// 获取文件名
	_, name := filepath.Split(filename)
	log.Println("file name:" + name)
	// 设置回复头，告诉浏览器下载文件，为文件名添加双引号解决文件文件名包含空格的情况
	w.Header().Set("Content-Disposition", "attachment; filename=\""+name+"\"")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	// 使用io.copy
	io.Copy(w, file)
}
