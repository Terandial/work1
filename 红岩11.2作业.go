package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup
var d string

func main() {
	over := make(chan bool, 1)
	for i := 1; i < 10; i++ {

		b(i) //goroutine具有无序性，其运行过程会使得前后的顺序紊乱，导致在运行之前就会使i达到9进而结束进程,虽然可以用锁，但去掉go不是最简单吗（决带笑）

		if i == 9 {
			over <- true
		}
	}

	fmt.Println("over!!!")
}
func b(i int) {

	fmt.Println(i)

}
