package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//struct中的各项首字母要大写，否则渲染时用"."获取不到
type Info struct {
	Title string
	Ret   string
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmpl := make(map[string]*template.Template)
	//渲染多个模板文件
	tmpl["upload"] = template.Must(template.ParseFiles("template/html/layout.html", "template/html/upload.html"))
	info := Info{Title: "上传测试", Ret: ""}
	if r.Method == "GET" {
		tmpl["upload"].ExecuteTemplate(w, "layout", info)
	} else {
		file, handle, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
		}
		f, err := os.OpenFile(filepath.Join("files", handle.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		defer file.Close()
		io.Copy(f, file)
		info.Ret = "上传成功"
		tmpl["upload"].ExecuteTemplate(w, "layout", info)
	}
}
