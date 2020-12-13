package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var j int64
	j = 5434123412312431212
	p1 := &j
	p2 := (*int32) (unsafe.Pointer(p1))

	fmt.Println(p1, *p1)
	fmt.Println(p2, *p2)
}