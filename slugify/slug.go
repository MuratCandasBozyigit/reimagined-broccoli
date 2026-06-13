package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Paste the link below")
	fmt.Println("*********************")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputTL := strings.ToLower(input)
	dash := strings.ReplaceAll(inputTL, " ", "-")
	replacer := strings.NewReplacer(
		"ç", "c",
		"ğ", "g",
		"ı", "i",
		"ö", "o",
		"ş", "s",
		"ü", "u",
	)
	done := replacer.Replace(dash)
	fmt.Println(done)
}
