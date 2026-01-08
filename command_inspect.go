package main

import (
	"fmt"
	"strings"

	"github.com/vtusharr/pokedex/internal/pokeapi"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: inspect <pokemon-name>")
	}

	n := strings.ToLower(strings.TrimSpace(args[0]))
	if n == "" {
		return fmt.Errorf("usage: inspect <pokemon-name>")
	}

	var p pokeapi.PokeInfo
	var ok bool

	// Prefer cached/caught entry if available
	if cfg != nil && cfg.Pokedex != nil {
		p, ok = cfg.Pokedex[n]
	}

	// If not in Pokedex map, fetch from API
	if !ok {
		var err error
		p, err = cfg.pokeapiClient.Getpoke(&n)
		if err != nil {
			return err
		}
	}

	// Print details in the simple format requested
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	fmt.Printf("  -hp: %d\n", getStat(p, "hp"))
	fmt.Printf("  -attack: %d\n", getStat(p, "attack"))
	fmt.Printf("  -defense: %d\n", getStat(p, "defense"))
	fmt.Printf("  -special-attack: %d\n", getStat(p, "special-attack"))
	fmt.Printf("  -special-defense: %d\n", getStat(p, "special-defense"))
	fmt.Printf("  -speed: %d\n", getStat(p, "speed"))
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
