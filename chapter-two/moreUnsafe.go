package main


import (
	"fmt"
	"unsafe"
)


func main(){
	array := []int{1, 3, 6, 6, 775}
	pointer := &array[0]
	mem := uintptr(unsafe.Pointer(pointer)) +unsafe.Sizeof(array[0])
	for i:=0; i< len(array)-1; i++ {
		pointer = (*int) (unsafe.Pointer(mem))
		fmt.Println("the current value", *pointer)
		mem = uintptr(unsafe.Pointer(pointer)) +unsafe.Sizeof(array[0])
	}
}