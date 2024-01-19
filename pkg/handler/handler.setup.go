package handler

import (
	"reflect"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	*fiber.App
}

func BootHandler() *Handler {
	app := fiber.New()

	// do some default config here

	return &Handler{app}
}

func (h *Handler) RegisterModule(modules []any) {
	for _, module := range modules {
		h.registerHandler(module, h.Group(""))
	}
}

func (h *Handler) registerHandler(module any, args ...any) {
	var route reflect.Value
	if len(args) >= 1 {
		route = reflect.ValueOf(args[0])
	}

	methods := reflect.TypeOf(module)
	_, isHandler := methods.MethodByName("Boot")

	if !isHandler {
		for i := 0; i < methods.NumMethod(); i++ {
			method := methods.Method(i)

			execMethod := reflect.ValueOf(module).MethodByName(method.Name).Call([]reflect.Value{route})

			for _, instance := range execMethod[1].Interface().([]any) {
				h.registerHandler(instance, execMethod[0].Interface())
			}
		}
	} else {
		reflect.ValueOf(module).MethodByName("Boot").Call([]reflect.Value{route})
	}
}
