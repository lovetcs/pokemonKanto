package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PokemonEntry struct {
	EntryNumber    int `json:"entry_number"`
	PokemonSpecies struct {
		Name string `json:name`
	} `json:"pokemon_species"`
}

type Response struct {
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(responseData))

	var resp Response
	err = json.Unmarshal(responseData, &resp)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range resp.PokemonEntries {
		fmt.Printf("Entry Number: %d\n", entry.EntryNumber)
		fmt.Printf("Pokemon Species: %s\n", entry.PokemonSpecies.Name)
	}
}
