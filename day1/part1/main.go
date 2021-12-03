package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	increasedTimes,_ := findIncreasingPattern()
	fmt.Println(increasedTimes)
}

func findIncreasingPattern() (int,error) {
	var timesOfIncrease int
	var lastOccurance int

	f, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {

		currentNum,_ := strconv.Atoi(sc.Text())
		if(lastOccurance > 0 && currentNum > lastOccurance ) {
			timesOfIncrease++
		}

		lastOccurance = currentNum
		_ = sc.Text() // GET the line string
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		panic(err)
	}

	return timesOfIncrease,err
}
