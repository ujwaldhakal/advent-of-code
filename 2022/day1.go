package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main()  {

	calories := sortedCalories()
	fmt.Println("highest calorie",calories[0] );
	fmt.Println("top 3 total highest calorie",calories[0] + calories[1] + calories[2]);
}

func readFileLineByLine(fileName string) (*os.File, *bufio.Scanner) {
	data,err := os.Open(fileName)

	if err != nil {
		log.Fatalln("something went wrong")
	}


	scanner := bufio.NewScanner(data)

	return data,scanner
}

func sortedCalories() []int {
	file,input := readFileLineByLine("input.txt")
	defer file.Close()

	var highestCalorie []int

	totalCalorie := 0
	for input.Scan() {

		if input.Text() != "" {
			calorie, err := strconv.Atoi(input.Text())
			if err != nil {
				log.Fatalln("could not convert calorie to int")
			}
			totalCalorie = totalCalorie +  calorie
		}

		if input.Text() == "" {
			highestCalorie = append(highestCalorie,totalCalorie)
			totalCalorie = 0
		}
	}

	sort.Slice(highestCalorie, func(a, b int) bool {
		return highestCalorie[b] < highestCalorie[a]
	})

	return highestCalorie

}

