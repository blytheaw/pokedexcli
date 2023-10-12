package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *config, params []string) error {
	if len(params) > 1 {
		return errors.New("Expected format: 'explore <location>'")
	}

	pokemonName := params[0]

	fmt.Printf("\nThrowing a Pokeball at %s\n", pokemonName)
	pokemon, err := config.pokeapiClient.GetPokemonByName(pokemonName)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("base xp: %v\n", pokemon.BaseExperience)
	fmt.Printf("chance: %v\n", chance)

	return nil
}
