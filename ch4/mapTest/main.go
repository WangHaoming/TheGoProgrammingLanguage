package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{}
	// 和 map[string]int{}一样会生成一个空map
	var nilAges map[string]int
	fmt.Println(ages)
	fmt.Println(ages == nil)
	// will get an error -- nilAges["alice"] = 32
	// nilAges["alice"] = 32
	fmt.Println(nilAges)
	fmt.Println(nilAges == nil)
	nilAges = make(map[string]int)
	nilAges["alice"] = 32
	fmt.Println(nilAges["bob"])
	_, found := nilAges["bob"]
	fmt.Println(found)

	// _ = &ages["bob"] // compile error: cannot take address of map element

	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42, "A": 0}))

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
