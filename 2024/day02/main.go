package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Check if the sequence is strictly increasing
func isIncreasing(parts []int) bool {
	for i := 0; i < len(parts)-1; i++ {
		if parts[i] >= parts[i+1] || (parts[i+1]-parts[i] > 3) {
			return false
		}
	}
	return true
}

// Check if the sequence is strictly decreasing
func isDecreasing(parts []int) bool {
	for i := 0; i < len(parts)-1; i++ {
		if parts[i] <= parts[i+1] || (parts[i]-parts[i+1] > 3) {
			return false
		}
	}
	return true
}

// Check if removing one element makes the sequence valid
func canBecomeValid(parts []int) bool {
	if len(parts) <= 2 {
		return true
	}
	for i := 0; i < len(parts); i++ {
		modified := append([]int(nil), parts[:i]...)
		modified = append(modified, parts[i+1:]...)
		if isIncreasing(modified) || isDecreasing(modified) {
			return true
		}
	}
	return false
}

func main() {
	// Open the input file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans := 0
	scanner := bufio.NewScanner(file)
	lineNumber := 1

	// Process each line in the input
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		nums := make([]int, len(parts))
		for i, value := range parts {
			num, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalf("Invalid number in line %d: %s", lineNumber, value)
			}
			nums[i] = num
		}

		// Debugging the original sequence
		// fmt.Printf("Line %d: Sequence: %v\n", lineNumber, nums)

		// Check if the sequence is already valid
		isSafe := isDecreasing(nums) || isIncreasing(nums)
		if isSafe {
			// fmt.Printf("Line %d: Safe (increasing or decreasing)\n", lineNumber)
			ans++
			lineNumber++
			continue
		}

		// Check if it can be made valid by removing one element
		if canBecomeValid(nums) {
			fmt.Printf("Line %d: Can become safe by removing an element\n", lineNumber)
			ans++
		} else {
			// fmt.Printf("Line %d: Unsafe\n", lineNumber)
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total safe sequences:", ans)
}
