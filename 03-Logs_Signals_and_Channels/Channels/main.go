package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Without channels
func generateRUndonNumber1() int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(100)
}

func test1() {

	startTime := time.Now()
	for num := 0; num < 100; num++ {
		fmt.Printf("%d\n", generateRUndonNumber1())
	}
	fmt.Println(time.Since(startTime))

}

// With channels

func generateRUndonNumber2(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	ch <- r.Intn(100)

}

func test2() {
	var ch chan int = make(chan int)
	var wg sync.WaitGroup
	var startTime time.Time = time.Now()

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go generateRUndonNumber2(ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
	fmt.Println(time.Since(startTime))

}

func main() {
	test1()
	test2()
}
