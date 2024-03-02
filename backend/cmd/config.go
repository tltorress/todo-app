package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	ErrInvalidEnvParams = errors.New("there are missing env variables")
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")
	cfg.DB.Pass = os.Getenv("DB_PASSWORD")
	cfg.DB.Database = os.Getenv("DB_DATABASE")

	if ok := cfg.validate(); !ok {
		return nil, ErrInvalidEnvParams
	}

	return cfg, nil
}

type Config struct {
	DB DB
}

type DB struct {
	Host     string
	Port     string
	Pass     string
	Database string
}

func (c Config) validate() bool {
	if c.DB.Host != "" || c.DB.Port != "" || c.DB.Database != "" {
		return true
	}
	return false
}
