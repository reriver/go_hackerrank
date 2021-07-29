package main

import (
	"fmt"
	"strings"
)

func drawLine(spaces int32, hashes int32) {
	fmt.Print(strings.Repeat(" ", int(spaces)))
	fmt.Print(strings.Repeat("#", int(hashes)))
	fmt.Println("")
}

func staircase(n int32) {
	var spaces int32 = n - 1

	for ; spaces >= 0; spaces-- {
		drawLine(spaces, n-spaces)
	}
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	//nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	n := int32(4)

	staircase(n)
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
