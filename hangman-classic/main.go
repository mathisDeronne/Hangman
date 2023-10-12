package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func openfile(f string, sep string) []string {
	file, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(string(file), sep)
	return str
}
func Affiche(index int) {
	hangman := openfile("hangman.txt", "\n")
	min := 0
	max := 8
	for i := 0; i < index; i++ {
		min += 8
		max += 8
	}
	for i := min; i < max; i++ {
		fmt.Println(hangman[i])
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func CheckLetters(tabofanswer, tabofwordselected []string, answer string) (bool, []int) {
	res := []int{}
	for index := range tabofwordselected {
		if answer == tabofwordselected[index] {
			res = append(res, index)
		}
	}
	if len(res) > 0 {
		return true, res
	}
	return false, res
}

func GetHiddenword(word string) []string {
	res := []string{}
	indexes := []int{}
	nb := len(word)/2 - 1
	for i := 0; i < nb; i++ {
		randindex := rand.Intn(len(word))
		for ArrayContains(indexes, randindex) {
			randindex = rand.Intn(len(word))
		}
		indexes = append(indexes, randindex)
	}
	for i, v := range word {
		if ArrayContains(indexes, i) {
			res = append(res, string(v))
		} else {
			res = append(res, "_")
		}
	}
	return res
}

func ArrayContains(tab []int, v int) bool {
	for _, c := range tab {
		if c == v {
			return true
		}
	}
	return false
}

func main() {
	lifes := 10
	args := os.Args[1]
	words, err := os.ReadFile(args)                    // on déclare un tableau de byte "words" qui lis et enregistre les mots contenus dans "word.txt"
	check(err)                                         // on affiche l'erreur "err" si "words.txt" est vide
	stringtosplit := string(words)                     // on fait une string de "words" (tableau de bytes) nommée "stringtosplit"
	tabofstrings := strings.Split(stringtosplit, "\n") // on sépare la string "stringtosplit" en tableau de string "tabofstrings" a chaque retoour a la ligne
	rand.Seed(time.Now().UnixNano())                   // on utilise la fonction rand pour choisir un mot au hasard dans le tableau de string "tabofstrings"
	guesswordplace := rand.Intn(len(tabofstrings))     // on créé une variable guessword qui sera un mot au hasard dans le fichier "word.txt" (ou plutot dans le tableau de string "tabofstrings")
	//fmt.Println(tabofstrings[guesswordplace])          // on affiche le mot qui a été sélectionné (cette ligne est uniquement pour tester le programme car le mot ne doit pas être affiché au joueur)
	wordselected := tabofstrings[guesswordplace]
	var answer string
	allanswers := ""
	tabofwordselected := strings.Split(wordselected, "")
	tabOfHiddenWord := GetHiddenword(wordselected)
	fmt.Println("Votre ami José a été condamné à mort par pendaison, mais c'est votre unique fournisseur de saucisson donc vous voulez le sauver. Pour cela, vous devez retrouver le mot qui pourra l'inocenter. Vous devez donc compléter le mot caché avec les lettres manquantes.")
	for true { 											// on créé une string answer qui sera demandée au joueur (c'est la que le joueur pourra essayer de deviner une lettre ou le mot en entier)
		//fmt.Println(strings.Split(wordselected, ""))
		fmt.Println(tabOfHiddenWord)
		fmt.Println("essayez de trouver une lettre ou le mot en entier ! \nEntrez exitgame si vous voulez quitter le jeu ") // on affiche la phrase "essayer de trouver une lettre ou le mot en entier !" au joueur
		fmt.Scanln(&answer)                                                                                                 // la string "answer" deviens la string que le joueur a entré dans la commande
		tabofanswer := strings.Split(answer, "")
		//fmt.Println(answer)																								// on affiche le mot ou la lettre que le joueur a choisi (cette ligne sert seulement a tester le programme)
		allanswers = allanswers + " " + answer
		if answer == "exitgame" {
			fmt.Println("Fin de partie, Merci d'avoir joué !")
			break
		}
		if answer == "showword" {
			fmt.Println("Hé tu n'était pas censé faire ça !\nBon d'accord voici le mot :", tabofstrings[guesswordplace])
		}
		if answer == "godmode" {
			lifes = 99999999999
			fmt.Println("godmode activé, vous avez ", lifes, "vies")
		}
		//fmt.Println(tabofwordselected)
		if answer != "showword" {
			if len(tabofanswer) == len(tabofwordselected) {
				if answer == wordselected {
					fmt.Println("Bravo, vous avez gagné !\nJosé est sauvé et le saucisson aussi !")
					return
				} else {
					if answer != wordselected {
						fmt.Println("Dommage, ce n'est pas le bon mot !")
						lifes--
					}
				}
			}
			if len(tabofanswer) > 1 && len(tabofanswer) != len(tabofwordselected) {
				fmt.Println("vous ne pouvez essayer de deviner qu'une lettre a la fois ou tout le mot d'un seul coup")
			}
			if cond, index := CheckLetters(tabofanswer, tabofwordselected, answer); cond {
				fmt.Println("Bravo, cette lettre est dans le mot !")
				for _, v := range index {
					tabOfHiddenWord[v] = answer
				}
			} else {
				lifes--
				fmt.Println("Dommage, cette lettre n'est pas dans le mot !")
			}
			Affiche(9 - lifes)
			if lifes == 0 {
				fmt.Println("Vous n'avez plus de vies !\nGame Over.")
				break
			}
			hiddenwordtoverify := strings.Join(tabOfHiddenWord, "")
			//fmt.Println(hiddenwordtoverify)
			if hiddenwordtoverify == wordselected {
				fmt.Println("Vous avez trouvés toutes les lettres du mot !\nBravo, vous avez gagné !\nJosé est sauvé et le saucisson aussi !")
				break
			}
			fmt.Println("Il vous reste", lifes, "vies.")
		}
		fmt.Println("Vous avez joué les lettres suivantes :", allanswers)
	}
}
