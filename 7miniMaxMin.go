package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func miniMaxSum(arr []int32) {
	var sum int32 = 0
	var sums []int32

	fmt.Println(arr)
	for i := range arr {
		sum += int32(i)
	}
	fmt.Println(sum)
	for i := range arr {
		sums = append(sums, sum-int32(i))
	}
	fmt.Println(sums)
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	//arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr = []int32{1, 3, 5, 7, 9}

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
