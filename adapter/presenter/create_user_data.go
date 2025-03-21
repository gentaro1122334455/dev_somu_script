package presenter

import (
	"github.com/gsabadini/go-clean-architecture/domain"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type createUserDataPresenter struct{}

func NewCreateUserDataPresenter() usecase.CreateUserDataPresenter {
	return createUserDataPresenter{}
}

func (a createUserDataPresenter) Output(userData domain.UserData) usecase.CreateUserDataOutput {
	return usecase.CreateUserDataOutput{
		ID: userData.ID().Int(),
	}
}
