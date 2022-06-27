package player

import "fmt"

type Player struct {
	name   string
	number int
	age    int
}

func NewPlayer(name string, number int, age int) Player {
	return Player{name: name, number: number, age: age}
}

func AgeString(player Player) string {
	return fmt.Sprintf("Player Age: %d", player.age)
}

func NameString(player Player) string {
	return fmt.Sprintf("Player Name: %s", player.name)
}

func PrintName(player Player) {
	fmt.Printf("Player Name: %s\n", player.name)
}
