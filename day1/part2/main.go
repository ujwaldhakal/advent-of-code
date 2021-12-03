package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main()  {

	data := getInput()

	var lastSumIndex int
	numOfTimesIncreased := 0


	totalDatalength := len(data)
	for i := 0; i < totalDatalength; i++ {

		lastIndex := totalDatalength - 3
		if i == lastIndex {
			return
		}

		lastSum := data[i+1] + data[i+2] + data[i+3]

		if lastSum > lastSumIndex {
			numOfTimesIncreased = numOfTimesIncreased + 1
		}

		fmt.Println(numOfTimesIncreased)
		lastSumIndex = lastSum

	}

	fmt.Println(numOfTimesIncreased)

}



func getInput() []int {
	var input []int

	f, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {

		currentNum,_ := strconv.Atoi(sc.Text())
		input = append(input, currentNum)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		panic(err)
	}

	return input
}