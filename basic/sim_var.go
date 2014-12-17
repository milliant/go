package main
import (
    "fmt"
    "os"
    )
func main(){
    var(
    a int
    b bool
    c float64
    HOME = os.Getenv("HOME")
    USER = os.Getenv("USER")
    GOROOT = os.Getenv("GOROOT")
    )
    a = 5
    b = true
    c = 3.14
    fmt.Printf("the value of a :%d,\n",a)
    a = int(c)
    fmt.Printf("convert float to int ,value is :%d,the bool value is %b\n",a,b);
    fmt.Printf("the infomation get from the operation system is: HOME: %s\n USER:%s\n GOROOT:%s\n",HOME,USER,GOROOT)
}
