package hangman

import (
	"bufio"
	"os"

	"github.com/nsf/termbox-go"
)

func DrawHangman(hangman *HangManData) {
	drawHangmanBox()
	err := displayDrawingFromFile("hangman.txt", 8, 8, hangman)
	if err != nil {
		panic(err)
	}
}

func drawHangmanBox() {
	// box size
	width, height := 25, 18
	x, y := 0, 0

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
	// Angle
	termbox.SetCell(0, 0, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(24, 0, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(0, 17, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(24, 17, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	// Box tittle
	TbPrint(1, 0, termbox.ColorCyan, termbox.ColorDefault, "Hangman")
}

func displayDrawingFromFile(filename string, startX, startY int, hangman *HangManData) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	lastStop := 0

	for scanner.Scan() {
		line := scanner.Text()
		drawLine(startX, startY+lineNumber, line)
		lineNumber++
		if lineNumber == hangman.Attempts*7 || hangman.Attempts == 0 {
			break
		}
		if lineNumber == lastStop+7 {
			startY -= 7

			lastStop = lineNumber
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func drawLine(x, y int, line string) {
	for i, ch := range line {
		termbox.SetCell(x+i, y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}
