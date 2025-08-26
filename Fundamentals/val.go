package main

import "fmt"

// Global variable declaration
var my_val int = 100 // Global variable

func main() {
	my_val1 := 20
	my_val2 := "TranDoanMinh"
	my_val3 := 34.84
	var localVar int = 200       // Local variable
	fmt.Printf("%d\n", localVar) // Accessible here
	const PI = 3.14
	fmt.Printf("The value of myvariable1 is : %d\n", my_val1)
	fmt.Printf("The type of myvariable1 is : %T\n", my_val1)

	fmt.Printf("The value of myvariable2 is : %s\n", my_val2)

	fmt.Printf("The type of myvariable2 is : %T\n", my_val2)

	fmt.Printf("The value of myvariable3 is : %f\n", my_val3)

	fmt.Printf("The type of myvariable3 is : %T\n", my_val3)

	fmt.Printf("%f \n", PI)
	const flag = true
	fmt.Println(flag)
}
