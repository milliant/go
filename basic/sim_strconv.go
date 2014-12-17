package main
import (
    "fmt"
    "strconv"
)
//type news int
//new and int aren't the same type
func main(){
    var a string = "789"
    var an int
    var ann string

    an,_=strconv.Atoi(a)
    fmt.Printf("convert string to int:%v\n",an)
    an-=700
    ann=strconv.Itoa(an)
    fmt.Printf("convert int to string:%v\n",ann)



}
