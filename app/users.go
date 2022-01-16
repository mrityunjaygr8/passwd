package app

import (
	"log"
	"time"

	"github.com/mrtyunjaygr8/passwd/ent"
	"github.com/mrtyunjaygr8/passwd/ent/user"
	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/o1egl/paseto"
)

func (a *App) CreateUser(email, password string) (*ent.User, error) {
	user, err := a.Client.User.Create().SetEmail(email).SetPassword(password).Save(a.Context)
	if err != nil {
		log.Printf("err creating user, %v: %v", email, err)
		return &ent.User{}, err
	}

	return user, nil
}

func (a *App) LoginUser(email, password string) (string, error) {
	user, err := a.Client.User.Query().Where(user.Email(email)).Only(a.Context)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			log.Printf("user: %v not found", email)
			return "", utils.NOT_FOUND
		}
		log.Printf("err querying user, %v, %v", email, err)
		return "", err
	}

	if password != user.Password {
		return "", utils.BAD_REQUEST
	}

	token, err := signToken(*user)
	if err != nil {
		log.Printf("err signing token: %v", err)
		return "", utils.BAD_REQUEST
	}
	return token, nil
}

func (a *App) GetUser(token string) (*ent.User, error) {
	user_email, err := verifyToken(token)
	if err != nil {
		log.Printf("err getting user: %v", err)
		return &ent.User{}, utils.BAD_REQUEST
	}

	user, err := a.Client.User.Query().Where(user.Email(user_email)).Only(a.Context)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			log.Printf("user: %v not found", user_email)
			return &ent.User{}, utils.NOT_FOUND
		}
		log.Printf("err querying user, %v, %v", user_email, err)
		return &ent.User{}, err
	}

	return user, nil
}

func signToken(user ent.User) (string, error) {
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "test",
		Issuer:     "test_service",
		Jti:        "test_service-test-jti",
		Subject:    "test_service-test-subject",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}

	jsonToken.Set("user", user.Email)

	return paseto.NewV2().Encrypt(utils.SYMMETRIC_KEY, jsonToken, "")
}

func verifyToken(token string) (string, error) {
	if token == "" {
		return "", utils.BAD_REQUEST
	}
	var paseto_token paseto.JSONToken
	var footer string
	err := paseto.NewV2().Decrypt(token, utils.SYMMETRIC_KEY, &paseto_token, &footer)
	if err != nil {
		log.Printf("An error has occurred while decryting the token: %v", err)
		return "", err
	}

	return paseto_token.Get("user"), nil
}
