package database

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

type dynamoHandler struct {
	db *dynamo.DB
}

func NewDynamoDBHandler(c *dynamodbConfig) *dynamoHandler {
	cred := credentials.NewStaticCredentials(c.accessKey, c.secretKey, "")

	dynamoDB := dynamo.New(session.New(), &aws.Config{
		Credentials: cred,
		Region:      aws.String(c.region),
	})

	return &dynamoHandler{
		db: dynamoDB,
	}
}

func (dynamo dynamoHandler) FindOne(ctx context.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	resp, err := dynamo.db.Client().GetItem(input)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (dynamo dynamoHandler) PutItem(ctx context.Context, input *dynamodb.PutItemInput) error {
	_, err := dynamo.db.Client().PutItem(input)

	if err != nil {
		return err
	}

	return nil
}
