package main

import (
	"fmt"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	api "github.com/k5ki/openapigen-test-go/api/gen"
	"github.com/labstack/echo/v4"
)

func RegisterServer(e *echo.Echo, s *Server) *echo.Echo {
	h := api.NewStrictHandler(s, []api.StrictMiddlewareFunc{})
	api.RegisterHandlers(e, h)

	// Enable validation
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil
	e.Use(oapiMiddleware.OapiRequestValidator(swagger))

	return e
}

func main() {
	e := echo.New()
	RegisterServer(e, &Server{})

	fmt.Println(e.Start(fmt.Sprintf(":8080")))
}
