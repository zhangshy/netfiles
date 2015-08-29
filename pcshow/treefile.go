package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type treestate struct {
	Opened   bool `json:"opened"`
	Disabled bool `json:"disabled"`
	Selected bool `json:"selected"`
}

// 1. 使用,omitempty如果是空的话在生成json时就不生成
// http://stackoverflow.com/questions/17306358/golang-removing-fields-from-struct-or-hiding-them-in-json-response
// 2. 使用map[string]string生成nested json
type filetree struct {
	Id       string            `json:"id"`
	Text     string            `json:"text"`
	Icon     string            `json:"icon,omitempty"`
	State    *treestate        `json:"state,omitempty"`
	Children bool              `json:"children,omitempty"`
	Li_attr  map[string]string `json:"li_attr,omitempty"`
	A_attr   map[string]string `json:"a_attr,omitempty"`
}

func treeFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dirpath := r.Form["id"][0]
	log.Println("dirpath:" + dirpath)
	if strings.HasPrefix(dirpath, "C:") {
		log.Println("Windows C do not browse!")
		io.WriteString(w, "[{\"text\" : \"Do not use C:\"}]")
		return
	}
	filetrees := make([]filetree, 0)
	fileinfos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Println("ReadDir error!" + dirpath)
	}
	for _, info := range fileinfos {
		tree := filetree{}
		// 使用文件路径作为id
		tree.Id = filepath.Join(dirpath, info.Name())
		tree.Text = info.Name()
		if info.IsDir() {
			// 是文件夹
			tree.Icon = "jstree-folder"
			tree.Children = true
		} else {
			// Use the rel attribute
			// http://stackoverflow.com/questions/4899520/jstree-types-plugin-does-not-display-custom-icons
			// tree.Li_attr = map[string]string{"rel": "file"}
			tree.Icon = "jstree-file"
			// 设置文件的href
			tree.A_attr = map[string]string{"href": "/download?file=" + tree.Id}
		}

		// 如果文件名为test的话，就选中这个文件
		if info.Name() == "test" {
			state := &treestate{Selected: true}
			tree.State = state
		}
		filetrees = append(filetrees, tree)
	}

	// filetreeData := "[ { \"text\" : \"Root node\", \"children\" : [ \"Child node 1\", \"Child node 2\" ] } ]"
	// io.WriteString(w, filetreeData)

	filetreeData, err := json.Marshal(filetrees)
	io.Copy(w, bytes.NewReader(filetreeData))
}
