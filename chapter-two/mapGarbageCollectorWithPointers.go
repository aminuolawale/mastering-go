package main

import (
	"runtime"
)


func main(){
	theMap := make(map[int]*int)
	for i:=0; i< 40000000; i++{
		theMap[i] = &i
	}

	runtime.GC()
	_ = theMap[0]
}