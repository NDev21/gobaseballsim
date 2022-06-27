package main

import (
	"fmt"
	"project/baseballsim/player"
)

func main() {
	fmt.Println("Welcome to Baseball Sim in Go!")
	nick := player.NewPlayer("Nick", 21, 26)

	fmt.Println(player.NameString(nick))
}
