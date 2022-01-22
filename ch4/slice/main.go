package main

import "fmt"

var months = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var sliceString = "hello world!"

func main() {
	Q2 := months[4:7]
	summer := months[6:9]
	summer[0] = 32
	fmt.Println(Q2) // ["April" "May" "June"]
	fmt.Println(summer)
	// fmt.Println(summer[:20])
	endlessSummer := summer[:5]                // extend a slice (within capacity)
	fmt.Println(endlessSummer)                 // "[June July August September October]"
	fmt.Println("gap of summer ", cap(summer)) //
	fmt.Println("gap of Q2 ", cap(Q2))
	substring := sliceString[3:]
	var nilSlice []int
	emptySlice := make([]int, 0)

	emptyArray := [...]int{}

	// go语言字符串不可更改 所以下面的语句无法编译通过
	// substring[2] = "!"
	fmt.Println(substring)
	fmt.Println(nilSlice)
	fmt.Println(nilSlice == nil)

	fmt.Println(emptySlice)
	fmt.Println(emptySlice == nil)

	fmt.Println(emptyArray)
	// 数组不可以和nil比较？ 数组不可能等于nil，因为声明数组的时候一定会给数组分配一块内存
	// slice是一个引用类型可以有nil，也就相当于空指针。而array不能不指向任何内存
	// fmt.Println(emptyArray == nil)
}
