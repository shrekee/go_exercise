package main

import "fmt"

func main() {
    a, b := 100, 200
    a,b = swap(a,b)
    fmt.Println("第一次交换后\t",a,b)
    swap_1(&a,&b)
    fmt.Println("第二次交换后\t",a,b)


}
//具有返回值的管用方法，实现两个数据的交换
func swap(x,y int) (int,int) {
    return y, x

}
//使用指针作为参数的写法
func swap_1(x, y *int) {
    *x, *y = *y, *x

}
