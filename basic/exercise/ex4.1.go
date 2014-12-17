package main
import "fmt"
var a="G"
func main(){
    n()
    m()
    n()

}

func n(){
    fmt.Printf("the value of a is :%v\n",a)
}

func m(){
    a := "O"
    fmt.Printf("the value of a is:%v\n",a)
}
