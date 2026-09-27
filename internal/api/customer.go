package api

import (
	"context"
	"net/http"
	"time"

	"github.com/giandiport80/fiber-book-api/domain"
	"github.com/giandiport80/fiber-book-api/dto"
	"github.com/giandiport80/fiber-book-api/internal/util"
	"github.com/gofiber/fiber/v3"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := customerApi{
		customerService: customerService,
	}

	app.Get("/customers", ca.Index)
	app.Post("/customers", ca.Create)
	app.Put("/customers/:id", ca.Update)
}

func (ca customerApi) Index(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error(), nil))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))

}

func (ca customerApi) Create(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.Bind().JSON(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseError("Error get data", fails))
	}

	util.TrimStrings(&req)

	err := ca.customerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error(), nil))
	}

	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess(""))
}

func (ca customerApi) Update(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	if err := ctx.Bind().JSON(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseError("validation error", fails))
	}

	req.ID = ctx.Params("id")
	err := ca.customerService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error(), nil))
	}

	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(""))
}
