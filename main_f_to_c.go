package main

import "fmt"

func main() {
	var f float64
	fmt.Println("Capture grados en Fahrenheit: ")
	fmt.Scanf("%f", &f)

	result := (f - 32) * 5 / 9
	if result > 32 {
		fmt.Println("Hace calor")
	}
	if result < 0 {
		fmt.Println("Hace frio")
	}

	fmt.Println("El resultado en Celcius es:", result)
}
