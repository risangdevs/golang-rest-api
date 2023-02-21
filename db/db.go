package db

import (
	"database/sql"
	"fmt"

	packr "github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var DbConnection *sql.DB

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migrations"),
	}
	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}
	DbConnection = dbParam
	fmt.Println("Applied", n, " migrations!")
}
