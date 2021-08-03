package main

import (
	"fmt"
	"math"
)

func hourglass(arr [][]int32, x int, y int) int32 {
	//var res [][]int32
	//var row1 = []int32 {arr[x-1][y-1], arr[x+0][y-1], arr[x+1][y-1]}
	//var row2 = []int32 {            0, arr[x+0][y+0],             0}
	//var row3 = []int32 {arr[x-1][y+1], arr[x+0][y+1], arr[x+1][y+1]}
	//
	//res = append(res, row1)
	//res = append(res, row2)
	//res = append(res, row3)
	res := arr[y-1][x-1] + arr[y-1][x] + arr[y-1][x+1] + arr[y][x] + arr[y+1][x-1] + arr[y+1][x] + arr[y+1][x+1]

	fmt.Println(res)

	return res

}

func hourglassSum(arr [][]int32) int32 {
	fmt.Println(arr)
	max := int32(math.MinInt32)
	curMax := int32(0)

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			curMax = hourglass(arr, i, j)
			if max < curMax {
				max = curMax
			}
		}
	}
	return max
}

func main() {
	var arr [][]int32
	//var row1 = []int32 {1, 1, 1, 0, 0, 0}
	//var row2 = []int32 {0, 1, 0, 0, 0, 0}
	//var row3 = []int32 {1, 1, 1, 0, 0, 0}
	//var row4 = []int32 {0, 0, 2, 4, 4, 0}
	//var row5 = []int32 {0, 0, 0, 2, 0, 0}
	//var row6 = []int32 {0, 0, 1, 2, 4, 0}
	var row1 = []int32{-9, -9, -9, 1, 1, 1}
	var row2 = []int32{0, -9, 0, 4, 3, 2}
	var row3 = []int32{-9, -9, -9, 1, 2, 3}
	var row4 = []int32{0, 0, 8, 6, 6, 0}
	var row5 = []int32{0, 0, 0, -2, 0, 0}
	var row6 = []int32{0, 0, 1, 2, 4, 0}

	arr = append(arr, row1)
	arr = append(arr, row2)
	arr = append(arr, row3)
	arr = append(arr, row4)
	arr = append(arr, row5)
	arr = append(arr, row6)

	fmt.Println(arr)

	result := hourglassSum(arr)

	fmt.Printf("%d\n", result)
}
