package main
import  "fmt"
var num int = 5
var num1,num2 int
func main(){
    num1,num2 = f1(num)
    printinfo()
    f2(num)
    printinfo()

}

func f1(input int)(int,int){
    return 1*input,2*input

}


func f2(input int){
    num1 = input*2
    num2 = input*4

}

func printinfo(){
    fmt.Println("the value of num1 and num2 is ",num1,num2)

}
