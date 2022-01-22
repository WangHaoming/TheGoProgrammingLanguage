package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// []byte("xx") 字符串强制转换成[]byte类型 []byte是一个slice类型
	c1 := sha256.Sum256([]byte("xx"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
