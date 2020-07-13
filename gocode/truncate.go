package main
import "fmt"
func main(){
	var number int
	fmt.Print("please enter a floating point number\n")
	fmt.Scanln(&number)
	fmt.Print("converted integer",number)
}
	