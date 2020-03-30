package main

import "fmt"

func main() {

	// var numberA int = 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value)   :", numberA)  // 4
	// fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	// fmt.Println("numberB (value)   :", *numberB) // 4
	// fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	//============= Efek Perubahan Nilai Pointer
	// var numberA int = 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value)   :", numberA)
	// fmt.Println("numberA (address) :", &numberA)
	// fmt.Println("numberB (value)   :", *numberB)
	// fmt.Println("numberB (address) :", numberB)

	// fmt.Println("")

	// numberA = 5

	// fmt.Println("numberA (value)   :", numberA)
	// fmt.Println("numberA (address) :", &numberA)
	// fmt.Println("numberB (value)   :", *numberB)
	// fmt.Println("numberB (address) :", numberB)

	//================== menggunakan parameter

	var a int = 11
	var p *int = &a
	var pp **int = &p
	var ppp ***int = &pp

	fmt.Println(a)
	fmt.Println(&a)

	fmt.Println(p)

	fmt.Println("================")
	fmt.Println(&p)

	fmt.Println(pp)
	fmt.Println("================")
	fmt.Println(*pp)
	fmt.Println(**pp)
	fmt.Println(***ppp)

	fmt.Println("================")
	change(&a, 18)
	fmt.Println(a)
	fmt.Println("")
	// param
	var number = 9
	fmt.Println("before :", number) // 4

	change(&number, 15)
	fmt.Println("after  :", number) // 10

	//bbbbbbbbbbbbbbbbbbbbbbbb
	// var a int = 20
	// var ip *int

	// ip - &a

	// fmt.Println("Adress of a variable: %x", &a)
	// fmt.Println("Adress stored in ip variable: %x", ip)
	// fmt.Println("Value of *ip variable: %x", *ip)

}

func change(original *int, value int) {
	*original = value
}
