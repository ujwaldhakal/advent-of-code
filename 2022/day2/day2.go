package day2

import (
	"fmt"
	"github.com/ujwaldhakal/advent-of-code/day1"
	"strings"
)

var winingRules,drawRules,losingRules  map[string]string
var points  map[string]int

func SolvePuzzle()  {

	 winingRules = map[string]string{
		"X":"C",
		"Y": "A",
		"Z":"B",
	}

	losingRules = map[string]string{
		"X":"B",
		"Y": "C",
		"Z":"A",
	}
	drawRules = map[string]string{
		"X":"A",
		"Y": "B",
		"Z":"C",
	}

	points = map[string]int{
		"X":1,
		"Y":2,
		"Z":3,
	}

	//puzzle1()
	puzzle2()
}



func puzzle1()  {

	file,input := day1.ReadFileLineByLine("day2/input.txt")



	defer file.Close()

	totalPoints := 0
	for input.Scan() {

		data := strings.Split(input.Text()," ")
		opponentMove := data[0];
		elfMove:= data[1]

		point := puzzle1Point(opponentMove,elfMove)
		totalPoints = totalPoints + point
	}

	fmt.Println("total point for first puzzle",totalPoints)
}

func puzzle2()  {
	file,input := day1.ReadFileLineByLine("day2/input.txt")

	defer file.Close()

	totalPoints := 0
	for input.Scan() {

		data := strings.Split(input.Text()," ")
		opponentMove := data[0];
		outcome:= data[1]

		point := puzzle2Point(opponentMove,outcome)
		totalPoints = totalPoints + point
	}

	fmt.Println("total point for second puzzle",totalPoints)
}

func findOutcome(rules map[string]string,opponentMove string) string {

	for key,value := range rules {

		if opponentMove == value {
			return key
		}
	}
	return ""
}


func puzzle2Point(opponentMove string, outcome string) int  {

	if outcome == "X" {
		return points[findOutcome(losingRules,opponentMove)] + 0
	}

	if outcome == "Y" {
		return points[findOutcome(drawRules,opponentMove)] + 3
	}

	if outcome == "Z" {
		return points[findOutcome(winingRules,opponentMove)] + 6
	}

	return 0
}

func puzzle1Point(opponentMove string, yourMove string) int {

	if drawRules[yourMove] == opponentMove {
		return points[yourMove] + 3
	}

	if winingRules[yourMove] == opponentMove {
		return points[yourMove] + 6
	}

	return points[yourMove] + 0 // lost

}