package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/offluck/ilove2rest/internal/config"
	"github.com/offluck/ilove2rest/internal/repository"
	"github.com/offluck/ilove2rest/internal/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	pgDriver       = "postgres"
	redisDriver    = "redis"
	migrationsPath = "file://migrations/"
)

func initLogger(conf config.Config) (*zap.Logger, error) {
	zapConf := zap.NewDevelopmentConfig()
	level, err := zapcore.ParseLevel(conf.LoggingLevel)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse level: %+v\n", err)
	}
	log.Printf("Setting logging level: %+v\n", level)

	zapConf.Level.SetLevel(level)
	logger, err := zapConf.Build()
	if err != nil {
		return nil, fmt.Errorf("Failed to build logger: %+v\n", err)
	}

	log.Printf("Logger has been set up: %+v\n", logger)
	return logger, nil
}

func initDB(dbURL string) (*sql.DB, error) {
	dataBase, err := sql.Open(pgDriver, dbURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to open DB: %+v\n", err)
	}

	err = dataBase.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping DB: %+v\n", err)
	}

	return dataBase, nil
}

func migrateDB(database *sql.DB) error {
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Failed to create postgres driver: %+v\n", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("Failed to create migration instance: %+v\n", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Printf("Migration did not change DB")
			return nil
		}
		return fmt.Errorf("Failed to migrate: %+v\n", err)
	}

	return nil
}

func main() {
	configPath := ""
	flag.StringVar(&configPath, "config", "", "Configuration file path")
	flag.Parse()

	if configPath == "" {
		log.Fatal("Please, provide config path via -config flag")
	}
	log.Printf("Config path: %s\n", configPath)

	conf, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to read the config: %+v\n", err)
	}
	log.Printf("Config: %+v\n", conf)

	logger, err := initLogger(conf)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %+v", err)
	}

	dataBase, err := initDB(conf.DB.GetDBURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}

	err = migrateDB(dataBase)
	if err != nil {
		log.Fatalf("Failed to migrate database: %+v", err)
	}

	logger.Info("Starting server")
	s := server.New(conf.Port, repository.NewPGClient(dataBase, logger), logger)
	s.Start()
}
