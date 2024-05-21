package userController

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/halo-suster/src/exceptions"
	"github.com/kolaboradev/halo-suster/src/helper"
	userRequest "github.com/kolaboradev/halo-suster/src/models/web/request/user"
	"github.com/kolaboradev/halo-suster/src/models/web/response"
	userResponse "github.com/kolaboradev/halo-suster/src/models/web/response/user"
	userService "github.com/kolaboradev/halo-suster/src/services/user"
)

type UserController struct {
	userService userService.UserServiceInterface
}

func NewUserController(userService userService.UserServiceInterface) UserControllerInterface {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) RegisterIt(c *fiber.Ctx) error {
	userRequestParse := userRequest.UserItCreateParse{}
	if err := c.BodyParser(&userRequestParse); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}
	nipStr := strconv.Itoa(userRequestParse.Nip)
	userItRequest := userRequest.UserItCreate{
		Nip:      nipStr,
		Name:     userRequestParse.Name,
		Password: userRequestParse.Password,
	}
	userResponse := controller.userService.CreateIt(context.Background(), userItRequest)

	c.Set("X-Author", "Kolaboradev")
	return c.Status(201).JSON(response.Web{
		Message: "User registered successfully",
		Data:    userResponse,
	})
}

func (controller *UserController) LoginIt(c *fiber.Ctx) error {
	userRequestParse := userRequest.UserItLoginParse{}
	if err := c.BodyParser(&userRequestParse); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}
	if userRequestParse.Nip == nil {
		panic(exceptions.NewBadRequestError("nip required"))
	}
	nipStr := strconv.Itoa(*userRequestParse.Nip)

	userItRequest := userRequest.UserItLogin{
		Nip:      nipStr,
		Password: userRequestParse.Password,
	}

	userResponse := controller.userService.LoginIt(context.Background(), userItRequest)

	c.Set("X-Author", "Kolaboradev")
	return c.Status(200).JSON(response.Web{
		Message: "User login successfully",
		Data:    userResponse,
	})
}

func (controller *UserController) RegisterNurse(c *fiber.Ctx) error {
	userRequestParse := userRequest.UserNurseCreateParse{}
	if err := c.BodyParser(&userRequestParse); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}
	if userRequestParse.Nip == nil {
		panic(exceptions.NewBadRequestError("nip required"))
	}
	nipStr := strconv.Itoa(*userRequestParse.Nip)
	userRequest := userRequest.UserNurseCreate{
		Nip:                   nipStr,
		Name:                  userRequestParse.Name,
		IdentityCardScanImage: userRequestParse.IdentityCardScanImage,
	}

	userResponse := controller.userService.CreateNurse(context.Background(), userRequest)
	c.Set("X-Author", "Kolaboradev")
	return c.Status(201).JSON(response.Web{
		Message: "User registered successfully",
		Data:    userResponse,
	})
}

func (controller *UserController) LoginNurse(c *fiber.Ctx) error {
	userRequestParse := userRequest.UserNurseParse{}
	if err := c.BodyParser(&userRequestParse); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}
	if userRequestParse.Nip == nil {
		panic(exceptions.NewBadRequestError("nip required"))
	}
	nipStr := strconv.Itoa(*userRequestParse.Nip)
	userRequest := userRequest.UserNurseLogin{
		Nip:      nipStr,
		Password: userRequestParse.Password,
	}

	userResponse := controller.userService.LoginNurse(context.Background(), userRequest)
	c.Set("X-Author", "Kolaboradev")
	return c.Status(200).JSON(response.Web{
		Message: "User registered successfully",
		Data:    userResponse,
	})
}

func (controller *UserController) GetAllUsers(c *fiber.Ctx) error {
	filters := userRequest.UserFilter{
		UserId:    c.Query("userId"),
		Limit:     c.QueryInt("limit", 5),
		Offset:    c.QueryInt("offset", 0),
		Name:      c.Query("name", ""),
		Nip:       helper.QueryIntPointer(c, "nip"),
		Role:      c.Query("role"),
		CreatedAt: c.Query("createdAt", ""),
	}

	userResponses := controller.userService.GetAllUsers(context.Background(), filters)

	c.Set("X-Author", "Kolaboradev")
	if len(userResponses) == 0 {
		return c.Status(200).JSON(response.Web{
			Message: "Get all users successfully",
			Data:    []interface{}{},
		})
	}

	return c.Status(200).JSON(response.Web{
		Message: "Get all users successfully",
		Data:    userResponses,
	})
}

func (controller *UserController) EditNurseById(c *fiber.Ctx) error {
	userRequestParse := userRequest.UserNurseEditParse{}
	if err := c.BodyParser(&userRequestParse); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}
	if userRequestParse.Nip == nil {
		panic(exceptions.NewBadRequestError("nip required"))
	}
	nipStr := strconv.Itoa(*userRequestParse.Nip)

	userId := c.Params("userId")
	userRequest := userRequest.UserNurseEdit{
		Nip:    nipStr,
		Name:   userRequestParse.Name,
		UserId: userId,
	}

	controller.userService.EditNurseById(context.Background(), userRequest)

	c.Set("X-Author", "Kolaboradev")
	return c.Status(200).JSON(response.Web{
		Message: "Successfully edit nurse",
		Data:    "OK",
	})
}

func (controller *UserController) DeleteNurseById(c *fiber.Ctx) error {
	userId := c.Params("userId")

	controller.userService.DeleteNurseById(context.Background(), userId)

	c.Set("X-Author", "Kolaboradev")
	return c.Status(200).JSON(response.Web{
		Message: "Successfully deleted nurse",
		Data:    "OK",
	})
}

func (controller *UserController) GetAccessNurse(c *fiber.Ctx) error {
	userRequestParse := userRequest.UserNurseAccessParse{}
	if err := c.BodyParser(&userRequestParse); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}
	userId := c.Params("userId")

	userRequest := userRequest.UserNurseAccess{
		UserId:   userId,
		Password: userRequestParse.Password,
	}

	controller.userService.AddNurseAccess(context.Background(), userRequest)

	c.Set("X-Author", "Kolaboradev")
	return c.Status(200).JSON(response.Web{
		Message: "Successfully get access nurse",
		Data:    "OK",
	})
}

func (controller *UserController) PostImage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		panic(exceptions.NewBadRequestError("Invalid data"))
	}

	maxSize := 2 * 1024 * 1024
	minSize := 10 * 1024
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" {
		return fmt.Errorf("invalid file format, only *.jpg or *.jpeg allowed")
	}

	if file.Size > int64(maxSize) {
		panic(exceptions.NewBadRequestError("file maks 2MB"))
	}
	if file.Size < int64(minSize) {
		panic(exceptions.NewBadRequestError("file min 10KB"))
	}

	image := userResponse.Image{
		ImageUrl: fmt.Sprintf("https://www.google.com" + file.Filename),
	}

	c.Set("X-Author", "Kolaboradev")
	return c.Status(200).JSON(response.Web{
		Message: "File uploaded sucessfully",
		Data:    image,
	})
}
