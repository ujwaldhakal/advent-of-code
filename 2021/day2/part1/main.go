package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {

	totalHorizontal, totalDepth := getTotalHorizontalAndDepth()
	fmt.Println("did it", totalHorizontal * totalDepth)
}


type position struct {

	direction string
	point int
}

func getTotalHorizontalAndDepth() (int,int) {

	var totalHorizontalPosition int
	var totalDepth int
	f, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {

		explodedStr := strings.Split(sc.Text()," ")

		direction := explodedStr[0]
		point,_ := strconv.Atoi(explodedStr[1])

		if direction == "forward" {
			totalHorizontalPosition = totalHorizontalPosition + point
		}

		if direction == "down" {
			totalDepth = totalDepth + point
		}

		if direction == "up" {
			totalDepth = totalDepth - point
		}
		//input = append(input, position{direction: explodedStr[0], point: point})
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		panic(err)
	}

	return totalHorizontalPosition, totalDepth
}