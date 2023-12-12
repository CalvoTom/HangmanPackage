
# Hangman

This project is a hangman's game fully realized in go in the terminal. In order to offer a more interesting user experience, a console interface is available.


## 🔖 Rules

The aim of the game is simple: guess all the letters that make up a word within a maximum of 10 attempts. Each time the player guesses a letter, it is displayed. If not, a picture of a hanged man appears... Be careful, if the hanged man is whole, or if the number of attempts exceeds 10, the game is lost...
## 🚀 Deployment
To deploy this project you will need :
- go 1.21.1

Then run,

```bash
 git clone "https://ytrack.learn.ynov.com/git/catom/Hangman.git"
```

To launch the game if you're in the main folder, ,
```bash
 go run main/main.go [arguments]
```
If you're not then run,

```bash
 cd hangman
```
## 🎮 How to play
Le jeu est  modulable vous pouvez choisir votre bibliothèque de mot ou encore votre police de jeu parmis trois disponnible. La progression d'une partie est sauvegardable a tout moment puis peut être reprise plus tard. Voici la liste des arguments a votre disposition pour :

| Flags             | Utility                                                                |
| ----------------- | ------------------------------------------------------------------ |
| [words.txt] | Runs the program with the word library contained in words.txt |
| --letterFile [police.txt] | Run the game with the font in [file.txt], default is not in ASCII |
| --startWith [save.txt] | Restores the game to its last saved state |



## 📚 Authors

- [@catom](https://ytrack.learn.ynov.com/git/catom)


