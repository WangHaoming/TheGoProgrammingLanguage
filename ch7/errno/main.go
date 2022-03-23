package main

import (
	"fmt"
	"syscall"
)

// type Errno unitptr
// type unitptr int

// var errors = [...]string{
// 	1: "operation not permitted",       // EPERM
// 	2: "no such file or directory ！！！", // ENOENT
// 	3: "no such process",               // ESRCH
// }

// func (e Errno) Error() string {
// 	if 0 <= int(e) && int(e) < len(errors) {
// 		return errors[e]
// 	}
// 	return fmt.Sprintf("errno %d", e)
// }

func main() {
	// syscall.Errno(2)是一个强制转型，把int转换为Errno类型
	var err error = syscall.Errno(2)
	fmt.Printf("%T\n", err)
	fmt.Println(err.Error())
	fmt.Println(err)
}
