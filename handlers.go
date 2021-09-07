package main

import (
	"fmt"
	"net/http"
)

func declareHandlers() {
	http.HandleFunc("/", fetchFile)
}

func fetchFile(w http.ResponseWriter, r *http.Request) {
	p := "htdocs" + r.URL.Path
	if p == "htdocs/" {
		p = "htdocs/index.htm"
	}
	t, e := readFile(p)
	if e != nil || t == nil {
		http.Error(w, "Page was not found in htdocs.\nIf you own this server, please check htdocs.", 404)
		return
	}
	fmt.Fprintf(w, "%s", *t)
}
