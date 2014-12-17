package main
import "fmt"
func main(){
	fmt.Printf("%d is a even %v\n",7,even(7))
	fmt.Printf("%d is a odd %v\n",5,odd(5))
}

func even(a int )bool{
	if a==0 {return true}
	if odd(p(a-1)) {return true}
	return false
}
func odd(a int )bool{
	if a==0 {return false}
	if even(p(a-1)){ return true}
	return false
}
func p(a int)int{
	if a<0 {return -a}
	return a
}
