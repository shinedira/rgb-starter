package apps

import (
	"rgb/starter/src/modules/apps/user"

	"github.com/gofiber/fiber/v3"
)

type ModuleApp struct{}

func (m *ModuleApp) ApiV1(route fiber.Router) (fiber.Router, []any) {
	group := route.Group("/api/v1/app")

	return group, []any{
		&user.ModuleUser{},
		// &cms.ModuleCms{},
		// &auth.ModuleAuth{},
		// other modules
	}
}
