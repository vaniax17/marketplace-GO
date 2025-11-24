package main

import (
	"log"
	"marketplace/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()

	err := r.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		panic(err)
	}

	if err = r.Run(":8080"); err != nil {
		panic(err)
	}

}
