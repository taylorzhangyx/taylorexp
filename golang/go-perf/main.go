package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

//
// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"reflect"
// 	"runtime"
// 	"runtime/pprof"
// 	"sync"
// 	"time"
//
// 	"taylorzhangyx.com/golang-perf/logic"
// )
//
// type Func func(interface{})
//
// func main() {
// 	var p int64 = 0
// 	flag.Int64Var(&p, "loop", 10000000, "use general loop to test performance")
//
// 	if err := profileMemory(); err != nil {
// 		fmt.Println("failed to profile memory")
// 		os.Exit(-1)
// 	}
//
// 	funcs := []Func{
// 		logic.ListLoop,
// 		logic.ListLoopWithSize,
// 		logic.MapLoop,
// 		logic.MapLoopWithSize,
// 	}
// 	startTs := time.Now()
// 	wg := sync.WaitGroup{}
// 	for _, f := range funcs {
// 		wg.Add(1)
// 		go func(f Func) {
// 			defer wg.Done()
// 			MeasureTime(f, p)
// 		}(f)
// 	}
// 	wg.Wait()
// 	println("finished all time cost", time.Since(startTs).String())
// 	// globalMap = map[string]string{}
// 	if err := profileMemory(); err != nil {
// 		fmt.Println("failed to profile memory")
// 		os.Exit(-1)
// 	}
// }
//
// func profileMemory() (err error) {
// 	file, err := os.Create("mem.prof")
// 	if err != nil {
// 		return
// 	}
// 	runtime.GC()
// 	err = pprof.WriteHeapProfile(file)
// 	return
// }
//
// func MeasureTime(f Func, p interface{}) {
// 	startTs := time.Now()
// 	defer func() {
// 		endTs := time.Since(startTs)
// 		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
// 		println(fmt.Sprintf("%s time cost: %v", funcName, endTs.String()))
// 	}()
// 	f(p)
// }

type GlobalStruct struct {
	bar map[string]string
}

var globalStruct = GlobalStruct{
	bar: map[string]string{},
}

var globalMap = map[string]string{}

func profileMemory() (err error) {
	file, err := os.Create("mem.prof")
	if err != nil {
		return
	}
	runtime.GC()
	err = pprof.WriteHeapProfile(file)
	return
}

func bar() {
	// var prev *Foo
	for i := 0; i < 1000000; i++ {
		key := fmt.Sprintf("map key: %d", i)
		globalStruct.bar[key] = key
		// globalMap[key] = key
	}
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	bar()
	globalStruct = GlobalStruct{}
	// globalMap = map[string]string{}
	if err := profileMemory(); err != nil {
		fmt.Println("failed to profile memory")
		os.Exit(-1)
	}
}
