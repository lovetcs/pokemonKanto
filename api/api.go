package api

import (
	"fmt"
	"log"
	// Import the database package
	// Import the models package
)

func fetchAndDisplayFeedback(pokemonName string) {
	var feedbacks []FeedbackEntry

	//Fetch feedback for a specific Pokemon
	if err := db.Where("pokemon_id = ?", pokemonName).Find(&feedbacks).Error; err != nil {
		log.Fatal(err)
	}

	//Display Feedback
	if len(feedbacks) > 0 {
		for _, feedback := range feedbacks {
			fmt.Printf("Feedback Type: %d\n", feedback.Type)
		}
	} else {
		fmt.Println("No feedback available for", pokemonName)
	}
}

func addFeedback(pokemonID uint, feedbackType int) {
	feedback := FeedbackEntry{
		PokemonID: pokemonID,
		Type:      feedbackType,
	}
	if err := db.Create(&feedback).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Feedback added for Pokémon ID: %d\n", pokemonID)
}
