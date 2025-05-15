package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	api "github.com/k5ki/openapigen-test-go/api/gen"
	"github.com/oapi-codegen/runtime/types"
	"github.com/samber/lo"
)

func Some[T any](x T) *T {
	return &x
}

var _ api.StrictServerInterface = &Server{}

type Server struct{}

func (s *Server) GetUser(ctx context.Context, req api.GetUserRequestObject) (api.GetUserResponseObject, error) {
	return api.GetUser200JSONResponse(
		api.User{
			Id:        Some(types.UUID(uuid.MustParse(req.UserId.String()))),
			Name:      Some("Jacque Yue"),
			CreatedAt: Some(lo.Must(time.Parse(time.RFC3339, "2024-05-26T16:38:11Z"))),
			UpdatedAt: Some(lo.Must(time.Parse(time.RFC3339, "2024-05-26T16:38:11Z"))),
		},
	), nil
}

func (s *Server) CreateUser(ctx context.Context, req api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return api.CreateUser200JSONResponse(
		api.User{
			Id:        Some(types.UUID(id)),
			Name:      Some(req.Body.Name),
			CreatedAt: Some(time.Now()),
			UpdatedAt: Some(time.Now()),
		},
	), nil
}
