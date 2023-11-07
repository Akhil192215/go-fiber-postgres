package helper

import (
	"errors"

	"github.com/Akhil192215/go-fiber/config"
	"github.com/Akhil192215/go-fiber/models"
)

type User struct {
	//serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName}
}

func FindUser(id int, user *models.User) error {
	config.Database.Db.Find(&user, "id =?", id)
	if user.ID == 0 {
		return errors.New("no user found for this id")
	}
	return nil
}
