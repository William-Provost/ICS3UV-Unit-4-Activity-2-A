// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-27
// Fileoverview: This program calculates average.

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var marksAsString string
	var markAsNumber int
	var sum int
	var average float64
	var numberOfMarks int
	var err error

	reader := bufio.NewReader(os.Stdin)

	// input number
	fmt.Print("Please enter the test mark. To quit enter a -1: ")
	marksAsString, _ = reader.ReadString('\n')
	marksAsString = strings.TrimSpace(marksAsString)
	markAsNumber, err = strconv.Atoi(marksAsString)

	// keep looping until the user enters -1 to quit
	for markAsNumber != -1 {
		// check if the input is a valid number AND between 0 and 100
		if err == nil && markAsNumber >= 0 && markAsNumber <= 100 {
			// add the mark to our running total
			sum = sum + markAsNumber
			// count this as a valid mark
			numberOfMarks++
		} else if err == nil {
			// if it's a number but not in the valid range (0-100)
			fmt.Println("Invalid mark. Please enter a mark between 0 and 100.")
		} else {
			// if the input is text (not a number at all)
			fmt.Println("Invalid input. Text will not be counted.")
		}

		// ask for the next mark
		fmt.Print("Please enter the test mark. To quit enter a -1: \n")
		marksAsString, _ = reader.ReadString('\n')
		marksAsString = strings.TrimSpace(marksAsString)
		markAsNumber, err = strconv.Atoi(marksAsString)
	}

	// calculate average
	if numberOfMarks > 0 {
		average = float64(sum) / float64(numberOfMarks)
		fmt.Println()
		fmt.Println("The average is: " + fmt.Sprintf("%.2f", average) + "%")
		fmt.Println("The number of marks entered is " + fmt.Sprintf("%d", numberOfMarks))
	} else {
		fmt.Println("No marks were entered.")
	}

	fmt.Println("\nDone.")
}
