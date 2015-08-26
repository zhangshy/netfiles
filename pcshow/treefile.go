package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// type filetree struct {
// 	Id   string `json:"id"`
// 	Text string `json:"text"`
// }

type treestate struct {
	Opened   bool `json:"opened"`
	Disabled bool `json:"disabled"`
	Selected bool `json:"selected"`
}

type filetree struct {
	Id       string     `json:"id"`
	Text     string     `json:"text"`
	Icon     string     `json:"icon"`
	State    *treestate `json:"state"`
	Children []string   `json:"children"`
	Li_attr  string     `json:"li_attr"`
	A_attr   string     `json:"a_attr"`
}

func treeFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filepath := r.Form["id"][0]
	log.Println("filepath:" + filepath)
	filetrees := make([]filetree, 0)
	fileinfos, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Println("ReadDir error!" + filepath)
	}
	for _, info := range fileinfos {
		tree := filetree{}
		tree.Id = info.Name()
		tree.Text = info.Name()
		filetrees = append(filetrees, tree)
	}

	// filetreeData := "[ { \"text\" : \"Root node\", \"children\" : [ \"Child node 1\", \"Child node 2\" ] } ]"
	// io.WriteString(w, filetreeData)

	filetreeData, err := json.Marshal(filetrees)
	io.Copy(w, bytes.NewReader(filetreeData))
}
