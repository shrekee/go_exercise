package main

import "fmt"

func main() {
    var i int
    for i < 10{
      i++
      fmt.Println(i)
    }
    for i := 0; i <= 10; i++ {
        fmt.Printf("%d ",i)
    }
    traverString()
    travelSlice()

}
func traverString() {
    str := "123ABCabc一丁丂"
    for i, value := range str {
        fmt.Printf("第%d位的字符是：%c \n",i, value)
    }
}

func travelSlice() {
    arr := []int{100,200,300}
    for  i, value := range arr {
        fmt.Println(i,": ",value)
    
    }

}
