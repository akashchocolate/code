package main
import("fmt")
const count=10;//given 10 integer
func main(){
	a:=make([]int,count)//creating array
	for i:=0;i<count;i++{
	    fmt.Print(i)
		fmt.Print("enter an integer \n")
		fmt.Scanln(&a[i])//elements input
	    
	} 
	fmt.Print(a)
	bubblesort(a)//calling fuction
	fmt.Print("sorted array \t",a)
	}
func bubblesort(a []int){ //bubble sort function
	swap:=true
	for swap{
		for i:=0;i<(len(a)-1);i++{
			if a[i+1]<a[i]{
				Swap(a,i)//calling Swap function
				swap=true
			    
			}
			
			    swap=false
			    
			
		}
	}
}
func Swap (a []int, current int) { //swapping function
	a[current], a[current+1] = a[current+1],a[current]

}