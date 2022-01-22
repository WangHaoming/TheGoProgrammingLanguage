package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee = Employee{
		ID: 123,
	}
	fmt.Println(EmployeeByID(123))
	// 如果将EmployeeByID函数的返回值从*Employee指针类型改为Employee值类型，那么更新语句将不能编译通过，因为在赋值语句的左边并不确定是一个变量（译注：调用函数返回的是值，并不是一个可取地址的变量）。
	// EmployeeByID(123).Salary = 0
	x := EmployeeByID(123)
	x.Salary = 0
	dilbert.Salary = 0
}

func EmployeeByID(id int) Employee {
	return Employee{
		ID: 123,
	}
}
