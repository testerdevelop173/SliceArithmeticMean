package main

import (
	"fmt"
	"sync"
	//"bufio"
)

const goroutines int = 2

var data = []int{34, 50, 97, 46, 53, 457, 89, 101}
var mu sync.Mutex // Mutex for synchron access to sum

func sumSlice(in []int, start, end int) int {
	localSum := 0
	for j := start; j < end; j++ {
		localSum += in[j]
	}
	return localSum
}

func main() {
	var sum int //переменная, куда функция sumSlice помещает результат суммир-я слайса
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		start := i * len(data) / goroutines
		end := (i + 1) * len(data) / goroutines
		go func(start, end int) {
			defer wg.Done()
			partialSum := sumSlice(data, start, end)
			mu.Lock()
			sum += partialSum
			mu.Unlock()
		}(start, end)
	}

	wg.Wait()

	mean := float64(sum) / float64(len(data))
	fmt.Println("Arithmetic Mean:", mean)
}
