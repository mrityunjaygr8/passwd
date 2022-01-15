package app

import (
	"log"

	"github.com/mrtyunjaygr8/passwd/ent"
)

func (a *App) CreateUser(email, password string) (*ent.User, error) {
	user, err := a.Client.User.Create().SetEmail(email).SetPassword(password).Save(a.Context)
	if err != nil {
		log.Printf("err creating user, %v: %v", email, err)
		return &ent.User{}, err
	}

	return user, nil

}
