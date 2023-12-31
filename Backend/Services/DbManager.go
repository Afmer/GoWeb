package Services

import (
	"database/sql"

	"GoWeb/Architecture/Interfaces"

	_ "github.com/go-sql-driver/mysql"
)

type dbManager struct {
	settings DbManagerSettings
}

type DbManagerSettings struct {
	ConnectionString string
	DatabaseName     string
}

func CreateDbManager(settings DbManagerSettings) dbManager {
	result := dbManager{
		settings: settings,
	}
	return result
}

func (dbManager dbManager) InitDatabase() {
	db, err := sql.Open("mysql", dbManager.settings.ConnectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SHOW DATABASES LIKE '" + dbManager.settings.DatabaseName + "'")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var databaseName string
	isDbExists := false
	for rows.Next() {
		err := rows.Scan(&databaseName)
		if err != nil {
			panic(err.Error())
		}
		if databaseName == "your_database_name" {
			isDbExists = true
			break
		}
	}
	if !isDbExists {
		_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbManager.settings.DatabaseName)
		if err != nil {
			panic(err.Error())
		}
		_, err = db.Exec("USE " + dbManager.settings.DatabaseName)
		query := `
		CREATE TABLE IF NOT EXISTS users (
			username VARCHAR(50) NOT NULL PRIMARY KEY,
			password_hash VARCHAR(255) NOT NULL
		)
		`
		_, err = db.Exec(query)
		if err != nil {
			panic(err.Error())
		}
	}
}

var _ Interfaces.IDbInit = dbManager{}
