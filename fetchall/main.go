package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	switch {
	case strings.HasPrefix(url, "https://"):
	case strings.HasPrefix(url, "http://"):
	default:
		url = "https://" + url
	}

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("error: get(%q) -> %v", url, err)
		return
	}

	n, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("error: read(%q) -> %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%s\t%.2fs\t%dbytes", url, time.Since(start).Seconds(), n)
}
