package server

import (
	"context"
	"net/http"

	"github.com/davidporos92/margin-cms/api/internal/api"
)

func (s server) GetHealth(ctx context.Context, request api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	return api.GetHealth200JSONResponse{
		Status: http.StatusText(http.StatusOK),
	}, nil
}
