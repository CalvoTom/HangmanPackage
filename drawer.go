package hangman

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	borderTopLeft     = 0x250C
	borderTopRight    = 0x2510
	borderBotomLeft   = 0x2514
	borderBottomRight = 0x2518
	borderHorizontal  = 0x2502
	borderVertical    = 0x2500
)

type HangManData struct {
	Word         string // Word composed of '_', ex: H_ll_
	ToFind       string // Final word chosen by the program at the beginning. It is the word to find
	Attempts     int    // Number of attempts left
	IsASCII      bool
	Police       string
	LetterFind   string
	LettersTried string
}

func Drawer() {
	// Initialise variable
	var file string
	var lettersTried string
	var hangman *HangManData

	// Check flags
	flag.String("startWith", "default", "File name to start with")
	flag.String("letterFile", "default", "File name to choose ASCII")
	flag.Parse()
	if os.Args[1] == "--startWith" && os.Args[2] == "save.txt" {
		hangman = new(HangManData)
		LoadGame("save.txt", hangman)
	} else if len(os.Args[1:]) >= 3 {
		file = os.Args[1]
		if os.Args[2] == "--letterFile" && (os.Args[3] == "standard.txt" || os.Args[3] == "shadow.txt" || os.Args[3] == "thinkertoy.txt") {
			hangman = InitialiseStruc(file)
			hangman.IsASCII = true
			hangman.Police = os.Args[3]
		} else if os.Args[3] == "default.txt" {
			hangman = InitialiseStruc(file)
		}
	} else if len(os.Args[1:]) == 1 {
		file = os.Args[1]
		hangman = InitialiseStruc(file)
	} else {
		fmt.Println("Syntax Error")
		os.Exit(1)
	}

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputMouse)

	lettersTried = hangman.LetterFind

	// Game loop
loop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		DrawWordToGuess(hangman)
		DrawAttempts(hangman)
		DrawLettersTried(lettersTried)
		DrawLettersFind(hangman.LetterFind)
		DrawHangman(hangman)
		DrawButtonsSave()
		DrawButtonsQuit()
		termbox.Flush()

		if hangman.Word == hangman.ToFind {
			DrawVictory()
			time.Sleep(3 * time.Second)
			break loop
		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || hangman.Attempts >= 10 {
				DrawLoose()
				time.Sleep(3 * time.Second)
				break loop
			} else if hangman.Word == hangman.ToFind {
				break loop
			} else if IsAlpha(string(ev.Ch)) {
				guess := string(ev.Ch)
				if strings.Contains(string(hangman.ToFind), guess) {
					for i, letter := range hangman.ToFind {
						if string(letter) == guess {
							hangman.Word = ReplaceAtIndex(hangman.Word, letter, i)
							if !strings.Contains(string(lettersTried), guess) {
								hangman.LetterFind += guess
								lettersTried += guess
							}
						}
					}
				} else {
					if !strings.Contains(string(lettersTried), guess) {
						hangman.Attempts += 1
						lettersTried += guess
					}
				}
			}
		case termbox.EventMouse:
			if isMouseInsideButton(ev.MouseX, ev.MouseY, 32, 30, "QUIT") {
				break loop
			}
			if isMouseInsideButton(ev.MouseX, ev.MouseY, 26, 30, "SAVE") {
				Save(hangman)
				break loop
			}
		}
	}
}
