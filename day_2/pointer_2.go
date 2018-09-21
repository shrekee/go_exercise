package main

import "fmt"


/*
指针变量；
    变量是一种使用方便的占位符，变量都指向计算机的内存地址。
    一个指针可以指向任何一个值的内存地址。
go 语言指针的最大不同：
    Go指针不能运算


    var ip *int   //指向整型的指针
    var fp *float32   //指向浮点型的指针
*/

func main() {
    a := 20
    var ip *int
    ip = &a
    fmt.Printf("a 的类型是%T, 值是%v \n",a,a)
    fmt.Printf("&a 的类型是%T, 值是%v \n",&a,&a)
    fmt.Printf("ip 的类型是%T, 值是%v \n",ip,ip)
    fmt.Printf("*ip 的类型是%T, 值是%v \n",*ip,*ip)
    fmt.Printf("*&a 的类型是%T, 值是%v \n",*&a,*&a)


}
