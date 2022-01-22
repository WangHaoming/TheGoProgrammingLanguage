package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasSuffix(url, "http") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// 为什么 _ 后面加 err的声明方式就产生冲突， err必须换名字
		// _, err := io.Copy(os.Stdout, resp.Body)
		// b, err := io.Copy(os.Stdout, resp.Body)

		_, err1 := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		// fmt.Println(b)
		fmt.Println("status code: ", resp.Status)
	}

}
