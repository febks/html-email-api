package main

import (
	"go-send-email/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: Error loading .env file")
		}
	} else {
		log.Println("Warning: .env file not found. Proceeding without it.")
	}
}

func main() {
	r := router.SetupRouter()
	log.Println("Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
