package app

import (
	"log"

	"github.com/mrtyunjaygr8/passwd/ent"
	"github.com/mrtyunjaygr8/passwd/ent/user"
	"github.com/mrtyunjaygr8/passwd/utils"
)

func (a *App) CreateUser(email, password string) (*ent.User, error) {
	user, err := a.Client.User.Create().SetEmail(email).SetPassword(password).Save(a.Context)
	if err != nil {
		log.Printf("err creating user, %v: %v", email, err)
		return &ent.User{}, err
	}

	return user, nil
}

func (a *App) LoginUser(email, password string) error {
	user, err := a.Client.User.Query().Where(user.Email(email)).Only(a.Context)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			log.Printf("user: %v not found", email)
			return utils.NOT_FOUND
		}
		log.Printf("err querying user, %v, %v", email, err)
		return err
	}

	if password != user.Password {
		return utils.BAD_REQUEST
	}
	return nil
}
