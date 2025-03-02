package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}


func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}


func (a *ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty parameters")
	}

	//does a user already exist?
	userExists, err := a.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return err
	}

	if userExists {
		return fmt.Errorf("user already exists")
	}

	//create the user
	err = a.dbStore.CreateUser(event)
	if err != nil {
		return err
	}

	return nil
}
