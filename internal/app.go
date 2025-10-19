package internal

import (
	fmt "fmt"
	log "log"
	http "net/http"

	// init config
	config "go-template/internal/config"
	// init db
	db "go-template/internal/db"
	// init router
	myhttp "go-template/internal/http"
)

func Run() error {
	cfg := config.Load()

	database, err := db.Connect(cfg.DBUrl)
	if err != nil {
		return fmt.Errorf("db connection error: %w", err)
	}
	// close database connection after Run() func ended
	defer database.Close()

	router := myhttp.NewRouter(database)

	log.Printf("🚀 Server running on port %s", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, router)
	// if everything ok Run() return nil
}
