package main
import (
		"fmt"
		"os"
		"math/rand"
		"time"
		)
func main(){
	var a,b,c int = 1,2,'b'
	a,b=b,a
	a,c=c,a
	fmt.Printf("the value of a,b and c is:%c,%v,%v",a,b,c)
	d:=os.Getenv("HOME")
	fmt.Printf("the HOME is %v\n",d)
	for i:=0;i<5;i++{
	fmt.Printf("the random value is:%v",rand.Int())
	}
	fmt.Printf(">>>>>>>>>>>>>>>>>>>\n");
	for i:=0;i<5;i++{
	fmt.Printf("the random value is:(<9)%v\n",rand.Intn(9))
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>")
	ti:=int64(time.Now().Nanosecond())
	rand.Seed(ti)
	for i:=0;i<5;i++{
	fmt.Printf("%d / ",rand.Int())
	}


}
