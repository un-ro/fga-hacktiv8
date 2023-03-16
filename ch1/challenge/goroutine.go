package main

import (
	"fmt"
	"sync"
)

// Session 5 || D5W2
func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex

	data := [][]string{
		{"bisa1", "bisa2", "bisa3"},
		{"coba1", "coba2", "coba3"},
	}

	mtx.Lock()
	mtx.Unlock()

	// WaitGroup
	for i := 0; i < 2; i++ {
		for j := 1; j <= 4; j++ {
			wg.Add(1)
			go useWaitGroup(&wg, data[i], j)
		}
	}

	// Mutex
	for i := 0; i < 2; i++ {
		for j := 1; j <= 4; j++ {
			wg.Add(1)
			go useMutex(&mtx, &wg, data[i], j)
		}
	}

	wg.Wait()
}

func useWaitGroup(wg *sync.WaitGroup, data []string, i int) {
	defer wg.Done()
	fmt.Println(data, i)
}

func useMutex(mtx *sync.Mutex, wg *sync.WaitGroup, data []string, i int) {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(data, i)
	wg.Done()
}
