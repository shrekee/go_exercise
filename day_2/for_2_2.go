package main

//import "fmt"
var lines = 9
func main(){
//打印矩形
    printRectangle()
//打印直角三角形
    printRightTriangle()
//打印右下直角三角形。。
    printLeftTriangle()

}
//打印矩形
func printRectangle() {
    for i := 1; i <= 9;i++ {
        for j := 0 ;j<=9;j++ {
            print("-")
        }
        println()
    }
}
//打印左下角三角形
func printRightTriangle() {
    for i := 1;i <= 10; i++ {
        for j :=1; j<= 10-i ; j++ {
            print("-")
        } 
        println()
    }
}
//打印右下直角三角形。。。
func printLeftTriangle() {
    for i := 0; i <= 10; i++ {
        for j:=0;j<= 10; j++ {
            if j < i {
                print(" ")
            } else {
                print("*")
            }
        }
        println()
    }
}
