package main

import "fmt"

func main() {
   fmt.Println(max(12,22))
  // println(hello(23,11))
  x,y,z := hello(23,11)
  println(x,y,z)


}

func max(n1,n2 int) int{
    var result int

    if (n1 > n2) {
        result = n1
    } else {
        result = n2
    }
    return result
}

func hello(n1,n2 int) (int, int, int){
    var result int

    if (n1 < n2) {
        result = n2
    } else {
        result = n1
    }
    return result,n1,n2
}
