package main

import  "fmt"

/*
int uint8 uint16 uint64 uint32
int8 int16 int32 int64
 byte
 byte uint8
 rune
 rune int32
uint   32或者64

string 
char 字符
var a byte = 'a'
var 

字符串: " "  字面量 character literal
反引号:  ` `    

*/
func main(){
    a := 100
    var b byte = 100
    var c rune = 200
    var f rune = 'b'
    var g rune = '一'
    tmp :=`
    hello,what's a nice world 
    
    `
    fmt.Printf("%T, %v\n",a,a)
    fmt.Printf("%T, %v\n",b,b)
    fmt.Printf("%T, %v\n",c,c)
    fmt.Printf("%T, %v\n",f,f)
    fmt.Printf("%T, %v\n",g,g)
    fmt.Printf("%T, %v\n",tmp,tmp)


}
