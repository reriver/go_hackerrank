package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func miniMaxSum(arr []int32) {
	var sum int64 = 0
	//var sums []int32
	var max int32 = -2147483648
	var min int32 = 2147483647

	fmt.Println(arr)
	for i := range arr {
		sum += int64(arr[i])
		if arr[i] < min {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	//fmt.Println(sum, min, max)
	//for i := range arr {
	//	sums = append(sums, sum-int32(arr[i]))
	//}
	fmt.Println(sum-int64(max), sum-int64(min))
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	//arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	//var arr = []int32{1, 2, 3, 4, 5}
	var arr = []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	//for i := 0; i < 5; i++ {
	//	arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	//	checkError(err)
	//	arrItem := int32(arrItemTemp)
	//	arr = append(arr, arrItem)
	//}

	miniMaxSum(arr)
}

//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
