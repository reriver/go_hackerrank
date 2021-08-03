package main

import (
	"fmt"
)

func birthdayCakeCandles(candles []int32) int32 {
	var max int32 = -2147483648
	var sum int64 = 0
	var number int32 = 0

	fmt.Println(candles)

	for i := range candles {
		sum += int64(candles[i])
		if max < candles[i] {
			max = candles[i]
			number = 0
		}
		if max == candles[i] {
			number++
		}
	}
	return number
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
	//candlesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//
	//candlesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	//var candles = []int32 {2, 2, 1, 3, 2, 3, 1}
	//var candles = []int32 {3, 2, 1, 3}
	var candles = []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
	//var candles = []int32 {1, 2, 3, 4, 5, 6, 7, 7, 7, 8}

	//for i := 0; i < int(candlesCount); i++ {
	//	candlesItemTemp, err := strconv.ParseInt(candlesTemp[i], 10, 64)
	//	checkError(err)
	//	candlesItem := int32(candlesItemTemp)
	//	candles = append(candles, candlesItem)
	//}

	result := birthdayCakeCandles(candles)
	//
	fmt.Printf("%d\n", result)
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
