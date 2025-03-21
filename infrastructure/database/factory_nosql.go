package database

import (
	"errors"

	"github.com/gsabadini/go-clean-architecture/adapter/repository"
)

var (
	errInvalidNoSQLDatabaseInstance = errors.New("invalid nosql db instance")
)

const (
	InstanceDynamoDB int = iota
)

func NewDatabaseNoSQLFactory(instance int) (repository.NoSQL, error) {
	switch instance {
	case InstanceDynamoDB:
		return NewDynamoDBHandler(newConfigDynamoDB()), nil
	default:
		return nil, errInvalidNoSQLDatabaseInstance
	}
}
