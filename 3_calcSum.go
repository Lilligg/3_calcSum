package main

import "fmt"

type Number struct{}

func (s Number) calcSum(a, b, c int) {
	newA := a
	newB := b
	c = a + b

	fmt.Println("Копия A:", newA)
	fmt.Println("Копия B:", newB)
	fmt.Println("Измененная C:", c)
}

func main() {
	var a, b, c int
	fmt.Println("Введите значения a, b и c:")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)
	fmt.Print("c = ")
	fmt.Scanln(&c)

	result := Number{}
	result.calcSum(a, b, c)
}
