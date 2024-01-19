package user

import "github.com/gofiber/fiber/v3"

type ControllerUser struct {
	UserService ServiceUserInterface
}

func (c *ControllerUser) Boot(route fiber.Router) {
	// register service
	c.UserService = &ServiceUser{}

	// register route here
	route.Get("/", c.Index)
	route.Get("/:id", c.Show)
	route.Post("/", c.Store)
	route.Put("/:id", c.Update)
	route.Delete("/:id", c.Destroy)
}

func (c *ControllerUser) Index(ctx fiber.Ctx) error {
	// fetch data from service
	_ = c.UserService.FindAll()

	return ctx.SendString("Index")
}

func (c *ControllerUser) Store(ctx fiber.Ctx) error {
	// store data trough service
	_ = c.UserService.Store()

	return ctx.SendString("Store")
}

func (c *ControllerUser) Show(ctx fiber.Ctx) error {
	// fetch data from service
	_ = c.UserService.FindById(1)

	return ctx.SendString("Show")
}

func (c *ControllerUser) Update(ctx fiber.Ctx) error {
	// update data trough service
	_ = c.UserService.Update(1)

	return ctx.SendString("Update")
}

func (c *ControllerUser) Destroy(ctx fiber.Ctx) error {
	// delete data trough service
	_ = c.UserService.Delete(1)

	return ctx.SendString("Destroy")
}
