package main

import "fmt"

func writer(ch chan int) {
	for _, x := range []int{2, 3, 4, 5, 6} {
		ch <- x
	}
	close(ch)
}

func runComputation(ch chan int) chan int {
	result := make(chan int)

	go func() {
		result <- numberOfEvens(ch)
		close(result)
	}()

	return result
}

func main() {
	input := make(chan int)

	go writer(input) // non-blocking
	
	// res := runComputation(input) // non-blocking!
	// n := <-res // blocking
	
	n := <- runComputation(input) // blocking!
	
	fmt.Println("evens", n)
}

// 1 2 3 -> 1
// 2 3 4 -> 2
// 99 -> 0

// должно вернуть количество четных чисел в канале ch
func numberOfEvens(ch chan int) int {
	count := 0
	for v := range ch {
		if v%2 == 0 {
			count += 1
		}
	}

	return count
}
