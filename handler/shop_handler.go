package handler

import (
	"go_crud/database"
	"go_crud/model/shop/entity"
	"go_crud/model/shop/request"
	"go_crud/model/shop/response"
	"go_crud/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ShopHandlerGetAll(ctx *fiber.Ctx) error {

	var shop []entity.Shop
	result := database.DB.Debug().Find(&shop)

	if result != nil {

		log.Println(result.Error)
	}

	return ctx.JSON(shop)

}

func ShopHandlerCreate(ctx *fiber.Ctx) error {

	shop := new(request.ShopCreateRequest)

	if err := ctx.BodyParser(shop); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	//membuat validate request
	validate := validator.New()
	if err := validate.Struct(shop); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "validation failed",
			"error": err.Error(),
		})
	}

	newShop := entity.Shop{
		Name:    shop.Name,
		Email:   shop.Email,
		Address: shop.Address, 
		Contact:   shop.Contact,
		// Password: user.Password,
		UpdateAt: time.Now(),
	}

	hashedPassword, err := utils.HashPassword(shop.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	newShop.Password = hashedPassword

	errCreateUser := database.DB.Create(&newShop).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   errCreateUser.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    newShop,
	})
}

func ShopHandlerGetById(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")

	var user entity.Shop
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"massage": "user not found",
		})
	}

	shopResponse := response.ShopResponse{

		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Contact: user.Contact,
		CreatedAt: user.CreatedAt,
		UpdateAt: user.UpdateAt,
	}

	return ctx.JSON(fiber.Map{
		"massage": "success",
		"data": shopResponse,
	})
}

func ShopHandlerUpdate(ctx *fiber.Ctx) error {

	userRequest := new(request.ShopUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad reques",
		})
	}

	var shop entity.Shop

	shopId := ctx.Params("id")

	//check available user
	err := database.DB.First(&shop, "id = ?", shopId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//update user data
	if userRequest.Name != "" {
		shop.Name = userRequest.Name
	}
	shop.Address = userRequest.Address
	shop.Contact = userRequest.Contact
	errUpdate := database.DB.Save(&shop).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": "nama toko",
	})
}

func ShopHandlerDelete(ctx *fiber.Ctx) error {

	shopId := ctx.Params("id")
	var shop entity.Shop

	//check available user
	err := database.DB.Debug().First(&shop, "id = ?", shopId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"massage": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&shop).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"massage": "internal server error",
		})
	}
	
	return ctx.JSON(fiber.Map{
		"massage": "user was delete",
	})
}