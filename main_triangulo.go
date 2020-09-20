package main

import "fmt"

func main(){
	var base float64
	var heigt float64
	fmt.Println("Capture base del triangulo: ")
	fmt.Scan(&base)
	fmt.Println("Capture altura del triangulo: ")
	fmt.Scan(&heigt)

	result := (base * heigt) / 2

	fmt.Println("El area es:", result)
}