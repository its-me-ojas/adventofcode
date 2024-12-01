package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// open the file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var firstList, secondList []int
	scanner := bufio.NewScanner(file)

	// reading line by line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) != 2 {
			fmt.Println("Invalid line")
			continue
		}

		first, err1 := strconv.Atoi(parts[0])
		second, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting line: ", line)
		}

		firstList = append(firstList, first)
		secondList = append(secondList, second)

	}

	// sorting both lists
	sort.Ints(firstList)
	sort.Ints(secondList)

	ans := 0
	for i := 0; i < len(firstList); i++ {
		// applying logic given in the question
		ans += int(math.Abs(float64(firstList[i] - secondList[i])))
	}

	fmt.Println(ans)

}
