package main

import (
	"flag"
	"fmt"
	"net/http"
)

var store *URLStore

var (
	listenAddr = flag.String("http", ":3000", "http listen address")
	dataFile   = flag.String("file", "store.gob", "data store file name")
	hostname   = flag.String("host", "localhost:3000", "host name and port")
)

func main() {
	flag.Parse()
	fmt.Println("Server listening at", *hostname)
	store = NewURLStore(*dataFile)
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(*listenAddr, nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.Redirect(w, r, "/add", http.StatusFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprint(w, addForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "%s", key)
}

const addForm = `
<html><body>
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
</html></body>
`
