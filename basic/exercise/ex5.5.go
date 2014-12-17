package main
import "fmt"
func main(){
 for i:=1;i<=5;i++{
    for j:=1;j<=i;j++{
        fmt.Print("G")
    }
    fmt.Println("")
 }
 var s string ="a"
 //s :=make([]string{c})
 for i:=0;i<5;i++{
    fmt.Println(s)
    s+="a"
 }
}
