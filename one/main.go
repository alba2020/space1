package main

import "fmt"

func writer(ch chan<- int) {
	ch <- 100
	ch <- 2
	ch <- 3

	close(ch)
}


func main() {
	ch := make(chan int)
	
	go writer(ch)
	
	reader(ch)
}

func reader(ch <-chan int) {
	max := 0
	for val := range ch {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}


// 1. канал связывает горутины (а не просто функции)
// 2. в момент передачи данных через канал горутины должны одновременно
// - читать
// - писать
// т.е. чтение-запись в канал это точка синхронизации 2 горутин
// 3. кто пишет в канал, тот его и закрывает
// 4. практически все можно реализовать без wg и буферизации каналов
