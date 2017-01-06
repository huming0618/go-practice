package main

import (
	"fmt"
	"time"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	for _, url := range os.Args[1:]{
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		secs := time.Since(start).Seconds()
		fmt.Printf("%.2fs %7d  %s\n", secs, nbytes, url)
		//fmt.Printf("%s", b)

	}
}

