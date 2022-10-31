package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main2() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error to loading .env file")
	}

	fmt.Println()
}
