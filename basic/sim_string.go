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
    
    fmt.Println("splite by blankspace")
    for i,g:=range strings.Fields(a){
         fmt.Printf("%d:%v\n",i,g)
    }
    fmt.Println("splite by \"love\":")
    for i,g:= range strings.Split(a,"love"){
        fmt.Printf("%v:%v\n",i,g)
    }
    h:=strings.Split(a,"and")
    fmt.Printf("slice is: %v\n",h)
    fmt.Printf("after join:%v\n",strings.Join(h,"or"))
    fmt.Println("READ BY ASCII")
    var s string = "这个世界充满爱和罪恶，爱和罪恶，爱和罪恶..."
    strslice :=make([]byte,100)
    _,err:=strings.NewReader(s).Read(strslice)
    if nil!=err {
        fmt.Println("read err")
    }else{
        fmt.Printf("the read message is:%v\n",string(strslice))
    }
    //ReadByte()只读一个字节的数据
    data,error:=strings.NewReader(s).ReadByte()
    if nil!=error{
        fmt.Printf("readbyte err\n")
    }else{
        fmt.Printf("readbyte messge is:%v\n",string(data))
    }
   // for i,g:=range strings.NewReader(s).Read(strslice,10){
   //     fmt.Printf("%v:%v\n",i,g)
   // }
   // fmt.Println("READ BY UNICODE\n")
   // for i,g:=range strings.NewReader(s).ReadByte(s){
   //    fmt.Printf("%v:%v\n",i,g)
   // }

}
