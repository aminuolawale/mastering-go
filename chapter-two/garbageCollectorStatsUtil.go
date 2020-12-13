package main

import (
	"fmt"
	"runtime"
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



