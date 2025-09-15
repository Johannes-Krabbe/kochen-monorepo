package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal"
)

// @title kochen.app server
// @version 1.0
// @description this is the server for kochen.app
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @servers http://localhost:8080 Development server
// @servers https://localhost:8080 Development server (HTTPS)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config, err := internal.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	db, err := internal.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	router := internal.NewRouter(db)

	addr := fmt.Sprintf(":%d", config.Port)
	log.Printf("Server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
