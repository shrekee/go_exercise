package main

import "fmt"
import "math"
import "time"

func main() {
    t1 := time.Now()
    count := printPrimeNumber()
    t2 := time.Since(t1)
    fmt.Printf("40万自然数内有 %d 个素数，耗时：%s",count,t2)


}
/*
func eludeFourGoto() {
    print("hello")

}
*/
func printPrimeNumber() int {
    fmt.Println("1~400000的素数（借助goto跳转）")
    num := 0
    var count int = 0
LOOP:
    for num <= 400000 {
        num++
        for i := 2; i <= int(math.Sqrt(float64(num)))+1; i++ {
            if num % i == 0 {
                goto LOOP 
            }
        }
        count++
    
    
    
    }
    return count


}
