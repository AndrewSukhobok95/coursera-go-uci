package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)


type Animal interface {
	Eat()
	Move()
    Speak()
}


type Cow struct {}

func (c Cow) Eat() {
    fmt.Println("grass")
}

func (c Cow) Move() {
    fmt.Println("walk")
}

func (c Cow) Speak() {
    fmt.Println("moo")
}


type Bird struct {}

func (b Bird) Eat() {
    fmt.Println("worms")
}

func (b Bird) Move() {
    fmt.Println("fly")
}

func (b Bird) Speak() {
    fmt.Println("peep")
}


type Snake struct {}

func (s Snake) Eat() {
    fmt.Println("mice")
}

func (s Snake) Move() {
    fmt.Println("slither")
}

func (s Snake) Speak() {
    fmt.Println("hsss")
}


func RunAnimalAction(animal Animal, action string) {
	switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("The action is unknown.")
	}
}

func GetAnimalType(animalTypeName string) Animal {
	var animal Animal
	switch animalTypeName {
		case "cow":
			animal = Cow{}
		case "bird":
			animal = Bird{}
		case "snake":
			animal = Snake{}
		default:
			fmt.Println("The animal is unknown.")
	}
	return animal
}

func NewAnimal(animalName string, animalType string, animalMap map[string]Animal) {
	animal := GetAnimalType(animalType)
	if animal != nil {
		animalMap[animalName] = animal
		fmt.Println("Created it!")
	}
}

func Query(animalName string, animalAction string, animalMap map[string]Animal) {
	animal, animalExists := animalMap[animalName]
	if animalExists {
		RunAnimalAction(animal, animalAction)
	} else {
		fmt.Println("The animal with such name is not created.")
	}
}

func Help() {
	fmt.Println("The possible commands:")
	fmt.Println(" newanimal {name} {cow|bird|snake}  create new animal")
	fmt.Println(" query {name} {eat|move|speak}      get info about created animal")
	fmt.Println(" help                               get info about the possible commands")
	fmt.Println(" x                                  stop the program")
}


func main() {
	animalMap := make(map[string]Animal)

	for {
		fmt.Println("\nEnter the command (help command gives options):")
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)

		if input == "x" { break }

		params := strings.Split(input, " ")
		switch params[0] {
			case "newanimal":
				NewAnimal(params[1], params[2], animalMap)
			case "query":
				Query(params[1], params[2], animalMap)
			case "help":
				Help()
			default:
				fmt.Println("The command is unknown.")
		}
	}
}


