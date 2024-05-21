package recordController

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kolaboradev/halo-suster/src/exceptions"
	"github.com/kolaboradev/halo-suster/src/helper"
	recordRequest "github.com/kolaboradev/halo-suster/src/models/web/request/record"
	"github.com/kolaboradev/halo-suster/src/models/web/response"
	recordService "github.com/kolaboradev/halo-suster/src/services/record"
)

type RecordController struct {
	recordService recordService.RecordServiceInterface
}

func NewRecordController(rc recordService.RecordServiceInterface) RecordControllerInterface {
	return &RecordController{
		recordService: rc,
	}
}

func (controller *RecordController) Create(c *fiber.Ctx) error {
	recordRequestParse := recordRequest.RecordParse{}
	if err := c.BodyParser(&recordRequestParse); err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	createdBy := c.Locals("userId").(string)

	recordRequest := recordRequest.Record{
		IdentityNumber: recordRequestParse.IdentityNumber,
		Symptoms:       recordRequestParse.Symptoms,
		Medications:    recordRequestParse.Medications,
		CreatedBy:      createdBy,
	}
	controller.recordService.Create(context.Background(), recordRequest)

	c.Set("X-Author", "Kolaboradev")
	return c.Status(201).JSON(response.Web{
		Message: "record created successfully",
		Data:    "OK",
	})
}

func (controller *RecordController) GetAll(c *fiber.Ctx) error {
	recordFilter := recordRequest.RecordFilter{
		IdentityNumber: helper.QueryInt64Pointer(c, "identityDetail.identityNumber"),
		UserId:         c.Query("createdBy.userId", ""),
		Nip:            c.Query("createdBy.nip", ""),
		Limit:          c.QueryInt("limit", 5),
		Offset:         c.QueryInt("offset", 0),
		CreatedAt:      c.Query("createdAt", ""),
	}

	recordResponses := controller.recordService.GetAll(context.Background(), recordFilter)

	if len(recordResponses) == 0 {
		return c.Status(200).JSON(response.Web{
			Message: "Get all users successfully",
			Data:    []interface{}{},
		})
	}

	c.Set("X-Author", "Kolaboradev")
	return c.Status(200).JSON(response.Web{
		Message: "Get all users successfully",
		Data:    recordResponses,
	})
}
