package main

import (
	"log"
	"net/http"
	"regexp"
)

// 使用正则表达式绑定url进行访问
// 参考：http://stackoverflow.com/questions/6564558/wildcards-in-the-pattern-for-http-handlefunc
type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

//全局变量不能使用:=
var browsePath string

func main() {
	//使用http.StripPrefix可启动简单文件服务器
	// http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("D:\\lqbz\\test\\go"))))
	// http.HandleFunc("/hello", HelloServer)
	// http.HandleFunc("/browse", fileHandler)
	// http.HandleFunc("/upload", uploadHandler)
	browsePath = "D:\\BaiduYunDownload"
	handler := &RegexpHandler{}
	handler.HandleFunc(regexp.MustCompile("/download/*"), downloadHandler)
	handler.HandleFunc(regexp.MustCompile("/hello"), HelloServer)
	handler.HandleFunc(regexp.MustCompile("/browse"), fileHandler)
	handler.HandleFunc(regexp.MustCompile("/upload"), uploadHandler)
	err := http.ListenAndServe(":80", handler)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
