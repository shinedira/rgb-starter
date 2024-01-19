package user

import "github.com/gofiber/fiber/v3"

type ModuleUser struct{}

func (m *ModuleUser) User(route fiber.Router) (fiber.Router, []any) {
	group := route.Group("/user")

	return group, []any{
		&ControllerUser{},
	}
}
