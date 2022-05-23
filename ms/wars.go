package main

import "fmt"

func main() {
	n := 7
	c := 1
	for i := 1; i < n; i += 2 {
		c++
		fmt.Println("i =", i)
		fmt.Println("c =", c)
	}
	fmt.Println(c)

}
