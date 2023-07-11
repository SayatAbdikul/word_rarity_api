package main

import (
	"fmt"

	"github.com/SayatAbdikul/word_rarity_api/post"
	"github.com/SayatAbdikul/word_rarity_api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	server.Connect()
	defer server.DB.Close()
	fmt.Println(1)
	router := gin.Default()
	router.POST("/add_text", post.AddText)
	router.Run(":8080")
}
