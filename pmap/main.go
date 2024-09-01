package main

import "fmt"

func Map(f func(int) int, input []int) []int {
	res := []int{}

	for _, x := range input {
		res = append(res, f(x))
	}

	return res
}

func main() {
	xs := []int{1, 2, 3, 4, 5}

	res := Map(func(x int) int { return x * x * x }, xs)

	for _, x := range res {
		fmt.Print(x, " ")
	}
	fmt.Println()
}

// дз. написать многопоточную версию Map
// то есть каждое вычисление применяемой функции д.б. в отдельной горутине
