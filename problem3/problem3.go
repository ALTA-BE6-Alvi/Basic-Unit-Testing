package main

import "fmt"

func PrintNama(nama string) string {
	// your code here
	result := "Hello" + " " + nama +"! Saya Golang bahasa yang sangat menyenangkan"
	return result
}

func main() {
	var nama string = "kobar"
	fmt.Println(PrintNama(nama))
}
