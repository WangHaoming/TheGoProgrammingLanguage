package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf *bytes.Buffer
	var w io.Writer
	// io.Writer是一个接口，所以打印出来的类型是nil， bytes.Buffer是一个struct所以能打印出实际类型
	fmt.Printf("%T\n", buf)
	fmt.Printf("%T\n", w)
	// if debug {
	// 	buf = new(bytes.Buffer)
	// }
	// f(buf)
	// if debug {
	// }
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
