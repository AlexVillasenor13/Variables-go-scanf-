package main

import (
	"fmt"
	"math"
)

func main(){
	var r float64
	fmt.Println("Capture radio del circulo: ")
	fmt.Scanf("%f", &r)

	result := math.Pi * math.Pow(r, 2)

	fmt.Println("El area es:", result)
}