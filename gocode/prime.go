package main
import "fmt"
func main(){
    var number int
    var count int
    var i int
    fmt.Print("enter the number \n")
    fmt.Scan(&number)
    count=0
    if number>1{
    for i=2;i<=number;i++{
        if (number%i)==0{
            count++
        }
    }
    if count>2{
        fmt.Print("not a prime number")
    }else{
        fmt.Print(" prime number")
    }}else{
        fmt.Print("not a prime number")
    }}