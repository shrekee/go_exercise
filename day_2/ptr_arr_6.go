package main

import "fmt"

const COUNT int = 4
func main() {
//创建指针数组
    a := [COUNT]string{"abc","ABC","123","一二三"}
    fmt.Printf(" %T , %v \n" ,&a,&a)
    //定义指针数组
    var ptr [COUNT]*string
    fmt.Printf(" %T , %v \n" ,ptr,ptr)
    fmt.Printf(" %T , %v \n" ,&ptr,&ptr)

}
