package main

import "fmt"

type IntSet struct {
	body string
}

func (i *IntSet) String() string {
	fmt.Println((*i).body)
	i.body = "changed!"
	return i.body
}
func main() {
	s := IntSet{body: "Hello pointer!"}
	fmt.Println(s.String())
	fmt.Println(s.body)

	// compile error: String requires *IntSet receiver
	// var _ = IntSet{body: "Hello pointer!"}.String()

	// IntSet does not implement fmt.Stringer (String method has pointer receiver)
	// var _ fmt.Stringer = s

	var _ fmt.Stringer = &s
}
