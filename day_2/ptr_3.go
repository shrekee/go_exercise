package main

import "fmt"

func main() {
    a := 10
    fmt.Println("函数调用前的a值",a)

    b := &a
    change(b)
    fmt.Println("函数调用后的a值",a)



}
func change(num *int) {
    *num++
}
