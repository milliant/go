package main
import "fmt"
func main(){
	a()
	b()

	fmt.Println(">>>>>>>>>>>>>>>>>>>")
	c()
    
	fmt.Println(">>>>>>>>>>>>>>>>>>>")
	d()
}
func trace(a string)(string){
	fmt.Println("tracing ",a)
	return a
}
func untrace(a string){
	fmt.Println("untracing",a)
}
func a (){
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}
func b(){
	trace("b")
	defer untrace("b")
	fmt.Println("in b")

}
func c(){
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
}

func aa(){
	defer untrace(trace("a"))
	fmt.Println("in a ")
}

func d(){
	defer untrace(trace("b"))
	fmt.Println("in b")
	aa()
}



