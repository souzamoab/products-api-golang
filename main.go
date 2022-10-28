package main

import (
	"github.com/joho/godotenv"
	"github.com/souzamoab/products-api-golang/src/app/database"
	"github.com/souzamoab/products-api-golang/src/app/server/routes"
)

func main() {
	database.StartDatabase()
	godotenv.Load()
	router := routes.ConfigRoutes()
	router.Run(":8080")
}
