package main

import "fmt"

func sign(i int32) int32 {
	if i < 0 {
		return 0
	} else if i > 0 {
		return 2
	}
	return 1
}

func plusMinus(arr []int32) {
	var i int32 = 0
	res := [3]int{0, 0, 0}
	for ; i < int32(len(arr)); i++ {
		//fmt.Println("i =", i, ", arr[i] =", arr[i], ", sign[i] =", sign(arr[i]), "arr[sign(arr[i])] = ", res[sign(arr[i])], ", sign(arr[i]) = ", sign(arr[i]))
		res[sign(arr[i])] += 1
		//fmt.Println("res[sign(arr[i])] =", res[sign(arr[i])])
	}

	//fmt.Println("res[0] =", res[0], "res[1] =", res[1]," res[2] =", res[2])
	fmt.Printf("%.6f\n", float32(res[0])/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(res[1])/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(res[2])/float32(len(arr)))
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	//
	//nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//n := int32(nTemp)
	//
	//arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr = []int32{-4, 3, -9, 0, 4, 1}

	//for i := 0; i < int(n); i++ {
	//	arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	//	checkError(err)
	//	arrItem := int32(arrItemTemp)
	//	arr = append(arr, arrItem)
	//}

	plusMinus(arr)

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
