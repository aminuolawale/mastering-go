package main

import (
	"runtime"
)

type data struct {
	i, j int
}

func main(){

	var structure []data
	for i:=0; i< 40000000; i++{
		structure = append(structure, data{19, 10})
	}
	runtime.GC()
	_ = structure[40000000 -1]
}

