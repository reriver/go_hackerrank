package main

import "fmt"

type result int

const (
	resultWin  result = 0
	resultDraw result = iota
	resultLose
	resultError result = -1
)

type prize int

const (
	prizeBig   prize = 0
	prizeSmall prize = 1
	prizeNo    prize = 2
)

type goal int

type score struct {
	a goal
	b goal
}

func scoreToResult(sc score) result {
	if sc.a < sc.b {
		return resultLose
	} else if sc.a == sc.b {
		return resultDraw
	} else if sc.a > sc.b {
		return resultWin
	} else {
		return resultError
	}
}

func returnPrize(exp score, real score) prize {

	if exp.a == real.a && exp.b == real.b {
		return prizeBig
	}

	expRes := scoreToResult(exp)
	realRes := scoreToResult(real)

	if expRes == realRes {
		return prizeSmall
	}

	return prizeNo
}

func main() {
	fmt.Println(returnPrize(score{1, 2}, score{1, 1}))
}
