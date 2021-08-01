package main

import (
	"fmt"
	"sort"
)

func pickingNumbers(a []int32) int32 {

	if len(a) < 2 || len(a) > 100 {
		fmt.Errorf("2 <= n <= 100")
		return int32(0)
	}

	fmt.Println(a)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	fmt.Println(a)

	var count = [100]int{}

	for i := 0; i < len(a); i++ {
		count[a[i]]++
	}
	fmt.Println(count)

	var neighbour = []int{}
	for i := 0; i < len(count)-1; i++ {
		//if count[i] > 0 && count[i+1] > 0 {
		neighbour = append(neighbour, count[i]+count[i+1])
		//}
	}
	fmt.Println("neighbour ", neighbour)
	sort.Ints(neighbour)
	fmt.Println("neighbour", neighbour, "sorted")

	return int32(neighbour[len(neighbour)-1])
}

func main() {

	var a = []int32{4, 6, 5, 3, 3, 1}
	//var a = []int32{1, 2, 2, 3, 1, 2}
	//var a []int32
	//
	//for i := 0; i < 100; i++ {
	//	a = append(a, 66)
	//}

	result := pickingNumbers(a)

	fmt.Printf("%d\n", result)

}
