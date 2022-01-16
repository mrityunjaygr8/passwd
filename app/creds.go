package app

import (
	"log"

	"github.com/mrtyunjaygr8/passwd/ent"
	"github.com/mrtyunjaygr8/passwd/ent/creds"
	"github.com/mrtyunjaygr8/passwd/ent/passwords"
	"github.com/mrtyunjaygr8/passwd/utils"
)

func (a *App) ListCreds(token string) []*ent.Creds {
	user, err := a.GetUser(token)
	if err != nil {
		log.Fatal("an error has occurred: ", err)
	}

	creds, err := user.QueryCreds().All(a.Context)
	if err != nil {
		log.Fatal(err)
	}
	return creds
}

func (a *App) CreateCreds(token, name, username, password, url string) (*ent.Creds, error) {
	user, err := a.GetUser(token)
	if err != nil {
		log.Fatal("an error has occurred: ", err)
	}

	cred, err := a.Client.Creds.Create().SetUsername(username).SetUser(user).SetName(name).SetURL(url).Save(a.Context)
	if err != nil {
		log.Printf("an error occurred while creating the new cred: %v", err)
		return &ent.Creds{}, utils.BAD_REQUEST
	}

	_, err = a.Client.Passwords.Create().SetPassword(password).SetCred(cred).Save(a.Context)
	if err != nil {
		log.Printf("an error occurred while creating the new cred: %v", err)
		return &ent.Creds{}, utils.BAD_REQUEST
	}

	return cred, nil
}

func (a *App) GetCred(token, name string) (*ent.Creds, string, error) {
	user, err := a.GetUser(token)
	if err != nil {
		log.Fatal(err)
	}

	cred, err := user.QueryCreds().Where(creds.Name(name)).Only(a.Context)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			log.Printf("cred: %v not found", name)
			return &ent.Creds{}, "", utils.NOT_FOUND
		}
		log.Printf("err querying cred, %v, %v", name, err)
		return &ent.Creds{}, "", err

	}

	pass, _ := cred.QueryPasswords().Order(ent.Desc(passwords.FieldCreateTime)).Limit(1).All(a.Context)
	if err != nil {
		log.Println(err)
		return &ent.Creds{}, "", utils.NOT_FOUND

	}
	return cred, pass[0].Password, nil
}

func (a *App) DeleteCred(token, name string) error {
	user, err := a.GetUser(token)
	if err != nil {
		log.Fatal(err)
	}

	cred, err := user.QueryCreds().Where(creds.Name(name)).Only(a.Context)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			log.Printf("cred: %v not found", name)
			return utils.NOT_FOUND
		}
		log.Printf("err querying cred, %v, %v", name, err)
		return err

	}

	_, err = a.Client.Creds.Delete().Where(creds.ID(cred.ID)).Exec(a.Context)
	if err != nil {
		log.Printf("cred: %v can not be deleted", cred.Name)
		return utils.BAD_REQUEST
	}

	return nil
}

func (a *App) UpdateCred(token, name, new_pass string) (*ent.Creds, error) {
	user, err := a.GetUser(token)
	if err != nil {
		log.Fatal(err)
	}

	cred, err := user.QueryCreds().Where(creds.Name(name)).Only(a.Context)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			log.Printf("cred: %v not found", name)
			return &ent.Creds{}, utils.NOT_FOUND
		}
		log.Printf("err querying cred, %v, %v", name, err)
		return &ent.Creds{}, err

	}
	_, err = a.Client.Passwords.Create().SetPassword(new_pass).SetCred(cred).Save(a.Context)
	if err != nil {
		log.Printf("cred: %v, error updating new password", cred.Name)
		return &ent.Creds{}, utils.BAD_REQUEST
	}

	return cred, nil
}

func (a *App) HistoryCreds(token, name string) ([]*ent.Passwords, error) {
	user, err := a.GetUser(token)
	if err != nil {
		log.Fatal(err)
	}

	cred, err := user.QueryCreds().Where(creds.Name(name)).Only(a.Context)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			log.Printf("cred: %v not found", name)
			return []*ent.Passwords{}, utils.NOT_FOUND
		}
		log.Printf("err querying cred, %v, %v", name, err)
		return []*ent.Passwords{}, err

	}

	passwords, err := cred.QueryPasswords().Order(ent.Desc(passwords.FieldCreateTime)).All(a.Context)
	if err != nil {
		log.Printf("cred: %v, error getting cred history", name)
		return []*ent.Passwords{}, err
	}

	return passwords, nil
}
