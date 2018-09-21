package main

func main(){
    const (
        i = 1 << iota    //   << 是位移运算符
        j = 3 << iota
        k
        l
    )
    println(i,j,k,l)



}
