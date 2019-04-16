package main

import (
	"fmt"
)

var p = f()

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func main() {

	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	var x2, y2 int
	fmt.Println(&x2 == &x2, &x2 == &y2, &x2 == nil) // "true false false"

	fmt.Println(f() == f()) // "false"

	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)

}
