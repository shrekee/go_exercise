package main
import "time"

func Generate(ch chan<- int) {
for i := 2; ; i++ {
ch <- i
}

}

func Filter(in <-chan int, out chan<- int, prime int) {
for {
i := <-in
if i%prime != 0 {
out <- i
}
}

}

func main() {
t1 := time.Now() //获得当前时间
ch := make(chan int)
go Generate(ch)

for i := 0; i < 10000; i++ {
prime := <-ch
//print(prime, "\n")
ch1 := make(chan int)
go Filter(ch, ch1, prime)
ch = ch1
  }
  elapsed := time.Since(t1) //经过的时间。。
  println("40e万的素数耗时：",elapsed)
}
