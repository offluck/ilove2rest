package main

import (
	"flag"
	"log"

	"github.com/offluck/ilove2rest/internal/config"
	"github.com/offluck/ilove2rest/internal/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

	zapConf := zap.NewDevelopmentConfig()
	level, err := zapcore.ParseLevel(conf.LoggingLevel)
	if err != nil {
		log.Fatalf("Failed to parse level: %+v\n", err)
	}
	log.Printf("Setting logging level: %+v\n", level)

	zapConf.Level.SetLevel(level)
	logger, err := zapConf.Build()
	if err != nil {
		log.Fatalf("Failed to build logger: %+v\n", err)
	}
	log.Printf("Logger has been set up: %+v\n", logger)

	logger.Info("Starting server")
	s := server.New(conf.Port, nil, logger)
	s.Start()
}
