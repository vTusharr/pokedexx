package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/vtusharr/pokedex/internal/pokeapi"
)

// commandPokedex prints all caught Pokémon from cfg.Pokedex with their stats and types.

func commandPokedex(cfg *config, _ []string) error {
	if cfg == nil || cfg.Pokedex == nil || len(cfg.Pokedex) == 0 {
		fmt.Println("Pokedex is empty")
		return nil
	}

	// deterministic ordering
	names := make([]string, 0, len(cfg.Pokedex))
	for k := range cfg.Pokedex {
		names = append(names, k)
	}
	sort.Strings(names)

	for i, key := range names {
		p := cfg.Pokedex[key]

		// Capitalize name for nicer output (pokeapi returns lowercase names)
		displayName := p.Name
		if displayName == "" {
			// fall back to map key if PokeInfo.Name missing
			displayName = key
		}
		if len(displayName) > 0 {
			displayName = strings.ToUpper(displayName[:1]) + displayName[1:]
		}

		fmt.Println(displayName)
		fmt.Println("Stats:")

		// Print stats in canonical order
		fmt.Printf("  -hp: %d\n", getStat(p, "hp"))
		fmt.Printf("  -attack: %d\n", getStat(p, "attack"))
		fmt.Printf("  -defense: %d\n", getStat(p, "defense"))
		fmt.Printf("  -special-attack: %d\n", getStat(p, "special-attack"))
		fmt.Printf("  -special-defense: %d\n", getStat(p, "special-defense"))
		fmt.Printf("  -speed: %d\n", getStat(p, "speed"))

		fmt.Println("Types:")
		if len(p.Types) == 0 {
			fmt.Println("  -")
		} else {
			// preserve API slot order (sorted by Slot)
			types := make([]string, 0, len(p.Types))
			for _, t := range p.Types {
				types = append(types, t.Type.Name)
			}
			for _, tn := range types {
				fmt.Printf("  - %s\n", tn)
			}
		}


		if i != len(names)-1 {
			fmt.Println()
		}
	}

	return nil
}

func getStat(p pokeapi.PokeInfo, name string) int {
	for _, s := range p.Stats {
		if s.Stat.Name == name {
			return s.BaseStat
		}
	}
	return 0
}
