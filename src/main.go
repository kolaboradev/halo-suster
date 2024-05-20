package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	dbpostgres "github.com/kolaboradev/halo-suster/src/databases/postgres"
	httpServer "github.com/kolaboradev/halo-suster/src/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file" + err.Error())
	}
	db, err := dbpostgres.NewDB()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	fmt.Println("Sukses connect")

	serverHttp := httpServer.NewServer(db)

	serverHttp.Listen()
}
