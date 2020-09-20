package main

import "fmt"

func main(){
	var side float64
	fmt.Println("Capture lado del cuadrado: ")
	fmt.Scanf("%f", &side)

	result := side * side

	fmt.Println("El area es:", result)
}