package medicalV1Routes

import (
	"github.com/gofiber/fiber/v2"
	patientController "github.com/kolaboradev/halo-suster/src/http/controllers/patient"
	recordController "github.com/kolaboradev/halo-suster/src/http/controllers/record"
	"github.com/kolaboradev/halo-suster/src/http/middlewares"
)

func SetRoutesMedicals(router fiber.Router, pc patientController.PatientControllerInterface, rc recordController.RecordControllerInterface) {
	medicalGrup := router.Group("/medical")

	medicalGrup.Post("/patient", middlewares.AuthMiddleware, pc.Create)
	medicalGrup.Get("/patient", middlewares.AuthMiddleware, pc.GetAll)
	medicalGrup.Post("/record", middlewares.AuthMiddleware, rc.Create)
	medicalGrup.Get("/record", middlewares.AuthMiddleware, rc.GetAll)
}
