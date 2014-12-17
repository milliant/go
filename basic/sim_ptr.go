package main
import "fmt"
func main(){
    var i1 = 9
    fmt.Print("the value %d is at the memory location %p\n",i1,&i1)
    var ptr *int
    ptr = &i1
    fmt.Print("at the location of %p is the value of %d\n",ptr,*ptr)
}
