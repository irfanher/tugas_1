package main

import "fmt"

type Poin struct {
	X, Y float64
}

func (p Poin) IsAbove(y float64) bool {
	return p.Y > y
}

func main() {
	p := Poin{2.0, 4.0}

	fmt.Println("Poin : ", p)

	fmt.Println("is Point p located above the line y = 1,0 ? : ", p.IsAbove(1))
}
