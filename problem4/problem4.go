package main

import "fmt"

func PrintLuas(alas, tinggi float64) float64 {
	// your code here
	const pi = 3.14  
	var luas float64 = alas * tinggi / 2

	return luas
}

func main() {
	var alas float64 = 20
	var tinggi float64 = 25

	fmt.Println(PrintLuas(alas, tinggi))
}
