package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:4]
	fmt.Println(len(s), cap(s))
	newS := append(s, 55, 66)
	fmt.Printf("len=%d cap=%d\n", len(newS), cap(newS))
	s = nil
	fmt.Println(len(s), cap(s))
}
