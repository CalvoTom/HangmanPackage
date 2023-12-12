package hangman

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

func DrawAttempts(hangman *HangManData) {
	// box size
	width, height := 25, 9
	x, y := 0, 18

	// Verfical border
	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}
	// Horizontal border
	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	// Angle
	termbox.SetCell(0, 18, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(24, 18, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(0, 26, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(24, 26, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	//box title
	TbPrint(13, 22, termbox.ColorDefault, termbox.ColorDefault, strconv.Itoa(hangman.Attempts))
	TbPrint(1, 18, termbox.ColorCyan, termbox.ColorDefault, "Try")
}
