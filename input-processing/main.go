package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the lines")
	line, _ := reader.ReadString('\n')

	sentences := strings.Split(line, ".")
	for _, x := range sentences {

		temp := strings.Split(x, " ")
		for _, words := range temp {
			if words == "error" {
				fmt.Println(x)
			}
		}
	}

}
