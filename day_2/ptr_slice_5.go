package main

import "fmt"

func main() {
    a := []int{1,2,3,4}
    fmt.Print("调用前的切片的值:",a)
    modify(&a)
    fmt.Print("调用后的切片的值:",a)
    fmt.Println(a)

}
func modify( arr *[]int ) {
    (*arr)[0] = 250
}
