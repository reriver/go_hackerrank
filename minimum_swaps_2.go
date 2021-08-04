package main

import (
	"fmt"
)

func minimumSwaps(arr []int32) int32 {
	var countSwaps int32 = 0

	for i := int32(0); i < int32(len(arr)); i++ {
		minArr := arr[i]
		minI := i
		for j := i + 1; j < int32(len(arr)); j++ {
			if arr[j] < minArr {
				minArr = arr[j]
				minI = j
			}
		}
		if minI != i {
			arr[minI], arr[i] = arr[i], arr[minI]
			countSwaps++
		}
	}
	return countSwaps
}

func main() {
	var arr = []int32{7, 1, 3, 2, 4, 5, 6}
	res := minimumSwaps(arr)
	fmt.Printf("%d\n", res)
}
