package main

import (
	"errors"
	"fmt"
)

func main()  {

	re, er := Divide(100,0)
	if er != nil{
		fmt.Println(er)
	}else {
		fmt.Println(re)
	}
}

func Divide(x,y float64) (float64, error){
	var result float64
	if y == 0 {
		return result, errors.New("cannot divide by o")
	}
	result = x / y
	return result, nil
}


