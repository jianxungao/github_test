package main

import (
	"fmt"
)

func main() {

	s := "世界，hello！"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#X\t", s[i])
		fmt.Printf("%b\t", s[i])
		fmt.Printf("%#U\n", s[i])
	}
	//s = "hello"
	bs := []byte(s)
	fmt.Println(bs)

	for i, v := range s {
		fmt.Println(i, "\t", v)
	}
}
