package main

import (
	"fmt"
)

//包级别变量声明
var x1 int
var y1 string
var z1 bool

type intJx int //自定义变量声明

type person struct { //结构体类型声明
	name string
}

var person1 person //结构体类型的变量声明

var scale = 2 //变量声明，类型由初始化表达式推导

var q = [...]int{1, 2, 3} //数组的长度是根据初始化值的个数来计算
var count = [...]int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}

func main() {
	// var x = 42
	// var y = "James Bond"
	// var z = true

	intJx := 3                     //简短变量声明
	x, y, z := 42, "james,", false //简短变量声明,类型由初始化表达式推导 int, string, bool
	fmt.Println(x, y, z)

	fmt.Println(x1, y1, z1)        //包级别声明的变量会在main入口函数执行前完成初始化
	x1, y1, z1 = 45, "nihao", true //赋值
	fmt.Println(x1, y1, z1)

	fmt.Println(intJx)

	x = 1                // 变量的赋值
	p := &z              //一个指针对应变量在内存中的存储位置
	*p = true            // 通过指针间接赋值
	fmt.Println(p)       //变量的内存地址
	fmt.Println(z)       //变量值
	person1.name = "bob" // 结构体字段赋值

	count[x] = count[x] * scale // 数组、slice或map的元素赋值

	q1 := [...]int{1, 2, 3} //简短变量声明

	fmt.Println(x)
	fmt.Println(person1)
	fmt.Println(person1.name)
	fmt.Println(count)
	fmt.Println(q)
	fmt.Println(q1)

	// Print the indices and elements.
	for i, v := range count {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range count {
		fmt.Printf("%d\n", v)
	}

	type Currency int
	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
	for i, v := range symbol {
		fmt.Printf("%d %s\n", i, v)
	}

}
