package main

import "fmt"

func main() {
	a1 := [5]int{1, 2, 3, 4, 5}
	s1 := a1[2:4]
	fmt.Println(s1)
	fmt.Println("--------")
	s1[1] = 101
	fmt.Println(s1)
	fmt.Println(a1)
}
