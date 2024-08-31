package handler

import (
	"go_crud/database"
	"go_crud/model/user/entity"
	"go_crud/model/user/request"
	"go_crud/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(ctx *fiber.Ctx) error {

	LoginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(LoginRequest); err != nil {
		return err
	}
	log.Println(LoginRequest)

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(LoginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	//check available email
	var user entity.User
	err := database.DB.First(&user, "email = ?", LoginRequest.Email).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//check validation password
	isValid := utils.CheckPasswordHash(LoginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//generate JWT (generate token)
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(404).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}