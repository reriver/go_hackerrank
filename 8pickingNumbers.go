package main

import (
	"fmt"
	"sort"
)

func pickingNumbers(a []int32) int32 {
	fmt.Println(a)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	fmt.Println(a)

	var maxCount int32 = -1
	var count int32 = 1
	for i := 0; i < len(a)-1; i++ {
		if 0 < a[i+1]-a[i] && a[i+1]-a[i] < 2 {
			count++
			if count > maxCount {
				maxCount = count
			}
		}
		if a[i+1]-a[i] > 2 {
			count = 0
		}
		if a[i+1]-a[i] <= 0 {
			fmt.Errorf("erro array not sorted")
			return -1
		}
		fmt.Println("MaxCount = ", maxCount)
		fmt.Println("count = ", count)
	}
	return maxCount
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

	//aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a = []int32{4, 6, 5, 3, 3, 1}

	//for i := 0; i < int(n); i++ {
	//	aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
	//	checkError(err)
	//	aItem := int32(aItemTemp)
	//	a = append(a, aItem)
	//}

	result := pickingNumbers(a)

	fmt.Printf("%d\n", result)

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
