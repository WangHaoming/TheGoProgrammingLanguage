package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(b []byte) (n int, err error) {
	*c += ByteCounter(len(b))
	return len(b), nil
}
func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
