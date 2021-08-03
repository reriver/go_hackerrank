package main

import (
	"fmt"
)

func absInt32(a int32) int32 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func minimumBribes(q []int32) {
	var minBribes int32 = 0
	var i int32 = 0
	for ; i < int32(len(q)); i++ {
		fmt.Print("q[i] = ", q[i], " i + 1 = ", i+1, "abs = ", absInt32(q[i]-(i+1)))
		if absInt32(q[i]-(i+1)) > 2 {
			fmt.Println("Too chaotic")
			return
		}
		if absInt32(q[i]-(i+1)) > 0 {
			minBribes += absInt32(q[i] - (i + 1))
		}
		fmt.Println("abs = ", absInt32(q[i]-(i+1)), " miniBrides = ", minBribes)
	}
	fmt.Println(minBribes / 2)
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	//tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//t := int32(tTemp)
	//
	//for tItr := 0; tItr < int(t); tItr++ {
	//	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//	checkError(err)
	//	n := int32(nTemp)
	//
	//	qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	//var q = []int32 {1, 2, 3, 5, 4, 6, 7, 8, 9}
	//var q = []int32 {2, 5, 1, 3, 4}
	var q = []int32{1, 2, 5, 3, 7, 8, 6, 4}

	fmt.Println(q)
	//for i := 0; i < int(n); i++ {
	//	qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
	//	checkError(err)
	//	qItem := int32(qItemTemp)
	//	q = append(q, qItem)
	//}

	minimumBribes(q)
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
