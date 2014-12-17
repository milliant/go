package main
import (
    "fmt"
    "time"
    "math/rand"
)
func main(){
    for i:=0;i<5;i++{
    a := rand.Int()
    fmt.Printf("%d/\n",a)

    }
    fmt.Println(">>>>>>>>>>>>>>>")
    for i:=0;i<5;i++{
    a := rand.Intn(8)
    fmt.Printf("%d/\n",a)
    }
    fmt.Println(">>>>>>>>>>>>>>>")
    timens := int64(time.Now().Nanosecond())
    rand.Seed(timens)
    for i:=0;i<5;i++{
    fmt.Printf("%2.2 /\n",100*rand.Float32()) 
    }
    fmt.Println(time.Now())
}
