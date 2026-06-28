package config

import (
	"context"
	"database/sql"

	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func SetupWa(sqlDb *sql.DB) *sqlstore.Container {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container := sqlstore.NewWithDB(sqlDb, "postgres", dbLog)
	errWa := container.Upgrade(context.Background())
	if errWa != nil {
		panic("Failed setup load db for wa")
	}
	return container
}
