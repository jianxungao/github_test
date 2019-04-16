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

func bitshift(v *int) int {
	return *v << 1
}

func bitshiftu(v int) int {
	return v << 2
}

func main() {

	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Printf("%x\t%b\t%d\n", x, x, x)

	c := bitshiftu(x)
	fmt.Printf("%x\t%b\t%d\n", c, c, c)
	fmt.Println("now x is ", x)

	b := bitshift(p)
	fmt.Printf("%x\t%b\t%d\n", b, b, b)
	fmt.Println("now x is ", x)

	d := *p << 1
	fmt.Printf("%x\t%b\t%d\n", d, d, d)
	fmt.Println("now x is ", x)

	var x2, y2 int
	fmt.Println(&x2 == &x2, &x2 == &y2, &x2 == nil) // "true false false"

	fmt.Println(f() == f()) // "false"

	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)

}
