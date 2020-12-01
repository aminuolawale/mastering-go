package main

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)

func main(){
	result, err := commandLineArgsAverage()
	if err !=nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("this is the result", result)
}


func commandLineArgsAverage()(float64, error){
	commandLineArgs := os.Args
	if len(commandLineArgs) <2{
		err := errors.New("please enter at least one value to be added")
		return 0.0, err
	}
	sum := 0.0
	for _, val := range commandLineArgs {
		interim, _ := strconv.ParseFloat(val, 64)
		sum += interim
	}
	 result := sum / float64(len(commandLineArgs)-1)

	return result, nil
}