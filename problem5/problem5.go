package main

import "fmt"

func PrintLuasPermukaan(radius, tinggi float64) float64 {
	// your code here
	const pi = 3.14  
	var luasPermukaan float64 = 2 * pi * radius * ( radius + tinggi )
	
	return luasPermukaan
}

func main() {
	var T float64 = 20
	var r float64 = 4

	fmt.Println(PrintLuasPermukaan(r, T))
}
