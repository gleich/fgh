package main

import (
	"github.com/gleich/fgh/pkg/commands"
	"github.com/gleich/statuser/v2"
)

func main() {
	statuser.Emojis = false
	commands.Execute()
}
