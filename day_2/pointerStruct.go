package main

import "fmt"

type Student struct {
    name string
    age uint16
    married bool
    gender uint8

}
func main() {
    s1 := Student{"Steven",30,true,'f'}
    s2 := Student{"sunny",20,false,'m'}
    
    var a *Student = &s1
    var b *Student = &s2
    fmt.Println("\n-------------------")
    fmt.Printf("s1的类型： %T 值是：%v\n",s1,s1) 
    fmt.Printf("a的类型： %T 值是：%v\n",a,a) 
    println("b的值是：",b)
    /*
    fmt.Printf("s1",s1,a,b) 
    fmt.Printf("s1",s1,a,b) 
    */
    println(&a.name,&a.age,&a.married,&a.gender)
}
