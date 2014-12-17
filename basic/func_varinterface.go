package main
import "fmt"
func main(){
	var a,b int = 4,5
	var c string = "hello world"
	var d float64 = 3.14159
	f1(a,b,c,d)
}
func f1(a...interface{}){
	for _,temp:=range a{
	   fmt.Printf("the value is:%v\n",temp)
	}
}


