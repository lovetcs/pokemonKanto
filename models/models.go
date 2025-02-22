package models

type FeedbackEntry struct {
	ID        uint `json:"id"`
	PokemonID uint `json:"pokemin_id"`
	Type      int  `json:"type"`
}
