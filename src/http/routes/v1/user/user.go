package userV1Routes

import (
	"github.com/gofiber/fiber/v2"
	userController "github.com/kolaboradev/halo-suster/src/http/controllers/user"
	"github.com/kolaboradev/halo-suster/src/http/middlewares"
)

func SetRoutesUsers(router fiber.Router, uc userController.UserControllerInterface) {
	router.Post("/image", middlewares.AuthMiddleware, uc.PostImage)
	userGrup := router.Group("/user")

	userGrup.Post("/it/register", uc.RegisterIt)
	userGrup.Post("/it/login", uc.LoginIt)
	userGrup.Post("/nurse/register", middlewares.AuthMiddleware, middlewares.AuthorizeMiddleware("it"), uc.RegisterNurse)
	userGrup.Post("/nurse/login", uc.LoginNurse)
	userGrup.Get("", middlewares.AuthMiddleware, middlewares.AuthorizeMiddleware("it"), uc.GetAllUsers)
	userGrup.Put("/nurse/:userId", middlewares.AuthMiddleware, middlewares.AuthorizeMiddleware("it"), uc.EditNurseById)
	userGrup.Delete("/nurse/:userId", middlewares.AuthMiddleware, middlewares.AuthorizeMiddleware("it"), uc.DeleteNurseById)
	userGrup.Post("/nurse/:userId/access", middlewares.AuthMiddleware, middlewares.AuthorizeMiddleware("it"), uc.GetAccessNurse)
}
