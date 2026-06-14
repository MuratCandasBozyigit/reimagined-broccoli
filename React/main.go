package main

import "fmt"

func main() {
	num := 0

	// Parantez yok, noktalı virgül var ve i++ eklendi!
	for i := 0; i <= 1000000000; i++ {
		num++
	}

	fmt.Println("İşlem bitti! Num değeri:", num)
}
