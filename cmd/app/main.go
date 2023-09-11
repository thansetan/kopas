package main

import (
	"github.com/thansetan/kopas/internal/infrastructure/database"
	httproute "github.com/thansetan/kopas/internal/infrastructure/http"
	"github.com/thansetan/kopas/pkg/helpers"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	dbPath := helpers.GetEnvOrDefault("BADGER_PATH", "badger")
	db, err := database.NewBadger(dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := httproute.NewRoute(db)
	r.Run()
}
