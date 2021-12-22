package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	//open file
	file, err := os.Open("resources/en_US.dic")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	//create scanner from file
	scanner := bufio.NewScanner(file)
	//split by scan lines
	scanner.Split(bufio.ScanLines)

	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	run := true
	board := getBoard()
	for run {
		run = runBoggle(board, txtlines)
	}
	fmt.Print("exiting gracefully")
}

func listContains(match string, wordList []string) bool {
	for _, txtline := range wordList {
		stringSections := strings.Split(txtline, "/")
		word := stringSections[0]
		if word == match {
			return true
		}
	}
	return false
}

func getBoard() [4][4]string {
	//taken from: https://boardgames.stackexchange.com/questions/29264/boggle-what-is-the-dice-configuration-for-boggle-in-various-languages
	availableDice := [][]string{
		// get letters
		{"I  ", "F  ", "E  ", "H  ", "E  ", "Y  "},
		{"R  ", "I  ", "F  ", "O  ", "B  ", "X  "},
		{"U  ", "T  ", "O  ", "K  ", "N  ", "D  "},
		{"D  ", "E  ", "N  ", "O  ", "W  ", "S  "},
		{"H  ", "M  ", "S  ", "R  ", "A  ", "O  "},
		{"L  ", "U  ", "P  ", "E  ", "T  ", "S  "},
		{"Y  ", "L  ", "G  ", "K  ", "U  ", "E  "},
		{"QU ", "B  ", "M  ", "J  ", "O  ", "A  "},
		{"E  ", "H  ", "I  ", "S  ", "P  ", "N  "},
		{"V  ", "E  ", "T  ", "I  ", "G  ", "N  "},
		{"B  ", "A  ", "L  ", "I  ", "Y  ", "T  "},
		{"E  ", "Z  ", "A  ", "V  ", "N  ", "D  "},
		{"R  ", "A  ", "L  ", "E  ", "S  ", "C  "},
		{"U  ", "W  ", "I  ", "L  ", "R  ", "G  "},
		{"P  ", "A  ", "C  ", "E  ", "M  ", "D  "},
		{"A  ", "C  ", "I  ", "T  ", "O  ", "A  "},
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(availableDice), func(i, j int) { availableDice[i], availableDice[j] = availableDice[j], availableDice[i] })

	var board [4][4]string

	for i, column := range board {
		for j, _ := range column {
			board[i][j], availableDice = getRandomLetter(availableDice)
		}
	}
	return board

}

func getRandomLetter(allDice [][]string) (string, [][]string) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(allDice))
	j := rand.Intn(6)
	returnString := allDice[i][j]
	allDice[i] = allDice[len(allDice)-1]
	return returnString, allDice[:len(allDice)-1]
}

func printBoard(board [4][4]string) {
	for i, column := range board {
		for j, _ := range column {
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}
}

func isWordAvailable(board [4][4]string, word string) bool {
	//pick up all the words on the board

}

func runBoggle(board [4][4]string, txtlines []string) bool {

	var input string
	screen.Clear()
	printBoard(board)

	fmt.Print("Enter word to check: ")
	fmt.Scanln(&input)

	if !isWordAvailable(board, input) {
		return false
	}

	if listContains(input, txtlines) {
		fmt.Println("	yes")
	} else {
		fmt.Println("	no")
	}

	return input != "q"
}
