package repository

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gsabadini/go-clean-architecture/domain"
	"github.com/pkg/errors"
)

type UserDataNoSQL struct {
	db NoSQL
}

func NewUserDataNoSQL(db NoSQL) UserDataNoSQL {
	return UserDataNoSQL{
		db: db,
	}
}

func (u UserDataNoSQL) Create(ctx context.Context, userData domain.UserData) (domain.UserData, error) {

	param1 := &dynamodb.PutItemInput{
		TableName: aws.String("user_data"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(string(userData.ID().Int())), //number
			},
		},
	}

	if err := u.db.PutItem(ctx, param1); err != nil {
		return domain.UserData{}, errors.Wrap(err, "error creating transfer")
	}

	return userData, nil
}
