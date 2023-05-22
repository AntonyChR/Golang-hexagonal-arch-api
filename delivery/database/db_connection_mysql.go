package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Database string
	Host     string
	Port     string
	User     string
	Password string
}

func ConnectToDatabase(config DBConfig) (*sql.DB, error) {
	connectionUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", connectionUrl)

	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	return db, nil
}
