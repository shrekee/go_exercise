package main

//import "fmt"

func main(){

/*
    iota:  每当遇到 go语言的关键字后,重置为 零!!!零!零!
*/
    const (
        L = iota
        M = iota
        N = iota
    )
    println(L,M,N)
    

}
