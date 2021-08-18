package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func httpClientSend(image []byte, uri string, w http.ResponseWriter) {
	client := &http.Client{}
	println("send request: ", uri)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(image))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", body)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	println("imageHandler ....")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		println("error: ", err.Error())
		panic(err)
	}
	api := r.Header.Get("api")
	uri := ""
	if api == "go" {
		uri = "http://127.0.0.1:8086/api"
	} else {
		uri = "http://127.0.0.1:8087/api"
	}
	httpClientSend(body, uri, w)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	title, err := filepath.EvalSymlinks("." + r.URL.Path)
	if err != nil {
		println("error: ", err.Error())
	}
	types := map[string]string{
		".html": "text/html",
		".css":  "text/css",
		".js":   "application/javascript",
		".ico":  "image/vnd.microsoft.icon",
	}
	content, _ := loadFile(title)
	w.Header().Set("Content-Type", "text/html")
	for key, typ := range types {
		if strings.HasSuffix(title, key) {
			w.Header().Set("Content-Type", typ)
			break
		}
	}
	if content == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprintf(w, "%s", content)
	}
}

func loadFile(path string) ([]byte, error) {
	println("loading page: {}", path)
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	http.HandleFunc("/static/", staticHandler)
	http.HandleFunc("/api/hello", imageHandler)
	println("listen to 8085 ...")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
