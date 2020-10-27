package main

import "fmt"

func main() {
	//a := "hello"
	//b := "world"
	//c := a + b
	//var x = len(a)
	//fmt.Println(c, x)
	d := "hello,世界!"
	for i := 0; i < len(d); i++ {
		fmt.Printf("字符：%c,ASCII:%d \n", d[i], d[i])
	}
	for i, v := range d {
		fmt.Println(i, v)
	}
}
