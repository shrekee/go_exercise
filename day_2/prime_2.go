package main

import (
    "fmt"
    "math"
    "time"
)

func main() {
    t1 := time.Now()
    count := 1
    for i:= 3;i <= 400000; i+=2 {
        k := int(math.Sqrt(float64(i)))+1
        for j:=2; j <=k;j++ {
            if i % j ==0 {
                break
            }
            if j ==k {
                count++   
            }
        
        }
    
    
    }
    t2 := time.Since(t1)
    fmt.Printf("40万素数有 %d 个， 耗时：%s \n",count,t2)

}
