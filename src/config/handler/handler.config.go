package handler

import (
	"rgb/starter/src/modules/apps"
	"rgb/starter/src/utils/app"

	hdl "rgb/starter/pkg/handler"
)

type Handler struct{}

func (c *Handler) Boot() {
	h := hdl.BootHandler()

	// register modules
	h.RegisterModule([]any{
		&apps.ModuleApp{},
		// other module here
	})

	app.App.Handler = h
}
