package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // generate new seed - default is 1
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
	runChance := 30

	//running := true

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Dungeon!")
	fmt.Println("---------------------")
GAME:
	for {

		enemyHealth := rand.Intn(maxEnemyHealth)  // random number with upper bound of maxhealth
		enemy := enemies[rand.Intn(len(enemies))] // random enemy in slice
		fmt.Println("\t# " + enemy + " has appeared! #\n")

		for enemyHealth > 0 { // for is Go's "while loop"

			// few examples of converting integer to string
			fmt.Println(fmt.Sprintf("%s %d", "\tYour HP:", health))           // reminiscent of C
			fmt.Println("\t" + enemy + "'s HP: " + strconv.Itoa(enemyHealth)) // closer to java concat

			fmt.Println("\n\tWhat would you like to do?")
			fmt.Println("\t1. Attack")
			fmt.Println("\t2. Drink health potion")
			fmt.Println("\t3. Run!")

			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\r\n", "", -1)

			if strings.Compare(text, "1") == 0 { // If attack is selected
				damageDealt := rand.Intn(attackDamage)      // get number for damage dealt
				damageTaken := rand.Intn(enemyAttackDamage) // get number for damage received

				enemyHealth -= damageDealt
				health -= damageTaken

				// flavor text to inform user of new values
				fmt.Println("\t> You strike the " + enemy + " for " + strconv.Itoa(damageDealt) + " damage.")
				fmt.Println("\t> You receive " + strconv.Itoa(damageTaken) + " in retaliation!")

				// gameover condition check
				if health < 1 {
					fmt.Println("\t> You have taken too much damage, you are too weak to go on!")
					break
				}
			} else if strings.Compare(text, "2") == 0 {
				if numHealthPots > 0 { // check if user has any health potions to use
					health += healthPotionHealAmount
					numHealthPots-- // decrease number of remaining health potions
					fmt.Println("\t> You drink a health potion, healing yourself for " + strconv.Itoa(healthPotionHealAmount) + " ." +
						"\n\t> Your health is restored to " + strconv.Itoa(health) + " HP." +
						"\n\t> Total health potions remaining: " + strconv.Itoa(numHealthPots) + ".")
				} else {
					fmt.Println("\t> You have no health potions left! Defeat enemies for a chance to get one!")
				}
			} else if strings.Compare(text, "3") == 0 {
				runCheck := rand.Intn(100)
				if runCheck >= runChance {
					fmt.Println("\tYou run away from the " + enemy + "!")
					continue GAME // return to start of loop, but new enemy
				} else {
					damageTaken := rand.Intn(enemyAttackDamage)
					fmt.Println("\tYou failed to flee! " +
						"\n\t> You take " + strconv.Itoa(damageTaken) + " in retaliation!")
					health -= damageTaken

					// gameover condition check
					if health < 1 {
						fmt.Println("\t> You have taken too much damage, you are too weak to go on!")
						break
					}
				}
			} else {
				fmt.Println("\tInvalid command!")
			}
		}

		//game over check
		if health < 1 {
			fmt.Println("\n You limp out of the dungeon, weak from battle.")
			fmt.Println("\n  #########")
			fmt.Println("# Game Over #")
			fmt.Println("  #########")
			break
		}

		fmt.Println("-------------------------------------")
		fmt.Println("\t # " + enemy + " was defeated! # ")
		fmt.Println(fmt.Sprintf("%s %d %s", "\t # You have", health, "health left."))

		//check for a potion drop
		if rand.Intn(100) < healthPotionDropChance {
			numHealthPots++
			fmt.Println("\t # The " + enemy + " dropped a health potion! # ")
			fmt.Println(fmt.Sprintf("%s %d %s", "\t # You now have", numHealthPots, "health potion(s) left. # "))
		}

		//check if player wants to keep going
		fmt.Println("-------------------------------------")
		fmt.Println("\t What would you like to do now?")
		fmt.Println("\t 1. Continue fighting")
		fmt.Println("\t 2. Exit dungeon")

		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)

		//if not a command, give error
		for strings.Compare(text, "1") != 0 && strings.Compare(text, "2") != 0 {
			fmt.Println("\t Invalid Command!")
			fmt.Println("\n\t What would you like to do now?")
			fmt.Println("\t 1. Continue fighting")
			fmt.Println("\t 2. Exit dungeon")
			fmt.Print("-> ")
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\r\n", "", -1)
		}

		//exit game
		if strings.Compare(text, "1") == 0 {
			fmt.Println("\t You continue on your adventure!")
		} else if strings.Compare(text, "2") == 0 {
			fmt.Println("\nYou exit the dungeon, successful from your adventures!")
			fmt.Println("\n  ###################")
			fmt.Println("# THANKS FOR PLAYING! #")
			fmt.Println("  ###################")
			break
		}

	}

}
