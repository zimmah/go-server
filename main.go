package main

import (
	"github.com/joho/godotenv"
	"github.com/zimmah/go-server/internal/router"
)

func main() {
	godotenv.Load()
	router.Router()
}
