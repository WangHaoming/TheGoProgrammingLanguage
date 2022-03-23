package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	rw := w.(io.ReadWriter)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", rw)
}
