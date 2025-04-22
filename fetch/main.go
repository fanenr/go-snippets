package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fetch2()
}

func fetch1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		data, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Printf("%s\n\n", data)
	}
}

func fetch2() {
	for _, url := range os.Args[1:] {
		switch {
		case strings.HasPrefix(url, "https://"):
		case strings.HasPrefix(url, "http://"):
		default:
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s %d: %v\n", url, resp.StatusCode, err)
			continue
		}

		var sb strings.Builder
		_, err = io.Copy(&sb, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}

		fmt.Printf("%s\n\n", sb.String())
	}
}
