package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func main() {
	names := []string{
		"hao",
		"wang",
		"ming",
	}
	// sort.Sort(names)
	// 自己实现的StringSlice类型
	// sort.Sort(StringSlice(names))
	// sort包自带的StringSlice类型
	sort.Sort(sort.StringSlice(names))
	// 直接调佣sort包自带的Strings函数进行排序
	// sort.Strings(names)
	fmt.Println(names)
}
