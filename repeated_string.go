package main

import (
	"fmt"
	"strings"
)

func repeatedString(s string, n int64) int64 {
	asAtOneString := int64(strings.Count(s, "a"))
	fmt.Print("s = ", s, ", asAtOneString = ", asAtOneString)

	l := int64(len(s))
	fmt.Println(", l = ", l)

	aSum := int64(0)

	asAtN := int64(n / l)
	fmt.Println("asAtN = ", asAtN)

	aSum = asAtN * asAtOneString
	fmt.Println("aSum = ", aSum)

	asRemainder := int64(n % l)
	sByteSlice := []byte(s)

	lastString := sByteSlice[:asRemainder]

	aSum += int64(strings.Count(string(lastString), "a"))
	return aSum
}

func main() {
	//s := "aba"
	//n := int64(10)
	s := "abcac"
	n := int64(10)
	//s := "a"
	//n := int64(1000000000000)

	result := repeatedString(s, n)

	fmt.Printf("%d\n", result)
}
