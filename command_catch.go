package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vtusharr/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, s []string) error {
	if len(s) == 0 {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}

	poke := &s[0]
	Resp, err := cfg.pokeapiClient.Getpoke(poke)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *poke)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	

	// Higher base_experience => harder to catch.
	
	maxExp := 500.0
	baseExp := float64(Resp.BaseExperience)
	successChance := 0.75 - (baseExp / maxExp)
	if successChance < 0.05 {
		successChance = 0.05
	}
	if successChance > 0.95 {
		successChance = 0.95
	}

	roll := rand.Float64()
	if roll < successChance {
		// ensure Pokedex map exists
		if cfg.Pokedex == nil {
			cfg.Pokedex = make(map[string]pokeapi.PokeInfo)
		}

		// add to user's Pokedex (keyed by poke name)
		cfg.Pokedex[Resp.Name] = Resp
		fmt.Printf("%s was caught!\n", Resp.Name)
	} else {
		fmt.Printf("%s escaped!\n", Resp.Name)
	}

	return nil
}
