package main

import (
	"fmt"
	"runtime"
	_"time"
)


func PrintMemStats(mem runtime.MemStats){
	// runtime.MemStats can store runtime Memstats
	runtime.ReadMemStats(&mem)
	// we copy current mem stats into mem
	// now mem has some usefull attributes we can access
	// runtime.Memstats must therefore be a struct
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}



func main(){
	var mem runtime.MemStats
	PrintMemStats(mem)
	
	for i:=0; i<10; i++{
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("operation failed")
		}
	}
	PrintMemStats(mem)

	// doing more memory allocation
	for i:=0; i<10; i++{
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("operation failed")
		}
		// time.Sleep(5 * time.Second)
	}
	PrintMemStats(mem)
}