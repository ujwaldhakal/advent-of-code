package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func SolvePuzzle()  {

	fmt.Println("got here")
	calories := sortedCalories()
	fmt.Println("highest calorie",calories[0] );
	fmt.Println("top 3 total highest calorie",calories[0] + calories[1] + calories[2]);
}

func ReadFileLineByLine(fileName string) (*os.File, *bufio.Scanner) {
	data,err := os.Open(fileName)

	if err != nil {
		log.Fatalln("something went wrong reading file",err)
	}


	scanner := bufio.NewScanner(data)

	return data,scanner
}

func sortedCalories() []int {
	file,input := ReadFileLineByLine("day1/input.txt")
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

