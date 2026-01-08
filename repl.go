package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vtusharr/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	Pokedex          map[string]pokeapi.PokeInfo
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		subCommand:=words[1:]


		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg,subCommand)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config,[]string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore":{
			name:    "explore",
			description:"Explore a location",
			callback:  commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List caught Pokémon",
			callback:    commandPokedex,
		},
		"catch":    {
			name:		"Catch",
			description: "Catches the pokemon",
			callback : commandCatch,
		},
	}
}
