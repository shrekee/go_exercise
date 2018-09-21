package main

import "fmt"

func main() {
    //d函数的可变参数。。
    println(AddSum(1,2,3,4,5,6))
    println(AddSum())
    arr := []int{2,3,52,33,345,343}
    println(AddSum(arr...))
    //println(len(int))
    var flo float32
    var str string
    println(flo,"\thaha",str,"hello",str,"H")

}

func AddSum(nums ...int) (sum int) {
        fmt.Printf("%T\n", nums)
        for _, value := range nums {
            sum += value 
        }
        return 
}
