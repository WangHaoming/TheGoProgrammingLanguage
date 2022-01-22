package main

import "fmt"

func main() {
	s := "には尾\a"
	// s = "124"
	s = s[3:]
	// cannot assign to s[0] (strings are immutable)
	// s[0] = "L"
	s = "\xe4\xb8\x96\xe7\x95\x8c"
	s = "\u4e16\u754c"
	// s = "\U00004e16\U0000754c"
	fmt.Println(len(s))
	fmt.Println(s)
}
