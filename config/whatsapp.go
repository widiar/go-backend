package config

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var WA *whatsmeow.Client

func SetupWa() {
	sqlDb, err := DB.DB()
	if err != nil {
		panic("Failed to extract db for WA")
	}

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container := sqlstore.NewWithDB(sqlDb, "postgres", dbLog)
	errWa := container.Upgrade(context.Background())
	if errWa != nil {
		panic("Failed setup load db for wa")
	}

	deviceStore, err := container.GetFirstDevice(context.Background())
	if err != nil {
		panic("Failed to load devices")
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	waClient := whatsmeow.NewClient(deviceStore, clientLog)

	if waClient.Store.ID != nil {
		errs := waClient.Connect()
		if errs != nil {
			fmt.Println("Failed auto reconnect", errs)
		} else {
			fmt.Println("Success to reconnect")
		}
	}
	WA = waClient
}
