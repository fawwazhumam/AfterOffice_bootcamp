package handler

import (
	"go_crud/database"
	"go_crud/model/user/entity"
	"go_crud/model/user/request"
	"go_crud/model/user/response"
	"go_crud/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {

	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result != nil {

		log.Println(result.Error)
	}

	return ctx.JSON(users)

}

func UserHandlerCreate(ctx *fiber.Ctx) error {

	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	//membuat validate request
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "validation failed",
			"error": err.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address, 
		Phone:   user.Phone,
		// Password: user.Password,
		UpdateAt: time.Now(),
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   errCreateUser.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"massage": "user not found",
		})
	}

	userResponse := response.UserResponse{

		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
		CreatedAt: user.CreatedAt,
		UpdateAt: user.UpdateAt,
	}

	return ctx.JSON(fiber.Map{
		"massage": "success",
		"data": userResponse,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {

	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad reques",
		})
	}

	var user entity.User

	userId := ctx.Params("id")

	//check available user
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//update user data
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": "user",
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")
	var user entity.User

	//check available user
	err := database.DB.Debug().First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"massage": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"massage": "internal server error",
		})
	}
	
	return ctx.JSON(fiber.Map{
		"massage": "user was delete",
	})
}