package main

//import  "fmt"

func main() {
    var age uint8 = 20

    switch age {
        case age >10:
                println("大于10")
        case age > 20:
                println("大于20")
        case age > 30:
                println("大于30")
    }



}
//报错： int 和 bool型数据不匹配。。。
//所以：switch后 表达式的值要和 caseh 后跟的值要数据类型一致！！！
// switch expression   {  case val1 :   case val2 : } default :
// 关键字  ： faultthrough

// 如果在case中出现了 faultthrough y关键字，那么本case语句会贯穿到下一条case中，然后再下一条中case 自带了“break”
//所以，带有faultthrough，会跑到下一条语句中。
//
// faultrthrough  只能出现再case语句的最后一行。。。。
//
