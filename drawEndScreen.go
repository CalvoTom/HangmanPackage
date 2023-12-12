package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawVictory() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	size := 0
	for _, letter := range "victoire" {
		DisplayASCII("standard.txt", 28+size, 15, letter)
		size += 11
	}
	drawEndScreenBox(termbox.ColorGreen)
	termbox.Flush()
}

func DrawLoose() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	size := 0
	for _, letter := range "loose" {
		DisplayASCII("standard.txt", 40+size, 15, letter)
		size += 11
	}
	drawEndScreenBox(termbox.ColorRed)
	termbox.Flush()
}

func drawEndScreenBox(color termbox.Attribute) {
	// box size
	width, height := 140, 40
	x, y := 0, 0

	// Verfical border
	for i := 0; i < width; i++ {
		termbox.SetCell(x+i, y, '#', color, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, '#', color, termbox.ColorDefault)
	}
	// Horizontal border
	for i := 1; i < height; i++ {
		termbox.SetCell(x, y+i, '#', color, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, '#', color, termbox.ColorDefault)
	}
}
