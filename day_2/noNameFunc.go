package main

import "fmt"

func main() {
    // 1.定义式m调用无参匿名函数
    func(data int) {
        fmt.Println("你的成绩：%d",data)
}(98 )


}
