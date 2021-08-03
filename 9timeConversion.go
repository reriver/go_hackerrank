package main

import (
	"fmt"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	layout1 := "3:04:05PM"
	layout2 := "15:04:05"

	fmt.Println("S =", s)
	t, _ := time.Parse(layout1, s)
	//fmt.Println(t.Format(layout2))
	//fmt.Println(t.Format(layout2))
	//fmt.Println(t)
	return t.Format(layout2)
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

	//s := readLine(reader)
	s := "07:05:45PM"
	result := timeConversion(s)

	fmt.Printf("Main %s\n", result)

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
