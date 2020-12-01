package main

import (
	"strconv"
	"fmt"
)

func main(){
	result, err := strconv.ParseFloat("1112121", 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}