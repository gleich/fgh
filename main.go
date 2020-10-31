package main

import (
	"github.com/Matt-Gleich/fgh/pkg/commands"
	"github.com/Matt-Gleich/statuser/v2"
)

func main() {
	statuser.Emojis = false
	commands.Execute()
}
