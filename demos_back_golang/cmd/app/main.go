package main

import (
	"flag"
	"log"

	"demos_back_golang/internal/config"
	"demos_back_golang/internal/database"
)

func main() {
	configPath := flag.String("config", "./configs/local.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDatabase(cfg.Database.Addres)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	storage := database.NewStorage(db)
	log.Println("database connection pool established")

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
