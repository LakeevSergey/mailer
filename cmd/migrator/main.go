package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migratemysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"

	"github.com/LakeevSergey/mailer/internal/config"
	"github.com/LakeevSergey/mailer/internal/infrastructure/logger"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		fmt.Printf("Parse config error: %v", err)
		return
	}

	zlogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.Level(cfg.ConsoleLoggerLevel)).With().Timestamp().Logger()
	logger := logger.NewLogger(zlogger)

	dbconfig := mysql.NewConfig()
	dbconfig.Net = "tcp"
	dbconfig.Addr = cfg.DBHost + ":" + strconv.Itoa(cfg.DBPort)
	dbconfig.User = cfg.DBUser
	dbconfig.Passwd = cfg.DBPassword
	dbconfig.DBName = cfg.DBName

	db, err := sql.Open("mysql", dbconfig.FormatDSN())
	if err != nil {
		logger.ErrorErr(fmt.Errorf("open DB connection error: %w", err))
		return
	}

	driver, err := migratemysql.WithInstance(db, &migratemysql.Config{})
	if err != nil {
		logger.ErrorErr(fmt.Errorf("get DB driver error: %w", err))
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///usr/src/app/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		logger.ErrorErr(fmt.Errorf("create migrator error: %w", err))
		return
	}

	m.Up()
	if err != nil {
		logger.ErrorErr(fmt.Errorf("migration error: %w", err))
		return
	}
	logger.Info("Migrated")
}
