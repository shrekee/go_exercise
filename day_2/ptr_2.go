package main

import "fmt"

func main() {
    a := 10
    b := &a
    fmt.Printf("b的类型：%T， 数值是%v\n",b,b)
    fmt.Printf("a的地址：%d\n",b)
    fmt.Printf("*b的值：%d\n",*b)

    *b++
    fmt.Printf("a的地址：%d\n",b)



}
