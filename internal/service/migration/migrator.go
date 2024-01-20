package migration

import (
	config "github.com/DrNikita/People/internal/db"
)

func MigrateTable(tableStruct interface{}) error {
	dbConn := config.GetDBInstance()
	if err := dbConn.AutoMigrate(tableStruct); err != nil {
		return err
	}
	return nil
}
