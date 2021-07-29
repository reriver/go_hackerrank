package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	var s1 int32 = 0
	var s2 int32 = 0

	fmt.Println(len(arr), len(arr[0]))
	for j := 0; j < len(arr); j++ {
		s1 += arr[j][j]
		s2 += arr[len(arr)-1-j][j]
	}
	fmt.Println("s1 ", s1, " s2", s2)
	return int32(math.Abs(float64(s1 - s2)))
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()
	//
	//writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
	//
	//nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//n := int32(nTemp)

	var arrRow1 = []int32{11, 2, 4}
	var arrRow2 = []int32{4, 5, 6}
	var arrRow3 = []int32{10, 8, -12}
	var arr [][]int32
	arr = append(arr, arrRow1)
	arr = append(arr, arrRow2)
	arr = append(arr, arrRow3)

	fmt.Println("Main", arr)

	//for i := 0; i < int(n); i++ {
	//	arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")
	//
	//	var arrRow []int32
	//	for _, arrRowItem := range arrRowTemp {
	//		arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
	//		checkError(err)
	//		arrItem := int32(arrItemTemp)
	//		arrRow = append(arrRow, arrItem)
	//	}
	//
	//	if len(arrRow) != int(n) {
	//		panic("Bad input")
	//	}
	//
	//	arr = append(arr, arrRow)
	//}

	result := diagonalDifference(arr)
	fmt.Println("Result ", result)
	//fmt.Fprintf(writer, "%d\n", result)
	//
	//writer.Flush()
}

//
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
