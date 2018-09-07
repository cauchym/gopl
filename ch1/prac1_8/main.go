package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urlList := os.Args[1:]
	fetchWithStream(os.Stdout, urlList)
}

func fetchWithStream(w io.Writer, urlList []string) {
	for _, url := range urlList[0:] {
		if (!strings.HasPrefix(url, "http://")) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, errc := io.Copy(w, resp.Body)
		if errc != nil {
			fmt.Fprintf(os.Stderr, "copy: %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Printf("%s", b)
	}
}
