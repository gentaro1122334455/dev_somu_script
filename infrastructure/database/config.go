package database

import (
	"os"
	"time"
)

type postgresConfig struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

type dynamodbConfig struct {
	region    string
	accessKey string
	secretKey string
}

func newConfigDynamoDB() *dynamodbConfig {
	return &dynamodbConfig{
		region:    os.Getenv("AWS_REGION"),
		accessKey: os.Getenv("AWS_ACCESS_KEY"),
		secretKey: os.Getenv("AWS_SECRET_KEY"),
	}
}

func newConfigPostgres() *postgresConfig {
	return &postgresConfig{
		host:     os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DATABASE"),
		port:     os.Getenv("POSTGRES_PORT"),
		driver:   os.Getenv("POSTGRES_DRIVER"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
