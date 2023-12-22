package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите значения a, b и c:")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)
	fmt.Print("c = ")
	fmt.Scanln(&c)

	calcSum(&a, &b, &c)
}

func calcSum(a *int, b *int, c *int) {
	NewA := *a
	NewB := *b
	*c = *a + *b

	fmt.Println("Копия A: =", NewA)
	fmt.Println("Копия B: =", NewB)
}
