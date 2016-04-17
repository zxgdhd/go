package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = strings.Join([]string{"http://", url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		} else {
			//useBuffer(resp)
			noUseBuffer(resp)
		}
	}
}

func useBuffer(resp *http.Response, url string) {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", body)
}

func noUseBuffer(resp *http.Response) {
	_, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	fmt.Printf("resp.State=%d\n", resp.Status)
	if err != nil {
		fmt.Println("Exception happened!")
	}
}
