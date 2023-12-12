package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawLettersTried(letters string) {
	// Box size
	width, height := 100, 9
	x, y := 25, 18

	// Vertical border
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
	termbox.SetCell(25, 18, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(124, 18, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(25, 26, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(124, 26, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	// Title
	TbPrint(27, 22, termbox.ColorDefault, termbox.ColorDefault, letters)
	TbPrint(26, 18, termbox.ColorCyan, termbox.ColorDefault, "Letter try")
}
