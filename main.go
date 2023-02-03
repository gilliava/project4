// Names: Mary Awopileda, Vincent Gilliam, Myles Robinson
// VUnetIDs: awopilma, gilliava, robinmb4
// Emails: mary.a.awopileda@vanderbilt.edu, vincent.a.gilliam@vanderbilt.edu,
// myles.b.robinson@vanderbilt.edu
// Class: CS 3270 - Vanderbilt University
// Date: 12/06/2022
// Honor statement: We have neither given nor received aid from any unauthorized source
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * Determines if a number can be placed in the corresponding spot
 *
 * @param numToCheck the number to attempt to place
 * @param row the row of the empty spot
 * @param col the column of the empty spot
 * @param board the 2d array representing the sudoku board
 * @return true if the number can be placed and false if it cannot
 */

func canPlace(numToCheck int, row int, col int, board [9][9]int) bool {

	// check if the number already exists in the row
	for i := 0; i < 9; i++ {

		// skip over the spot that is being tried
		if i == col {

		} else if board[row][i] == numToCheck {
			return false
		}

	}

	// check if the number already exists in the column
	for i := 0; i < 9; i++ {

		// skip over the spot that is being tried
		if i == row {

		} else if board[i][col] == numToCheck {
			return false
		}

	}

	// get the first row to check in a 3x3 square
	var beginningRow = row - (row % (9 / 3))

	// get the second row to check in a 3x3 square
	var beginningColumn = col - (col % (9 / 3))

	// check if the number already exists in the 3x3 square
	for i := 0; i < (9 / 3); i++ {

		for j := 0; j < (9 / 3); j++ {

			if board[beginningRow+i][beginningColumn+j] == numToCheck {
				return false
			}
		}

	}

	// return true if all the specifications pass
	return true
}

/**
 * Determines whether a puzzle is solvable or not
 *
 * @param board the 2d array representing the sudoku board
 * @return true if the puzzle is solvable and false if it is not solvable
 */

func solve(board [9][9]int) bool {

	// stores the ith index of the last open slot
	var lastIndexOfi = -1

	// stores the jth index of the last open slot
	var lastIndexOfj = -1

	// gets the ith and jth indexes of the last open slot
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				lastIndexOfi = i
				lastIndexOfj = j
			}
		}
	}

	// puzzle is solved when every spot has been filled
	if lastIndexOfj == -1 && lastIndexOfi == -1 {
		fmt.Println("Solution:\n")
		printBoard(board)
		return true
	}

	// keeps track of the number to try and insert
	var numToInsert = 1

	// traverses through the puzzle
	for i := 0; i < 9; i++ {

		for j := 0; j < 9; j++ {

			// tries a number if the spot is equal to 0
			if board[i][j] == 0 {

				// executes if the number can be inserted
				if canPlace(numToInsert, i, j, board) {

					// places the number for that spot
					board[i][j] = numToInsert

					// recursively calls itself on the next open spot
					if solve(board) {

						return true
					}

					// sets the spot back to 0
					board[i][j] = 0

					// tries the next number in the sequence
					j--
					numToInsert++

					// exits when all possible numbers have been exhausted
					if numToInsert == 10 {
						return false
					}

				} else if !canPlace(numToInsert, i, j, board) {

					// tries the next number in the sequence
					j--
					numToInsert++

					// exit if you hit the limit of numbers
					if numToInsert == 10 {
						return false
					}
				}
			}
		}
	}

	// returns false if the puzzle can not be solved
	return false
}

/**
 * Prints the board in its current state
 *
 * @param board the 2d array representing the sudoku board
 */

func printBoard(boardToPrint [9][9]int) {

	//string that will hold the empty sudoku board format
	var board = ""

	// keeps track of when to add a bar and new line
	var barCounter = 1

	//print the empty board
	for k := 0; k < 9; k++ {

		for l := 0; l < 9; l++ {

			// places a space and bar where necessary if the number to insert is 0
			if boardToPrint[k][l] == 0 && ((barCounter == 3) || (barCounter == 6)) {
				board += "  | "
				barCounter++
			} else if boardToPrint[k][l] == 0 && (barCounter != 3 && barCounter != 6) {
				board += "  "
				barCounter++
			} else if boardToPrint[k][l] != 0 && ((barCounter == 3) || (barCounter == 6)) {

				// converts the integer to a string
				intToAdd := strconv.Itoa(boardToPrint[k][l])
				board += intToAdd + " | "
				barCounter++
			} else if boardToPrint[k][l] != 0 && (barCounter != 3 && barCounter != 6) {

				// converts the integer to a string
				intToAdd := strconv.Itoa(boardToPrint[k][l])
				board += intToAdd + " "
				barCounter++
			}

			// places a space and newline where necessary if the number to insert is 0
			if l == 8 {
				barCounter = 1
				board += " \n"
			}

		}
		// places the divider line
		if k == 2 || k == 5 {
			board += "------+-------+------\n"
		}

	}

	//print the empty board in sudoku format
	fmt.Println(board)
}

/**
 * Reads the file that is indicated from input and converts the string representation
 * into a 2d array representation of a sudoku puzzle
 *
 * @param sudokuBoard the 2d array representing the sudoku board
 */

func loadFile(sudokuBoard *[9][9]int) {

	// gets the name of the text file
	var filename string
	fmt.Scanln(&filename)

	// default to first puzzle if the user doesn't enter anything
	if len(filename) == 0 {
		filename = "txt/sudoku-test1.txt"
	} else {
		filename = "txt/" + filename
	}

	//keeps track of the index to insert the number at
	var i = 0
	var j = 0

	//opens the text file
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	//closes the text file
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	// load the text file into the array
	for scanner.Scan() {

		// convert the string character to an int
		intVar, err := strconv.Atoi(scanner.Text())

		// if there is an error, print it
		if err != nil {
			fmt.Println(err)
		}

		// if the end of the row is reached, reset the column # and increment the row #
		if j == 9 {
			i++
			j = 0
		}

		// place the character in the board
		sudokuBoard[i][j] = intVar

		// increment to the next column
		j++
	}

	// if there is an error, print it
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {

	// Ask the user for the text file name
	fmt.Println("Enter puzzle text file (assumes file is in \"txt\" folder).")
	fmt.Println("Pressing <Enter> will run the file \"sudoku-test1.txt\".\n")
	fmt.Println("\n")

	// 2d array to represent the sudoku boar
	var board [9][9]int

	// creates a pointer to the board
	var ptr *[9][9]int
	ptr = &board

	// loads the text file into the array
	loadFile(ptr)

	// prints the empty board
	fmt.Println("Puzzle:\n")
	printBoard(board)

	// begin trying to solve the board
	if solve(board) {

	} else {
		//if there is no solution, print it to the console
		fmt.Println("Solution:\n")
		fmt.Println("No solution")
	}
}
