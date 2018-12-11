package main

import (
	"fmt"
	"errors"
)

type operate func(x,y int) int

func calculate (x int, y int, op operate) (int,error) {
	if op == nil {
		return 0 ,errors.New("invalid error")
	}
	return op(x,y),nil
}

type calculateFunc func(x int, y int) (int, error)

func genCalculator (op operate) calculateFunc {
	return func(x int, y int) (int,error) {
		if op == nil {
			return 0 ,errors.New("invalid error")
		}
		return op(x,y),nil
	}
}

func main() {
	//1
	x,y := 12, 21
	op := func(x,y int) int {
		return x+y
	}
	result,err := calculate(x,y,op) 
	fmt.Printf("The result: %d (error: %v)\n",result, err)

	//2
	x, y = 13, 31
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n",result, err)
}