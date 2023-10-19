package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func main() {
// 	threadPtr := flag.Int("thread", 1, "number of threads to use")
// 	maxprocsPtr := flag.Int("maxprocs", 1, "number of cpu cores to use")

// 	flag.Parse()

// 	fmt.Printf("Using %d threads\n", *threadPtr)
// 	fmt.Printf("Using %d MAXPROC\n", *maxprocsPtr)

// 	runtime.GOMAXPROCS(*maxprocsPtr)

// 	// Your code goes here
// 	fmt.Println("starting routine")

// 	wg := sync.WaitGroup{}
// 	for i:=0; i<*threadPtr;i++ {
// 		wg.Add(1)
// 		go func(n int, wgp *sync.WaitGroup){
// 			for j:=0; j<5000000; j++{
// 					time.Sleep(time.Microsecond)
// 			}
// 			fmt.Printf("routine %d DONE\n", n)
// 			wgp.Done()
// 		}(i, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("exiting...")
// }

func runForever(id int) {
	fmt.Printf("id: %d\n", id)
	for {
		time.Sleep(1)
	}
}
func main() {
	// threadPtr := flag.Int("thread", 1, "number of threads to use")
	maxprocsPtr := flag.Int("maxprocs", 1, "number of cpu cores to use")
	fmt.Printf("Using %d MAXPROC\n", *maxprocsPtr)

	runtime.GOMAXPROCS(*maxprocsPtr)
// 	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go runForever(i)
	}
	wg.Wait()
}
