package main
import "fmt"
var a int = 5
func main(){
    fmt.Println("the value of a is:",a)
    fmt.Println("the address of a is",&a)
    f1(&a)
    f2(a)
    fmt.Println("int main func the value of a ",a)
}

func f1(e *int){
    fmt.Println("pass address ")    
    b:=e
    *b+=1
    fmt.Printf(">>>%d\n",a)

}

func f2(k int){
    fmt.Println("pass value,then ")
    c:=&k
    *c+=2
    fmt.Printf("<<<value of proceed paased param %v,value of a %v\n",*c,a)

}

