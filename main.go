package main

import (
	"fmt"
	"log"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func newCliCommand(name, desc string, cb func() error) *cliCommand {
	return &cliCommand{
		name:        name,
		description: desc,
		callback:    cb,
	}
}

var commands = make(map[string]cliCommand)

func commandHelp() error {
	fmt.Println("Usage")
	for _, cmd := range commands {
		fmt.Printf("%s : %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	defer os.Exit(0)
	fmt.Println("Exiting")
	return nil
}

func main() {
	commands["help"] = *newCliCommand("help", "display help", commandHelp)
	commands["exit"] = *newCliCommand("exit", "Exit the repl", commandExit)

	for {
		var input string
		fmt.Print("pokedex > ")
		fmt.Scanln(&input)

		if cmd, ok := commands[input]; ok {
			err := cmd.callback()
			if err != nil {
				log.Fatalf("An Error occured: %v", err)
			}
		} else {
			fmt.Printf("Unknown command")
		}
	}
}
