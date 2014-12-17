package main
import (
    "fmt"
    "math/rand"
    "runtime"
    "time"
)

func main(){


    if runtime.GOOS== "windows"{
        fmt.Println("we are test in the windows platform")
    }else{
        fmt.Println("we are running on the unix like system")
    }
    // above didn't get the random num ,every time is the same
    i:=rand.Intn(99)
    fmt.Printf("random number is:%d\n",i)
    if i>90 {
        fmt.Println(">90")
    }else if(i>80){
        fmt.Println("80<x<90")
    }else if(i>50){
        fmt.Println("50<x<80")
    }else if(i>25){
        fmt.Println("25<x<50")
    }else{
        fmt.Println("x<=25")
    }

    rand.Seed(int64(time.Now().Nanosecond()))
    if i:=rand.Float32();i<5{
        fmt.Println("true")
    }else{
        fmt.Println("false")
    }
        


}
