package main
import ("fmt")
    var choice int
    var value1 int 
    var value2 int
    var ans int
func main(){
  
    fmt.Print("enter you choice \n")
    fmt.Print("1.addition \n")
    fmt.Print("2.substraction \n")
    fmt.Print("3.multiplication \n")
    fmt.Print("4.Division \n")
    fmt.Scan(&choice)
    fmt.Print("enter the first value \n")
    fmt.Scan(&value1)
    fmt.Print("enter the second value \n")
    fmt.Scan(&value2)
    switch choice{
        case 1:
          ans=add(value1,value2)
        case 2:
            ans=sub(value1,value2)
        case 3:
            ans=multi(value1,value2)
        case 4:
            ans=div(value1,value2)
        default:
            fmt.Print("invalid input \n")
    }
    fmt.Print(ans)
    
}
func add(value1 int, value2 int) int{
    
   return value1 + value2;
}

func multi( value1 int,  value2 int) int{
    
   return value1 * value2;
}

func div( value1 int,  value2 int) int{
    
    if value2 ==  0 {
      fmt.Print("you can not divide to 0")
    }

   return value1 / value2;
}

func sub( value1 int,  value2 int) int{
    
   return value1 - value2;
}