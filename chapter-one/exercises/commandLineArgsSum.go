package main

import (
	"os"
	"fmt"
	"strconv"
	"errors"
)

func main(){
	result, err := commandLineArgsSum()
	if err !=nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("this is the result", result)
}


func commandLineArgsSum()(float64, error){
	commandLineArgs := os.Args
	if len(commandLineArgs) <3{
		err := errors.New("please enter at least one value to be added")
		return 0.0, err
	}
	sum := 0.0
	for _, val := range commandLineArgs {
		interim, _ := strconv.ParseFloat(val, 64)
		sum += interim
	}

	return sum, nil
}