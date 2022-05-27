package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type pair struct {
	row, col int
}

const length = 1000

var start time.Time
var rez [length][length]int

func main() {

	// utilize the max num of cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	var threadLength = flag.String("thread-len", "1", "Number of threads you want to use.")
	getThreadLength := func() int {
		i, err := strconv.Atoi(*threadLength)
		if err != nil {
			return 1
		}
		return i
	}
	flag.Parse()

	pairs := make(chan pair, 10)
	var wg sync.WaitGroup
	var a [length][length]int
	var b [length][length]int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			a[i][j] = rand.Intn(10)
			b[i][j] = rand.Intn(10)
		}
	}
	wg.Add(getThreadLength())
	for i := 0; i < getThreadLength(); i++ {
		go Calc(pairs, &a, &b, &rez, &wg)
	}
	start = time.Now()
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			pairs <- pair{row: i, col: j}
		}
	}
	close(pairs)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Binomial took ", elapsed)

}

func Calc(pairs chan pair, a, b, rez *[length][length]int, wg *sync.WaitGroup) {
	for {
		pair, ok := <-pairs
		if !ok {
			break
		}
		rez[pair.row][pair.col] = 0
		for i := 0; i < length; i++ {
			rez[pair.row][pair.col] += a[pair.row][i] * b[i][pair.col]
		}
	}
	wg.Done()
}
