package main

import (
	"fmt"
	"log"
	"time"

	"github.com/daglamier22/my-clients-be/internal/application"
	"github.com/daglamier22/my-clients-be/internal/database"
	"github.com/daglamier22/my-clients-be/internal/utils"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := utils.GetEnvInt("PORT", 8080)
	username := utils.GetEnvString("DB_USERNAME", "admin")
	password := utils.GetEnvString("DB_PASSWORD", "adminpassword")
	host := utils.GetEnvString("DB_HOST", "localhost")
	dbport := utils.GetEnvString("DB_PORT", "5432")
	databaseName := utils.GetEnvString("DB_DATABASE", "my-clients")
	ssl := utils.GetEnvString("DB_SSL", "disable")
	schema := utils.GetEnvString("DB_SCHEMA", "public")
	dbAddr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s", username, password, host, dbport, databaseName, ssl, schema)

	cfg := application.Config{
		Addr: fmt.Sprintf(":%d", port),
		Db: application.DbConfig{
			Addr:         dbAddr,
			MaxOpenConns: utils.GetEnvInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: utils.GetEnvInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  utils.GetEnvDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
	}

	db, err := database.New(cfg.Db.Addr, cfg.Db.MaxOpenConns, cfg.Db.MaxIdleConns, cfg.Db.MaxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Print("database connection pool established")

	app := &application.Application{
		Config: cfg,
		Db:     db,
	}

	mux := app.RegisterRoutes()
	log.Fatal(app.Run(mux))
}
