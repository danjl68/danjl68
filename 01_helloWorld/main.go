package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("This is a test!")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
