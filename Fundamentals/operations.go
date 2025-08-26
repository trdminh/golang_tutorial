package main

import "fmt"

func main() {
	q := 23
	p := 10

	fmt.Printf("Result of %d + %d = %d\n", p, q, q+p)
	fmt.Printf("Result of %d - %d = %d\n", p, q, p-q)
	fmt.Printf("Result of %d * %d = %d\n", p, q, p*q)
	fmt.Printf("Result of %d / %d = %d\n", p, q, q/q)
	fmt.Printf("Result of %d %% %d = %d\n", p, q, p%q)
}
