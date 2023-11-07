package routes

import (
	"github.com/Akhil192215/go-fiber/config"
	"github.com/Akhil192215/go-fiber/helper"
	"github.com/Akhil192215/go-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser((&user)); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	config.Database.Db.Create(&user)
	responseUser := helper.CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	config.Database.Db.Find(&users)
	responseUsers := []helper.User{}
	for _, user := range users {
		responseUser := helper.CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := helper.FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := helper.CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = helper.FindUser(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	config.Database.Db.Save(&user)

	responseUser := helper.CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = helper.FindUser(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := config.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully deleted user")
}
