package main

import (
	"log"
	"webapi/csv"
	"webapi/routes"
)

func main() {
	err := csv.LoadDataFromCSV("BoulderTrailHeads.csv")
	if err != nil {
		log.Fatalf("Error loading CSV: %v", err)
	}

	router := routes.SetupRouter()
	router.Run(":8080")
}
