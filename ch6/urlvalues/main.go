package main

import (
	"fmt"

	"example.com/url"
)

func main() {
	// m := url.Values{} // direct construction
	m := url.Values(nil)
	dict := make(map[string]int)
	fmt.Printf("%T\n", dict)
	fmt.Printf("%T\n", m)
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // "[1 2]"  (direct map access)

	m = nil
	fmt.Println(m.Get("item")) // ""
	// m.Add("item", "3")
}
