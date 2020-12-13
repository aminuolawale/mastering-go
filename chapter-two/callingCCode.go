package main

// #include <stdio.h>
// void saySomething(){
// 	printf("this is me just printing a number");
// }
import	"C"


import "fmt"


func main(){
	fmt.Println("I am going to call a C function now")
	C.saySomething()
	fmt.Println("I am done")
}