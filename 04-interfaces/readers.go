package interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Readers() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello there! How can I help you?")
	commandsCtr, nonBlankLinesCtr := 0, 0
	for {
		input, inputErr := reader.ReadString('\n')
		if inputErr != nil {
			panic(fmt.Sprint("Input error:", inputErr))
		}
		commandsCtr++
		word := strings.TrimSpace(input)
		if word == "" {
			continue
		}
		nonBlankLinesCtr++
		if strings.ToLower(word) == "hello" {
			fmt.Println("Greetings!")
		}
		if strings.ToLower(word) == "bye" {
			fmt.Println("Goodbye!")
		}
		if strings.ToLower(word) == "q" {
			fmt.Println("No. of commands: ", commandsCtr)
			fmt.Println("No. of non-blank lines: ", nonBlankLinesCtr)
			break
		}
	}
}
