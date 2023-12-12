package hangman

import "github.com/nsf/termbox-go"

func DrawButtonsSave() {
	//bouton size
	width, height := 6, 3
	x, y := 25, 29

	// Vertical border 
	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}
	// horizontal border
	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}

	//Angle
	termbox.SetCell(25, 29, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(30, 29, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(25, 31, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(30, 31, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	// title
	TbPrint(26, 30, termbox.ColorCyan, termbox.ColorDefault, "SAVE")
}

func DrawButtonsQuit() {
	// bouton size
	width, height := 6, 3
	x, y := 31, 29

	//border vertical
	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}
	//border horizontal
	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	//angle
	termbox.SetCell(31, 29, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(36, 29, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(31, 31, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(36, 31, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	// title
	TbPrint(32, 30, termbox.ColorCyan, termbox.ColorDefault, "QUIT")
}
