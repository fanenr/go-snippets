package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mtx sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	count++
	mtx.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	c := count
	mtx.Unlock()
	fmt.Fprintf(w, "count: %d\n", c)
}
