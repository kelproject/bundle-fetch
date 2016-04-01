package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string) error {
	var err error
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			io.Copy(os.Stdout, resp.Body)
			return nil
		}
		err = fmt.Errorf("%s returned %s", url, resp.Status)
	}
	return err
}

func main() {
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		fmt.Fprintln(os.Stderr, "bundle-fetch: missing URL")
	}
	backoff := time.Duration(1 * time.Second)
	for {
		err := fetch(url)
		if err == nil {
			break
		}
		fmt.Fprintf(os.Stderr, "bundle-fetch [ERROR] %s\n", err)
		time.Sleep(backoff)
		if backoff < (16 * time.Second) {
			backoff = backoff * 2
		}
	}
}
