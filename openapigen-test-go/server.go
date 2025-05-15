package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	api "github.com/k5ki/openapigen-test-go/api/gen"
	"github.com/oapi-codegen/runtime/types"
)

func Some[T any](x T) *T {
	return &x
}

var _ api.StrictServerInterface = &Server{}

type Server struct{}

func (s *Server) GetUser(ctx context.Context, req api.GetUserRequestObject) (api.GetUserResponseObject, error) {
	return api.GetUser200JSONResponse(
		api.User{
			Id:        Some(types.UUID(uuid.MustParse("fd6944af-e7e4-4e35-9a45-43c57b7a648a"))),
			Name:      Some("Jacque Yue"),
			CreatedAt: Some(time.Now()),
			UpdatedAt: Some(time.Now()),
		},
	), nil
}

func (s *Server) CreateUser(ctx context.Context, req api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	return api.CreateUser200JSONResponse(
		api.User{
			Id:        Some(types.UUID(uuid.MustParse("fd6944af-e7e4-4e35-9a45-43c57b7a648a"))),
			Name:      Some("Jacque Yue"),
			CreatedAt: Some(time.Now()),
			UpdatedAt: Some(time.Now()),
		},
	), nil
}
