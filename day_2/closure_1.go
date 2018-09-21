package main

import "fmt"

func main() {
    res := Count()
    fmt.Printf("%T\n", res)
    fmt.Println("res:", res)
    fmt.Println("res():", res())
    fmt.Println("res():", res())
    fmt.Println("res():", res())
    res2 := Count()
    fmt.Println("res2:", res2)
    fmt.Println("res2():", res2())
    fmt.Println("res2():", res2())
    fmt.Println("res2():", res2())
    //把res第四次调用。
    fmt.Println("res():", res())

}

//闭包函数，实现计数器功能

func Count() func() int {
    i := 0
    res := func() int {
        i++
        return i
    }
    fmt.Println("Counter内部：",res)
    return res
}
