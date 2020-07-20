package main
import "fmt"
func main(){
	var number int
	fmt.Print("enter the number \n")
	fmt.Scanln(&number)
	for i:=1;i<= (number/2);i++{
		if ((number%i)==0){
			fmt.Print(i,"\t")
		}
	}
}