package main

import (
	"fmt"

	"./models"
	"./routes"
)

func main() {
	// database
	models.StartDbConnection()
	// routers
	router := routes.SetupRouters()
	router.Run(":5555")
	fmt.Println("Server running on port 5555")
}
