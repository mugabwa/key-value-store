package main

import (
	"fmt"

	"github.com/mugabwa/little-key-value/internal/api"
)


func main() {
	fmt.Println("Starting Little Key Value Server...")
	server := api.New()
	server.Server(":9000")
}
