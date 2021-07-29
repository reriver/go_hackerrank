package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	var sum int32

	sum = 0
	fmt.Println(ar)
	for i := 0; i < len(ar); i++ {
		sum += int32(ar[i])
	}
	return sum
}

func main() {
	fmt.Println("bjhb")
	ar := []int32{1, 2, 3, 4, 10, 11}
	fmt.Println(len(ar))
	fmt.Println(simpleArraySum(ar))
}
