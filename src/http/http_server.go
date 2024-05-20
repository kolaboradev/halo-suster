package httpServer

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kolaboradev/halo-suster/src/helper"
	patientController "github.com/kolaboradev/halo-suster/src/http/controllers/patient"
	recordController "github.com/kolaboradev/halo-suster/src/http/controllers/record"
	userController "github.com/kolaboradev/halo-suster/src/http/controllers/user"
	"github.com/kolaboradev/halo-suster/src/http/middlewares"
	v1Routes "github.com/kolaboradev/halo-suster/src/http/routes/v1"
	medicalV1Routes "github.com/kolaboradev/halo-suster/src/http/routes/v1/medical"
	userV1Routes "github.com/kolaboradev/halo-suster/src/http/routes/v1/user"
	patientRepository "github.com/kolaboradev/halo-suster/src/repositories/patient"
	recordRepository "github.com/kolaboradev/halo-suster/src/repositories/record"
	userRepository "github.com/kolaboradev/halo-suster/src/repositories/user"
	patientService "github.com/kolaboradev/halo-suster/src/services/patient"
	recordService "github.com/kolaboradev/halo-suster/src/services/record"
	userService "github.com/kolaboradev/halo-suster/src/services/user"
)

type HttpServer struct {
	DB *sql.DB
}

func NewServer(db *sql.DB) HttpServerInterface {
	return &HttpServer{DB: db}
}

func (server *HttpServer) Listen() {
	validator := validator.New()
	validator.RegisterValidation("nip_it", helper.IsValidItNIP)
	validator.RegisterValidation("nip_nurse", helper.IsValidNurseNIP)
	validator.RegisterValidation("url_image", helper.IsValidUrl)
	validator.RegisterValidation("identity_number", helper.IdentityNumber)
	validator.RegisterValidation("gender", helper.IsGender)

	app := fiber.New(fiber.Config{
		ServerHeader: "Kolaboradev",
		ErrorHandler: middlewares.ErrorHandle,
	})

	userRepo := userRepository.NewUserRepository()
	userService := userService.NewUserService(server.DB, validator, userRepo)
	userController := userController.NewUserController(userService)

	patientRepo := patientRepository.NewPatientRepo()
	patientService := patientService.NewPatientService(server.DB, validator, patientRepo)
	patientController := patientController.NewPatientController(patientService)

	recordRepo := recordRepository.NewRecordRepository()
	recordService := recordService.NewRecordService(server.DB, validator, recordRepo, patientRepo, userRepo)
	recordController := recordController.NewRecordController(recordService)

	app.Use(recover.New())
	v1 := v1Routes.SetRoutesV1(app)
	userV1Routes.SetRoutesUsers(v1, userController)
	medicalV1Routes.SetRoutesMedicals(v1, patientController, recordController)

	app.Listen(":8080")
}
