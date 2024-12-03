package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//  open the file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// pattern, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	pattern := regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	result := 0
	performMultiplication := true

	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindAllStringSubmatch(line, -1)
		fmt.Println(matches)
		for i := 0; i < len(matches); i++ {
			if matches[i][0] == "do()" {
				performMultiplication = true
			} else if matches[i][0] == "don't()" {
				performMultiplication = false
			} else {
				if performMultiplication {
					temp1, _ := strconv.Atoi(matches[i][1])
					temp2, _ := strconv.Atoi(matches[i][2])
					result += temp1 * temp2
				}
			}
		}
	}
	fmt.Println(result)
}
