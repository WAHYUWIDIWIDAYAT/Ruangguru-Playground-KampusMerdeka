package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name    string `json:"name"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		}
	}
}

func GetPokemonData() (*Pokemon, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/1/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func main() {
	pokemon, err := GetPokemonData()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pokemon)
}
