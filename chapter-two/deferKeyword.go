package main


import (
	"fmt"
)


func main(){
	fmt.Println("runing d1")
	d1()
	fmt.Println("running d2")
	d2()
	fmt.Println("running d3")
	d3()
}


func d1(){
	defer fmt.Println()
	for i:=1; i<= 3; i++{
		defer fmt.Print(i, "  ")
	} 
}


func d2(){
	defer fmt.Println()
	for i:=1; i<= 3; i++{
		defer func(){
			fmt.Print(i, "  ")
		}()
	} 
}

func d3(){
	defer fmt.Println()
	for i:=1; i<= 3; i++{
		defer func(n int){
			fmt.Print(n, "  ")
		}(i)
	} 
}