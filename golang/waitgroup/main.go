package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cc := make(chan int)

	go func() {
		once := sync.Once{}
		count := 1
		for {
			time.Sleep(time.Millisecond * 200)
			once.Do(func() {
				cc <- count
			}		)
			count += 1
			fmt.Println("count:", count)
		}
	}()

	select {
	case n := <-cc:
		fmt.Println(fmt.Sprintf("got %d from channel", n))
		close(cc)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
		break
	}
	time.Sleep(time.Second * 2)
	fmt.Println("done")
}
