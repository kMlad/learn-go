package concurrency

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const basePath = "./assets/goroutines"

func GoRoutines() {

	totalSum := 0
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	getNumbersFromFile := func(path string) {
		sum := 0
		file, err := os.Open(fmt.Sprintf("%v/%v", basePath, path))
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			num, err := strconv.Atoi(line)
			if err != nil {
				continue
			}
			sum += num
		}

		fmt.Printf("Sum for %v: %v\n", path, sum)
		totalSum += sum
	}

	for _, file := range files {
		go getNumbersFromFile(file)
	}

	time.Sleep(1000 * time.Millisecond)

	fmt.Printf("Total sum: %v\n", totalSum)

}
