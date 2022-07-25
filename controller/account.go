package controller

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/go-openapi/runtime/middleware"
	"github.com/openziti-test-kitchen/zrok/controller/store"
	"github.com/openziti-test-kitchen/zrok/rest_model"
	"github.com/openziti-test-kitchen/zrok/rest_zrok_server/operations/identity"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func createAccountHandler(params identity.CreateAccountParams) middleware.Responder {
	logrus.Infof("received account request for username '%v'", params.Body.Username)
	if params.Body == nil || params.Body.Username == "" || params.Body.Password == "" {
		return middleware.Error(500, errors.Errorf("invalid username or password"))
	}

	token, err := generateApiToken()
	if err != nil {
		logrus.Errorf("error generating api token: %v", err)
		return middleware.Error(500, err.Error())
	}

	a := &store.Account{
		Username: params.Body.Username,
		Password: hashPassword(params.Body.Password),
		Token:    token,
	}
	tx, err := str.Begin()
	if err != nil {
		logrus.Errorf("error starting transaction: %v", err)
		return middleware.Error(500, err.Error())
	}
	id, err := str.CreateAccount(a, tx)
	if err != nil {
		logrus.Errorf("error creating account: %v", err)
		_ = tx.Rollback()
		return middleware.Error(400, err.Error())
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("error comitting: %v", err)
	}

	logrus.Infof("account created with id = '%v'", id)
	return identity.NewCreateAccountCreated().WithPayload(&rest_model.AccountResponse{
		Token: token,
	})
}

func hashPassword(raw string) string {
	hash := sha512.New()
	hash.Write([]byte(raw))
	return hex.EncodeToString(hash.Sum(nil))
}
