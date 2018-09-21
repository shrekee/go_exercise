	package main
	//import "fmt"
	type Student struct {
		name string
		age uint16
		married bool
		gender uint8
	}
	func main() {
		s1 := Student{"Steven",30,true,'f'}
		s2 := Student{"sunny",20,false,'m'}

		var a *Student = &s1
		var b *Student = &s2
		println(s1.gender)
		println(a.age,b.gender)
}
