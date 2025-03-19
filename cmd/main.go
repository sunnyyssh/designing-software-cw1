package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sunnyyssh/designing-software-cw1/internal/cli"
	"github.com/sunnyyssh/designing-software-cw1/internal/config"
)

const EnvPGConnString = "PG_CONN_STRING"

func main() {
	ctx := context.Background()

	connString := os.Getenv(EnvPGConnString)
	if connString == "" {
		log.Fatalf("%s env variable is not set", EnvPGConnString)
	}

	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("Connecting DB failed: %s", err)
	}

	dbConf := config.NewDB(db)
	svcConf := config.NewServices(dbConf)

	cli.CLI(svcConf).Execute()
}
