package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Game struct to represent a game
type Game struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Category string `json:"category" bson:"category"`
	URL      string `json:"url" bson:"url"`
}

// Placeholder in-memory game list (replace with MongoDB later)
var games = []Game{
	{ID: "1", Name: "Super Mario", Category: "Platform", URL: "https://example.com/game1"},
	{ID: "2", Name: "Space Invaders", Category: "Shooter", URL: "https://example.com/game2"},
}

// GetGames returns the list of games
func GetGames(c *gin.Context) {
	c.JSON(http.StatusOK, games)
}

// CreateGame adds a new game to the in-memory list
func CreateGame(c *gin.Context) {
	var newGame Game
	if err := c.ShouldBindJSON(&newGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ideally, save the new game to MongoDB here
	games = append(games, newGame)
	c.JSON(http.StatusCreated, newGame)
}
