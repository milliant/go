package main
import "fmt"
var a string
func main(){
    a = "G"
    fmt.Printf(a)
    f1()
}
func f1(){
    a:="O"
    fmt.Printf(a)
    f2()
}
func f2(){
    fmt.Printf(a)
}

