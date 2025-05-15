package main

import (
	"fmt"

	api "github.com/k5ki/openapigen-test-go/api/gen"
	"github.com/labstack/echo/v4"
)

func RegisterServer(e *echo.Echo, s *Server) *echo.Echo {
	h := api.NewStrictHandler(s, []api.StrictMiddlewareFunc{})
	api.RegisterHandlers(e, h)
	return e
}

func main() {
	e := echo.New()
	RegisterServer(e, &Server{})

	fmt.Println(e.Start(fmt.Sprintf(":8080")))
}
