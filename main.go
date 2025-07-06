package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-railway-app/controller" // Replace with your actual module name
)

func main() {
	_ = godotenv.Load()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://next-vercel-app-ten.vercel.app", "https://next-vercel-app-beta.vercel.app/", "https://next-vercel-app-git-main-surajs-projects-001e53ac.vercel.app/", "https://next-vercel-o9kmtt482-surajs-projects-001e53ac.vercel.app/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	router.GET("/admin/get-posts", controller.GetPosts)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port", port)

	if err := router.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Server failed:", err)
	}
}
