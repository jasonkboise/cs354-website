package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//game variables
	enemies := [4]string{"Skeleton", "Zombie", "Assassin", "Warrior"}

	maxEnemyHealth := 75
	enemyAttackDamage := 25

	//player variables
	health := 100
	attackDamage := 50
	numHealthPots := 3
	healthPotionHealAmount := 30
	healthPotionDropChance := 50

	running := true

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Dungeon!")
	fmt.Println("---------------------")

	for running == true {

		//TODO: game functionality below
		//I found some example code on how to read player input with a scanner in Golang
		//I put that below.
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}

}
