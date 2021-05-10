package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"sync"
	"time"
)

func median(values []float64) float64 {
	nums := make([]float64, len(values))
	copy(nums, values)
	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i]
	}

	return (nums[i-1] + nums[i]) / 2.0
}

func poolWorker(ch <-chan []float64, wg *sync.WaitGroup) {
	for values := range ch {
		m := median(values)
		log.Printf("median %v -> %f", values, m)
		wg.Done()
	}

	log.Printf("shutting down")
}

func multiDot(vectors [][]float64) {
	var wg sync.WaitGroup
	wg.Add(len(vectors))
	ch := make(chan []float64)

	for i := 0; i < runtime.NumCPU(); i++ {
		go poolWorker(ch, &wg)
	}

	for _, vec := range vectors {
		ch <- vec
	}

	wg.Wait()
	close(ch)
}

func main() {
	vectors := [][]float64{
		{1.1, 2.2, 3.3},
		{2.2, 3.3, 4.4},
		{3.3, 4.4, 5.5},
		{4.4, 5.5, 6.6},
		{5.5, 6.6, 7.7},
	}
	multiDot(vectors)
	time.Sleep(10 * time.Millisecond) // Let workers terminate
	fmt.Println("DONE")
}
