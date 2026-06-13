package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Paste the link below")
	fmt.Println("*********************")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
