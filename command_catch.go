package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *config, params []string) error {
	if len(params) != 1 {
		return errors.New("Expected format: 'catch <pokemon>'")
	}

	pokemonName := params[0]

	fmt.Printf("\nThrowing a Pokeball at %s\n", pokemonName)
	pokemon, err := config.pokeapiClient.GetPokemonByName(pokemonName)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	if chance < 40 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
