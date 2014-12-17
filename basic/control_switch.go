package main
import (
    "fmt"
)

func main(){
    var a int  = 89
    switch {
    case a==98||a==99:fmt.Println("98,99")
    case a<90:fmt.Println("<90")
    case a==12||a==13:fmt.Println("12,13")
    case a==10 :fallthrough
    case a==9:f()
    
    
    }


}

func f(){

    fmt.Println("fallthrough")

}
