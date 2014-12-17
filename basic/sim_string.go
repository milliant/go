package main
import (
		"fmt"
		"strings"
		)
func main(){
	a:="The world is full of love and evil,love and evil,love and evil..."
	b:=strings.HasPrefix(a,"The")
	fmt.Printf("the string a:%s prefix is \"The\":%v\n",a,b)
	
	c:=strings.Index(a,"love")
	d:=strings.LastIndex(a,"love")
	fmt.Printf("the index of \"love\" in string a is %d,and last index is %d\n",c,d)

	fmt.Printf("the string contains \"full\"?%v\n",strings.Contains(a,"full"))

	fmt.Printf("after replace\n%s\n",strings.Replace(a,"evil","love",2))

	e:=strings.Repeat(a,2)
	fmt.Printf("the len of a is %v,and the length of e is %v\n",len(a),len(e))

	fmt.Printf("to UPPER:%v\n",strings.ToUpper(a))
	fmt.Printf("to LOWER:%v\n",strings.ToLower(a))

	fmt.Printf("after remove head and tail blanks%s\n",strings.TrimSpace(a))
	f:=a
	fmt.Printf("after remove head&&tail \"The\"%v\n",strings.Trim(f,"The"))
}
