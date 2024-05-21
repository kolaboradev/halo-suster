package patientController

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/halo-suster/src/exceptions"
	"github.com/kolaboradev/halo-suster/src/helper"
	patientRequest "github.com/kolaboradev/halo-suster/src/models/web/request/patient"
	"github.com/kolaboradev/halo-suster/src/models/web/response"
	patientService "github.com/kolaboradev/halo-suster/src/services/patient"
)

type PatientController struct {
	patientService patientService.PatientServiceInterface
}

func NewPatientController(ps patientService.PatientServiceInterface) PatientControllerInterface {
	return &PatientController{
		patientService: ps,
	}
}

func (controller *PatientController) Create(c *fiber.Ctx) error {
	patientRequest := patientRequest.PatientCreate{}
	if err := c.BodyParser(&patientRequest); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	patientResponse := controller.patientService.Create(context.Background(), patientRequest)

	c.Set("X-Author", "Kolaboradev")
	return c.Status(201).JSON(response.Web{
		Message: "Successfully create patient",
		Data:    patientResponse,
	})
}

func (controller *PatientController) GetAll(c *fiber.Ctx) error {
	userFilters := patientRequest.PatientFilter{
		IdentityNumber: helper.QueryInt64Pointer(c, "identityNumber"),
		Limit:          c.QueryInt("limit", 5),
		Offset:         c.QueryInt("offset", 0),
		Name:           c.Query("name", ""),
		PhoneNumber:    c.Query("phoneNumber", ""),
		CreatedAt:      c.Query("createdAt", ""),
	}

	userResponses := controller.patientService.GetAll(context.Background(), userFilters)

	if len(userResponses) == 0 {
		return c.Status(200).JSON(response.Web{
			Message: "Get all users successfully",
			Data:    []interface{}{},
		})
	}

	c.Set("X-Author", "Kolaboradev")
	return c.Status(200).JSON(response.Web{
		Message: "Get all users successfully",
		Data:    userResponses,
	})
}
