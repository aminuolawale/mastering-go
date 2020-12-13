package main

import (
	"runtime"
)


func main(){
	split := make([]map[int]int, 200)
	for i := range split {
		split[i] = make(map[int]int)
	}

	for i:=0; i< 40000000; i++{
		split[i%200][i] =  i
	}
	runtime.GC()
	_ = split[0][0]
}