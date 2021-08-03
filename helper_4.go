package main

import (
	"fmt"
	"sort"
)

func AllPermutation(a []int) {
	var res [][]int
	calPermutation(a, &res, 0)
	fmt.Println(res)
}
func calPermutation(arr []int, res *[][]int, k int) {
	for i := k; i < len(arr); i++ {
		swapInt(arr, i, k)
		calPermutation(arr, res, k+1)
		swapInt(arr, k, i)
	}
	if k == len(arr)-1 {
		r := make([]int, len(arr))
		copy(r, arr)
		*res = append(*res, r)
		return
	}
}
func swapInt(arr []int, i, k int) {
	arr[i], arr[k] = arr[k], arr[i]
}

func main() {
	var a = []int{2, 1, 3}

	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)

	AllPermutation(a)

	//fmt.Printf("%d\n", result)
}
