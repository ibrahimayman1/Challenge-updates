package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var Db *bun.DB
var err error

func CreateDatabase() *bun.DB {
	/*config, err := pgx.ParseConfig("postgresql://ibrahim:Q8teKdOtIr3CubzlCsnppA@free-tier13.aws-eu-central-1.cockroachlabs.cloud:26257/gochallenge?sslmode=verify-full&options=--cluster%3Dwoozy-dryad-2598")
	if err != nil {
		panic(err)
	}
	config.PreferSimpleProtocol = true
	sqldb := stdlib.OpenDB(*config)
	Db := bun.NewDB(sqldb, pgdialect.New())
	return Db*/

	//ctx := context.Background()

	// Open a PostgreSQL database.
	dsn := "postgresql://ibrahim:Q8teKdOtIr3CubzlCsnppA@free-tier13.aws-eu-central-1.cockroachlabs.cloud:26257/woozy-dryad-2598.gochallenge?sslmode=verify-full"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return db
}
