package main

import (
	"flag"
	"log"

	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config, err := internal.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	var steps = flag.Int("steps", 0, "Number of migration steps (0 = all)")
	flag.Parse()

	db, err := internal.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db.Conn, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create postgres driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/schema",
		"postgres", driver)
	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}

	if *steps == 0 {
		log.Println("Running all migrations...")
		err = m.Up()
	} else {
		log.Printf("Running %d migrations...\n", *steps)
		err = m.Steps(*steps)
	}

	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to run")
		} else {
			log.Fatal("Migration failed:", err)
		}
	} else {
		log.Println("Migrations completed successfully")
	}
}
