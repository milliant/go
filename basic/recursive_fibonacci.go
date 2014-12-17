package main
import (
		"log"
	   )
func main(){
	for i:=3;i<10;i++{
	a,b:=fabonacci(i,1)
	log.Printf("the index %v,the value %v",a,b)
	}

}

func fabonacci(a ,level int)(int,int){
	var temp int 
	if level ==1{
		tem:=a
	}
	if a<=1{
		return 1,a
	}else{
	return f(fabonacci,fabonacci,a-1,a-2),temp
	}
}
func f(f1,f2 func(int ,int),pa1,pa2 int)(int){
   _,v1 :=f1(pa1)
	_,v2:=f2(pa2)
	return v1+v2
}
