package repository

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type NoSQL interface {
	PutItem(context.Context, *dynamodb.PutItemInput) error
	FindOne(context.Context, *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}
