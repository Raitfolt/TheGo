package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

func hasHttpPrefix(s string) bool {
	return strings.HasPrefix(s, httpPrefix)
}

func hasHttpsPrefix(s string) bool {
	return strings.HasPrefix(s, httpsPrefix)
}

func main() {
	for _, url := range os.Args[1:] {

		if !hasHttpPrefix(url) && !hasHttpsPrefix(url) {
			url = "http://" + url
			fmt.Println("URL was complited with prefix \"http://\"")
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		code := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy to os.Stdout %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Response status: %s", code)
	}
}
