package main

import "fmt"


func EvenOdd(){
    score :=78
    if score >=90 {
        fmt.Println("nice")
    } else if score >=80 {
        println("good")    
    } else if score >=70 {
        println("just so so")
    } else if score >= 60 {
        println("及格")
    }
}
func main() {
    EvenOdd()

}
/*
if else if  else 
总结:
    go编程中:
    else:
    if 后面的表达式可以不用不加括号,但是必须加 花括号{}
    左右花括号::都不能单独占用一行
    即使,后面只跟有一条语句,也必须加花括号...但是花括号不能单独一行..

    还有: if ;else if ;只能选择其一的分支...而且选择最近的一条分支,然后就pass过if-else语句

    60 70之间?80?score>=80

*/
