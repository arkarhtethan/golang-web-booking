package main

import (
	"fmt"
	"time"
)

type Summ struct {
	f  int
	t  int
	eq int
}

func worker(c chan Summ, a, b int) {
	res := 0
	for n := a; n < b; n++ {
		res += n
		time.Sleep(time.Millisecond)
	}
	c <- Summ{f: a, t: b, eq: res}
}

func main() {
	fmt.Println("Start")
	c := make(chan Summ, 5)
	for x := 0; x < 5; x++ {
		go worker(c, x, x+10)
	}
	for x := 0; x < 5; x++ {
		fmt.Printf("FTE = %v\n", <-c)
	}
	fmt.Println("End")
}
