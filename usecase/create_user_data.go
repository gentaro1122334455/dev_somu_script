package usecase

import (
	"context"
	"time"

	"github.com/gsabadini/go-clean-architecture/domain"
)

type (
	// CreateUserDataUseCase input port
	CreateUserDataUseCase interface {
		Execute(context.Context, CreateUserDataInput) (CreateUserDataOutput, error)
	}

	// CreateUserDataInput input data
	CreateUserDataInput struct {
		id int `dynamo:"id,hash"`
	}

	// CreateUserDataPresenter output port
	CreateUserDataPresenter interface {
		Output(domain.UserData) CreateUserDataOutput
	}

	// CreateUserDataOutput output data
	CreateUserDataOutput struct {
		ID int `json:"id"`
	}

	createUserDataInteractor struct {
		repo       domain.UserDataRepository
		presenter  CreateUserDataPresenter
		ctxTimeout time.Duration
	}
)

// NewCreateUserDataInteractor creates new createUserDataInteractor with its dependencies
func NewCreateUserDataInteractor(
	repo domain.UserDataRepository,
	presenter CreateUserDataPresenter,
	t time.Duration,
) CreateUserDataUseCase {
	return createUserDataInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (u createUserDataInteractor) Execute(ctx context.Context, input CreateUserDataInput) (CreateUserDataOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var userData = domain.NewUserData(
		domain.ID(input.id),
	)

	userData, err := u.repo.Create(ctx, userData)
	if err != nil {
		return u.presenter.Output(domain.UserData{}), err
	}

	return u.presenter.Output(userData), nil
}
