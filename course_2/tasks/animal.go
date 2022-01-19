package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
    fmt.Println(a.food)
}

func (a Animal) Move() {
    fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
    fmt.Println(a.noise)
}

func main() {
	animalMap := make(map[string]Animal)
	animalMap["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animalMap["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animalMap["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		fmt.Println("Enter the animal name and action divided by the space:")
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)
		
		if input == "x" {
			break
		}

		params := strings.Split(input, " ")

		fmt.Println(params)
		
		animal, animalExists := animalMap[params[0]]
		if !animalExists {
			fmt.Printf("The animal is unknown.")
			continue
		}

		switch params[1] {
			case "eat":
				animal.Move()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Printf("The action is unknown.")
				continue
		}
	}
}
