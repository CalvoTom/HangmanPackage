package hangman

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/nsf/termbox-go"
)

// Print word using termbox
func TbPrint(x, y int, fg, bg termbox.Attribute, text string) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

// Replace _ in hangman.word if letter is find
func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

// Try if letter is between A-Z, a-z and à-û
func IsAlpha(str string) bool {
	var alphanumeric = regexp.MustCompile("^[[A-zÀ-û]*$")
	return alphanumeric.MatchString(str)
}

func ReplaceAtIndex2(input string, replacement byte, index int) string {
	return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}

// Initilise hangman structure with default values
func InitialiseStruc(filename string) *HangManData {
	var hangman *HangManData
	hangman = new(HangManData)

	RandomWord(hangman, filename)
	arrayTf := []rune(hangman.ToFind)
	for i := 0; i < len(arrayTf); i++ {
		hangman.Word += "_"
	}
	hangman.LetterFind = HideWord(hangman)
	hangman.IsASCII = false
	hangman.Attempts = 10

	return hangman
}

// Chose random word in a file passed in parameter
func RandomWord(hangman *HangManData, filename string) {
	var arrayWord []string

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		arrayWord = append(arrayWord, scanner.Text())
	}
	randomWordIndex := rand.Intn(lineNumber - 1)
	hangman.ToFind = arrayWord[randomWordIndex]
}

// Randomly hiding a word by replacing letter with _
func HideWord(hangman *HangManData) string {
	var lettersFind string
	arrayWord := []rune(hangman.Word)
	arrayTofind := []rune(hangman.ToFind)
	nbLetterReveal := len(arrayTofind)/2 - 1

	for i := 0; i <= nbLetterReveal; i++ {
		randomIndex := rand.Intn(len(arrayWord) - 1)
		arrayWord[randomIndex] = arrayTofind[randomIndex]

		if !strings.Contains(lettersFind, string(arrayTofind[randomIndex])) {
			lettersFind += string(arrayTofind[randomIndex])
		}
		for index, ch := range arrayTofind {
			if ch == arrayTofind[randomIndex] {
				arrayWord[index] = arrayTofind[randomIndex]
			}
		}
	}
	hangman.Word = string(arrayWord)
	return lettersFind
}

// Display at screen a word letter by letter in a ASCII file passed in parameter
func DisplayASCII(filename string, startX, startY int, letter rune) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	liner := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		if letter == '_' {
			if liner >= 9 {
				break
			}
			if lineNumber >= 568 {
				drawLine(startX, startY+liner, line)
				liner++
			}

		} else if lineNumber >= 586 {
			if liner >= 9 {
				break
			}
			if lineNumber >= ((int(letter)-97)*9)+586 {
				drawLine(startX, startY+liner, line)
				liner++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func DrawLine(x, y int, line string) {
	for i, ch := range line {
		termbox.SetCell(x+i, y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}

// Save the current state of the hangman structure in save.txt file
func Save(hangman *HangManData) {
	save, err := json.Marshal(hangman)
	if err != nil {
		os.Exit(1)
	}

	file, err := os.Create("save.txt")
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()

	_, err = file.Write(save)
	if err != nil {
		os.Exit(3)
	}
}

// Update hangman structure in state of save.txt
func LoadGame(file string, hangman *HangManData) {
	load, err := ioutil.ReadFile(file)
	if err != nil {
		os.Exit(4)
	}
	err = json.Unmarshal(load, hangman)
	if err != nil {
		os.Exit(5)
	}
}

// Check if mouss is inside a bouton
func isMouseInsideButton(mouseX, mouseY, buttonX, buttonY int, label string) bool {
	return mouseX >= buttonX && mouseX < buttonX+len(label) && mouseY == buttonY
}

func Testeur(input string, hangman *HangManData) bool {
	switch {
	case len(input) <= 1:
		if strings.Contains(string(hangman.ToFind), input) {
			for i, letter := range hangman.ToFind {
				if string(letter) == input {
					hangman.Word = ReplaceAtIndex(hangman.Word, letter, i)
					if !strings.Contains(string(hangman.LettersTried), input) {
						hangman.LetterFind += input
						hangman.LettersTried += input
					}
				}
			}
			return true
		} else {
			if !strings.Contains(string(hangman.LettersTried), input) {
				hangman.Attempts -= 1
				hangman.LettersTried += input
			}
			return false
		}

	case len(input) == len(hangman.ToFind):
		returnValue := true
		for _, ch := range input {
			if returnValue == false {
				return false
			}
			if strings.Contains(string(hangman.ToFind), string(ch)) {
				for i, letter := range hangman.ToFind {
					if string(letter) == string(ch) {
						hangman.Word = ReplaceAtIndex(hangman.Word, letter, i)
						if !strings.Contains(string(hangman.LettersTried), string(ch)) {
							hangman.LetterFind += string(ch)
							hangman.LettersTried += string(ch)
						}
					}
				}
			} else {
				if !strings.Contains(string(hangman.LettersTried), string(ch)) {
					hangman.Attempts -= 2
					hangman.LettersTried += string(ch)
					returnValue = false
				}
			}
		}

	default:
		return false
	}
	return false
}
