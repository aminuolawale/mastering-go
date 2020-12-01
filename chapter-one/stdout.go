package main

import (
	"io"
	"os"
	"strings"
)


func main(){
	vals := os.Args
	var myString string
	if len(vals) == 1{
		myString = "Please pass the argument to be printed out"
	} else {
		interim := os.Args[1:]
		myString = strings.Join(interim, " ")
	}
	myString += "\n"
	io.WriteString(os.Stdout, myString)
}