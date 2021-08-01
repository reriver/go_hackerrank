package main

import (
	"fmt"
	"sort"
)

func CppPermutation(arr tInt32) int32 {

	return 0
}

type tInt32 []int32

func (f tInt32) Len() int {
	return len(f)
}

func (f tInt32) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f tInt32) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func main() {
	var z = tInt32{3, 2, 1}

	fmt.Println(z)
	sort.Sort(z)
	fmt.Println(z)

	CppPermutation(z)
}

//
//type tInt32 []int32
//
//func (f tInt32) Len() int {
//	return len(f)
//}
//
//func (f tInt32) Less(i, j int) bool {
//	return f[i] < f[j]
//}
//
//func (f tInt32) Swap(i, j int) {
//	f[i], f[j] = f[j], f[i]
//}
//
//func main() {
//
//	var f = tInt32 {2, 4, 1, 3}
//
//	fmt.Println(f)
//	sort.Ints(f)
//	fmt.Println(f)
//
//	CppPermutation(f)
//}
