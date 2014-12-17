package main
import "fmt"

func main(){

	f1(0,1,3,2,7,4,5,6)
	a:=[]int{3,6,9,2,5,1}//use new or make is not right :[]int isn't a type
	f1(a...)//"a" must be a slice not a array
}

func f1(a...int){
	if len(a)<1{
		return
	}
	temp:=a[0]
	for _,b:=range a {
		if b>temp {
		temp = b
		}
	
	}
fmt.Println("the max num is:",temp)
}
