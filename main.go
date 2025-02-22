package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//initialise db

	//FeedbackEntry into table

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
