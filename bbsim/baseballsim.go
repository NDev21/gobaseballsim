package main

import (
	"fmt"
	"project/baseballsim/player"
)

func main() {
	fmt.Println("Welcome to Baseball Sim in Go!")
	fmt.Printf("Hello\n")
	nick := player.NewPlayer("Nick", 21, 26)
	player.PrintName(nick)
}
