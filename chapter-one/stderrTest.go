package main


import (
	"os"
	"io"
)

func main(){
	io.WriteString(os.Stderr, "this is not me")
}