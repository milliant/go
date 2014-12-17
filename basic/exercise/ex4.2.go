package main
import "fmt"

var a = "G"
func main(){
    m()
    n()
    m()

}

func m(){
    fmt.Printf(a)
}

func n(){
    a = "o"
    fmt.Printf(a)

}
