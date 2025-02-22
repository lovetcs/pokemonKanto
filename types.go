package main

type PokemonSpecies struct {
	Name string `json:"name"`
}

type PokemonEntry struct {
	EntryNumber    int            `json:"entry_number"`
	PokemonSpecies PokemonSpecies `json:"pokemon_species"`
}

type Response struct {
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
}
