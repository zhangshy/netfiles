package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 截取文件名称
	filename := r.RequestURI[len("/download/"):]
	// 使用QueryUnescape解决中文路径乱码的问题
	// 参考：http://stackoverflow.com/questions/13826808/recommended-way-to-encode-decode-urls
	filename, err := url.QueryUnescape(filename)
	if err != nil {
		log.Println(err)
	}
	file, err := os.OpenFile(filepath.Join(browsePath, filename), os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
		return
	}
	// 设置回复头，告诉浏览器下载文件
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	// 使用io.copy
	io.Copy(w, file)
}
