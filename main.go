package main

import (
	"log"

	"github.com/SayatAbdikul/word_rarity_api/post"
	"github.com/SayatAbdikul/word_rarity_api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := gin.Default()
	router.POST("/add_text", post.AddText)
	router.Run(":8080")
}
