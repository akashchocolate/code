package main

import (
	"fmt"
)
var number int
var temp int
var sum int
var digit int

func main() {
	fmt.Print("enter the number \n")
	fmt.Scan(&number)
	sum=0
	digit=0
	temp=number
	for (temp>0){
		digit=temp%10
		sum=sum+(digit*digit*digit)
		temp=temp/10
	}
	if number==sum{
	    fmt.Print("you have entered an amstrong number")
	    
	}else{
	    fmt.Print("the number you have entered is not an amstrong number")
	}
	
}