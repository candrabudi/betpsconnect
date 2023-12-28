package main

import (
	"betpsconnect/database"
	"betpsconnect/database/migration"
	"betpsconnect/database/seeder"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/http"
	"betpsconnect/pkg/util"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var m string
	var s string

	database.CreateMongoConnection() // Menggunakan koneksi ke MongoDB

	flag.StringVar(
		&m,
		"m",
		"none",
		`This flag is used for migration`,
	)

	flag.StringVar(
		&s,
		"s",
		"none",
		`This flag is used for seeder`,
	)

	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
	}

	if s == "seeder" {
		seeder.Seed()
	}

	f := factory.NewFactory() // Instance database initialization
	g := gin.New()

	http.NewHttp(g, f)

	if err := g.Run(":" + util.GetEnv("APP_PORT", "8080")); err != nil {
		log.Fatal("Can't start server.")
	}
}
