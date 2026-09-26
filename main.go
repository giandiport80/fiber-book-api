package main

import (
	"fmt"
	"log"

	"github.com/giandiport80/fiber-book-api/internal/api"
	"github.com/giandiport80/fiber-book-api/internal/config"
	"github.com/giandiport80/fiber-book-api/internal/connection"
	"github.com/giandiport80/fiber-book-api/internal/repository"
	"github.com/giandiport80/fiber-book-api/internal/service"
	"github.com/gofiber/fiber/v3"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)
	api.NewCustomer(app, customerService)

	app.Get("/developer", developers)
	addr := fmt.Sprintf("%s:%s", cnf.Server.Host, cnf.Server.Port)
	log.Fatal(app.Listen(addr))
}

func developers(ctx fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Hello",
	})
}
