package day3

import (
	"fmt"
	"github.com/ujwaldhakal/advent-of-code/day1"
	"strings"
)

func SolvePuzzle()  {
	//points := mapLettersToPoints()
	//puzzle1(points)

	puzzle2()
}

func puzzle2()  {
	file,input := day1.ReadFileLineByLine("day3/input.txt")
	defer file.Close()
	sum := 0
	for input.Scan() {
		row1 := input.Text()

		input.Scan()
		row2 := input.Text()

		input.Scan()
		row3 := input.Text()

		saw := map[int32]bool{}
		for _, v1 := range row1 {

			foundIn2 := false
			foundIn3 := false

			if saw[v1] {
				continue
			}

			saw[v1] = true

			for _, v2 := range row2 {
				if v2 == v1 {
					foundIn2 = true
					break
				}
			}

			for _, v3 := range row3 {
				if v3 == v1 {
					foundIn3 = true
					break
				}
			}

			if foundIn2 && foundIn3 {
				switch {
				case v1 >= 97 && v1 <= 122:
					sum += int(v1) - 97 + 1

				case v1 >= 65 && v1 <= 90:
					sum += int(v1) - 65 + 26 + 1
				}
				break
			}
		}
	}

	fmt.Println("this is puzle 2",sum)
}

func findRepeatedString(data []string) map[string]bool {

	firstHashMap :=map[string]bool{}


	fmt.Println("tota",len(data))
	for _,char := range data[0] {
		firstHashMap[string(char)] = true
	}


	//repeatedString := []string{}
	for _,char := range data[1] {
		//str := string(char)
		firstHashMap[string(char)] = true
		//if firstHashMap[str]  {
		//	firstHashMap[string(char)] = true
		//	//repeatedString = append(repeatedString,str)
		//}
	}


	finalString := []string{}
	for _,char := range data[2] {
		str := string(char)
		if firstHashMap[str]  {
			finalString = append(finalString,str)
		}
	}

	fmt.Println("map",finalString)

	return firstHashMap
}



func mapLettersToPoints() map[string]int {

	letter := map[string]int{}
	startingPointForCapsLetter := 26
	for i := 1; i < 27; i++ {
		char := string(rune(64+i))
		letter[char] = startingPointForCapsLetter + i
		letter[strings.ToLower(char)] = i
	}

	return letter
}

func puzzle1(points map[string]int)  {

	file,input := day1.ReadFileLineByLine("day3/input.txt")
	defer file.Close()

	occurance := map[string]int{}
	totalPoints := 0
	for input.Scan() {
		stringLen := len(input.Text())
		firstHalf := input.Text()[0:stringLen/2]
		secondHalf := input.Text()[stringLen/2:stringLen]
		for _,firstHalfLetter := range firstHalf {
			for _,secondHalfLetter := range secondHalf {
				if firstHalfLetter == secondHalfLetter {
					occurance[string(secondHalfLetter)] = occurance[string(secondHalfLetter)]
				}
			}

		}
		for char,_ := range occurance {
			totalPoints = totalPoints + points[char]
		}
		occurance = map[string]int{}
	}



	fmt.Println("sum of priorities is",totalPoints)
}