package main

import (
	"github.com/ftsog/goweb/database"
	"github.com/ftsog/goweb/handlers"
)

func NewDataCenter(dbconn database.DataCenterInterface) *database.DataCenter {
	dbmodel := database.DataCenter{
		DBconn: dbconn,
	}

	return &dbmodel
}

func NewApplication() *handlers.Application {
	dbconn := database.DatabaseConnection()
	dc := NewDataCenter(dbconn)
	app := handlers.Application{
		Model: dc,
	}

	return &app

}
