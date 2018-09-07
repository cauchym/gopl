package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// [TODO]途中でErr発生したらそれ以降のurlに関してfetchできないのだめっぽい
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "status: %s\nfetch: reading %s: %v\n", resp.Status, url, err)
			os.Exit(1)
		}
		fmt.Printf("status: %s\nbody: %s\n", resp.Status, b)
	}
}
