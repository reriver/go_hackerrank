package main

import (
	"fmt"
)

func rotLeft(a []int32, d int32) []int32 {
	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(len(a))
	res := []int32(a[d:len(a)]) //int32(len(a)) - int32(1) -
	fmt.Println(res)
	res = append(res, a[0:d]...)
	fmt.Println(res)
	return res
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
	//firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	//
	//nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	//checkError(err)
	//n := int32(nTemp)
	//
	//dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	//checkError(err)
	//d := int32(dTemp)
	//
	//aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a = []int32{1, 2, 3, 4, 5}
	//
	//for i := 0; i < int(n); i++ {
	//	aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
	//	checkError(err)
	//	aItem := int32(aItemTemp)
	//	a = append(a, aItem)
	//}

	result := rotLeft(a, 4)

	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if i != len(result)-1 {
			fmt.Printf(" ")
		}
	}

	fmt.Println("\n")
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
