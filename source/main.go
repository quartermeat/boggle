package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// for _, eachline := range txtlines {
	// 	fmt.Println(eachline)
	// }
	run := true
	for run {
		run = execute(txtlines)
	}
	fmt.Print("exiting gracefully")
}

func CheckInput(input string, txtLines []string) bool {

	return true
	// for _, line := range txtLines {

	// }
}

func execute(txtlines []string) bool {

	// var input string

	for _, txtline := range txtlines {
		fmt.Println(txtline)
		// fmt.Println("enter a word:")
		// fmt.Scanln(&input)
		var cont = "y"
		if !CheckInput(cont, txtlines) {
			break
		}
	}
	return false
}
