package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// cell contains it's x and y coordinates in the world
// and it's status
type cell struct {
	posX, posY int
	status     rune
}

// generation contains
// the state of the world at it's point in time,
// the length of the square representing the world
// and the number of cells alive in this generation
type generation struct {
	world               [][]cell
	sizeOfUniverse      int
	numberOfLivingCells int
}

// createBlankGeneration() returns a generation
// The returned generation contains only dead cells to be updated later
func createBlankGeneration(sizeOfUniverse int) generation {
	var blankGeneration generation

	blankGeneration.sizeOfUniverse = sizeOfUniverse

	blankGeneration.world = make([][]cell, sizeOfUniverse)

	// Loops over declaration and initialization of secondary rune slices
	for i := range blankGeneration.world {
		blankGeneration.world[i] = make([]cell, sizeOfUniverse)
	}

	// Loops over creation of all cells in the 2d slice
	for i := 0; i < sizeOfUniverse; i++ {
		for j := 0; j < sizeOfUniverse; j++ {
			blankGeneration.world[i][j].posX = i
			blankGeneration.world[i][j].posY = j
			blankGeneration.world[i][j].status = ' '
			blankGeneration.numberOfLivingCells = 0
		}
	}

	return blankGeneration
}

// initialize() does not return anything
// Fills it's generation according to given seed and sizeOfUniverse
func (g *generation) initialize(sizeOfUniverse int) {

	// Sets size of universe
	g.sizeOfUniverse = sizeOfUniverse

	// Declares and initializes 2 dimensional rune slice
	g.world = make([][]cell, sizeOfUniverse)

	// Loops over declaration and initialization of secondary rune slices
	for i := range g.world {
		g.world[i] = make([]cell, sizeOfUniverse)
	}

	var livingCellCounter int

	// Loops over creation of all cells in the 2d slice
	for i := 0; i < sizeOfUniverse; i++ {
		for j := 0; j < sizeOfUniverse; j++ {
			g.world[i][j].posX = i
			g.world[i][j].posY = j
			// If random number is 1
			if rand.Intn(2) == 1 {
				// Sets cell.status as alive
				g.world[i][j].status = 'O'
				// Counts number of living cells in current generation
				livingCellCounter++
				// If random number is 0
			} else {
				// Sets cell.status as dead
				g.world[i][j].status = ' '
			}
		}
	}
	// Assigns the counted number of living cells
	g.numberOfLivingCells = livingCellCounter
}

// countNeighbours returns an int representing the number of living neighbours of the given cell
// If given cell is at an edge of the square, treats first cell at other end of square as neighbour
func (g *generation) countNeighbours(c cell) int {
	// Counter variable to count alive neighbours
	var counter int
	// Loops over possible offsets from the cell
	for xSetOff := -1; xSetOff < 2; xSetOff++ {
		for ySetOff := -1; ySetOff < 2; ySetOff++ {
			// The x coordinate to be used for the neighbour
			var neighbourPosX int
			// The y coordinate to be used for the neighbour
			var neighbourPosY int

			// Determines if the cell is at the top edge of the square
			if c.posX+xSetOff < 0 {
				// Defines top neighbours as being from bottom edge of square
				neighbourPosX = len(g.world) - 1
				//    Determines if the cell is at the bottom edge of the square
			} else if c.posX+xSetOff == len(g.world) {
				// Defines bottom neighbours as being from top edge of square
				neighbourPosX = 0
			} else {
				// If cell is neither at top nor bottom edge,
				// define neighbours as being from the rows above and below of the cell
				neighbourPosX = c.posX + xSetOff
			}

			// Determines if the cell is at the left edge of the square
			if c.posY+ySetOff < 0 {
				// Defines left neighbours as being from right edge of square
				neighbourPosY = len(g.world) - 1
				// Determines if the cell is at the right edge of the square
			} else if c.posY+ySetOff == len(g.world) {
				// Defines right neighbours as being from left edge of square
				neighbourPosY = 0
			} else {
				// If cell is neither at left nor right edge,
				// define neighbours as being from the rows left and right of the cell
				neighbourPosY = c.posY + ySetOff
			}
			// If the neighbour is alive and NOT the cell itself
			if g.world[neighbourPosX][neighbourPosY].status == 'O' && !(ySetOff == 0 && xSetOff == 0) {
				// Count an additional living neighbour
				counter++
			}
		}
	}
	// Returns counted number of living neighbours
	return counter
}

// createNextGeneration() returns a new generation, based on the one this method is invoked on
func (g *generation) createNextGeneration() generation {
	// Variable for next generation
	nextGeneration := createBlankGeneration(g.sizeOfUniverse)

	var livingCellCounter int

	// Loops over each possible cell in the current universe
	// and fills sets it's state in the nextGeneration
	for i := 0; i < g.sizeOfUniverse; i++ {
		for j := 0; j < g.sizeOfUniverse; j++ {
			// Gets number of living neighbours the current cell had in the previous generation
			numberOfNeighbours := g.countNeighbours(cell{posX: i, posY: j})
			// Checks if current cell was alive in the previous generation
			if g.world[i][j].status == 'O' {
				// If current cell had 2 or 3 living neighbours in the previous generation
				if numberOfNeighbours == 2 || numberOfNeighbours == 3 {
					// Set status of cell as 'alive' in the next generation
					nextGeneration.world[i][j].status = 'O'
					// Counts number of living cells in current generation
					livingCellCounter++
				} else {
					// Sets status of cell in the next generation as 'dead',
					// if cell had less than 2 or more than 3 living neighbours
					// in the previous generation
					nextGeneration.world[i][j].status = ' '
				}
				//     If cell was not alive in the previous generation
			} else {
				// If cell had exactly 3 living neighbours in previous generation
				if numberOfNeighbours == 3 {
					// Set status of cell in next generation to 'alive'
					nextGeneration.world[i][j].status = 'O'
					// Counts number of living cells in current generation
					livingCellCounter++
				} else {
					// If not, set status of cell in next generation to 'dead'
					nextGeneration.world[i][j].status = ' '
				}
			}
		}
	}
	// Assigns the counted number of living cells
	nextGeneration.numberOfLivingCells = livingCellCounter
	// Returns the newly generated next generation
	return nextGeneration
}

func main() {
	// write your code here

	// Length of square
	var sizeOfUniverse int

	// Number of generations to create
	numberOfGenerations := 10

	// Scans input from user
	_, err := fmt.Scanln(&sizeOfUniverse)

	// If there was an error during scanning, log error and exit program
	if err != nil {
		log.Fatal(err)
	}

	// Variable for the current generation
	var currentGeneration generation

	// Initializes first generation
	currentGeneration.initialize(sizeOfUniverse)

	for h := 0; h < numberOfGenerations; h++ {
		// Creates new generation based on the previous generation
		currentGeneration = currentGeneration.createNextGeneration()

		fmt.Printf("Generation #%d\n", h+1)
		fmt.Printf("Alive: %d\n", currentGeneration.numberOfLivingCells)
		// Loops over printing of cells
		for i := 0; i < sizeOfUniverse; i++ {
			for j := 0; j < sizeOfUniverse; j++ {
				// Prints cell
				fmt.Print(string(currentGeneration.world[i][j].status))
			}
			// Prints newline after row is finished
			fmt.Print("\n")
		}

		// Pauses program for 500 ms
		time.Sleep(500 * time.Millisecond)

		// If this wasn't the last generation
		if h < numberOfGenerations-1 {
			// Clears console
			fmt.Print("\033c")
		}
	}
}
