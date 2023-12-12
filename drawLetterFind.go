package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawLettersFind(letters string) {
	// box size
	width, height := 100, 5
	x, y := 25, 13

	// vertical border
	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}
	// horizontal border
	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	// angle
	termbox.SetCell(25, 13, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(124, 13, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(25, 17, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(124, 17, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	//title
	TbPrint(27, 15, termbox.ColorDefault, termbox.ColorDefault, letters)
	TbPrint(26, 13, termbox.ColorCyan, termbox.ColorDefault, "Letter find")
}
