package main

import (
	"fmt"
	"sort"
)

func swapInt32(a *int32, b *int32) {
	*a, *b = *b, *a

}

func permutationsInt32(arr []int32) [][]int32 {
	var helper func([]int32, int32)
	var res [][]int32

	fmt.Println("Start", arr)

	helper = func(arr []int32, n int32) {
		if n == 1 {
			tmp := make([]int32, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; int32(i) < n; i++ {

				helper(arr, n-1)

				if n%2 == 1 {
					//swapInt32(&arr[i], &arr[n - 1])
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
					//swapInt32(&arr[0], &arr[n - 1])
				}
			}
		}
	}
	helper(arr, int32(len(arr)))
	return res
}

func solve(x []int32) int64 {
	var in []int32

	var xInt []int

	fmt.Println(x)
	for _, xI := range x {
		xInt = append(xInt, int(xI))
	}
	fmt.Println(xInt)
	sort.Ints(xInt)
	fmt.Println(xInt)

	for i := range x {
		x[i] = int32(xInt[i])
	}
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		in = append(in, x[i])
	}
	fmt.Println(x)

	//var res [][]int32
	//res = permutations(in)
	fmt.Println(permutationsInt32(in))
	//for rowIndex, row := range res {
	//
	//}
	return 10
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

	//aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a = []int32{4, 2, 1, 3}
	//n := int32(len(a))

	//for i := 0; i < int(n); i++ {
	//aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
	//checkError(err)
	//aItem := int32(aItemTemp)
	//a = append(a, aItem)
	//}

	result := solve(a)

	fmt.Printf("%d\n", result)

	//writer.Flush()
}

//
//func readLine(reader *bufio.Reader) string {
//str, _, err := reader.ReadLine()
//if err == io.EOF {
//return ""
//}
//
//return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//if err != nil {
//panic(err)
//}
//}
