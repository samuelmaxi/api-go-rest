package main

import (
	"fmt"

	"github.com/samuelmaxi/api-go-rest/database"
	"github.com/samuelmaxi/api-go-rest/routes"
)

func main() {
	database.ConnectDB()
	fmt.Println("Hello word")
	routes.HandlerRequest()
}
