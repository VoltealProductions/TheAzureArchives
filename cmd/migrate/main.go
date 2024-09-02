package main

import (
	"fmt"
	"log"
	"os"

	"github.com/VoltealProductions/TheAzureArcchives/config"
	"github.com/VoltealProductions/TheAzureArcchives/db"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := db.NewMySqlStorage(config.Envs.DBDriver, mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	var driver database.Driver
	var src string

	if config.Envs.DBDriver == "sqlite3" {
		driver, err = sqlite3.WithInstance(db, &sqlite3.Config{})
		if err != nil {
			log.Fatal(err)
		}
		src = "file://cmd/migrate/migrations/sqlite"
	} else {
		driver, err = mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			log.Fatal(err)
		}
		src = "file://cmd/migrate/migrations"
	}

	m, err := migrate.NewWithDatabaseInstance(
		src,
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("Database setup: finished.")
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("Database table drop: finished.")
	}
}
