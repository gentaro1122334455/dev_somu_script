package action

import (
	"encoding/json"
	"net/http"

	"github.com/gsabadini/go-clean-architecture/adapter/api/logging"
	"github.com/gsabadini/go-clean-architecture/adapter/api/response"
	"github.com/gsabadini/go-clean-architecture/adapter/logger"
	"github.com/gsabadini/go-clean-architecture/adapter/validator"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

type CreateUserDataAction struct {
	uc        usecase.CreateUserDataUseCase
	log       logger.Logger
	validator validator.Validator
}

func NewCreateUserDataAction(uc usecase.CreateUserDataUseCase, log logger.Logger, v validator.Validator) CreateUserDataAction {
	return CreateUserDataAction{
		uc:        uc,
		log:       log,
		validator: v,
	}
}

func (u CreateUserDataAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_userData"

	var input usecase.CreateUserDataInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			u.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error when decoding json")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if errs := u.validateInput(input); len(errs) > 0 {
		logging.NewError(
			u.log,
			response.ErrInvalidInput,
			logKey,
			http.StatusBadRequest,
		).Log("invalid input")

		response.NewErrorMessage(errs, http.StatusBadRequest).Send(w)
		return
	}

	output, err := u.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(
			u.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when creating a new userData")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(u.log, logKey, http.StatusCreated).Log("success creating userData")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

func (u CreateUserDataAction) validateInput(input usecase.CreateUserDataInput) []string {
	var msgs []string

	err := u.validator.Validate(input)
	if err != nil {
		for _, msg := range u.validator.Messages() {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}
