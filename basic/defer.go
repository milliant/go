package main
import "fmt"

func main(){
	for i:=0;i<5;i++{
	defer fmt.Println("the value is :",i)
	}
	fmt.Println(">>>>>>>>>>>")
	
	var j int = 9
	
	defer fmt.Println("defer value is",j)

	j++
	fmt.Println("the value of j is :",j)



}
