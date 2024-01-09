package main

import (
	"fmt"

	"github.com/samuelmaxi/api-go-rest/models"
	"github.com/samuelmaxi/api-go-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Samuel", History: "History 123..."},
		{Id: 2, Name: "Bob", History: "History 321"},
	}

	fmt.Println("Hello word")
	routes.HandlerRequest()
}
