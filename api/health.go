package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/RedBird96/b7s/models/response"
)

// Execute implements the REST API endpoint for function execution.
func (a *API) Health(ctx echo.Context) error {

	// respond with health check
	resp := response.Health{
		Type: "health",
		Code: http.StatusOK,
	}

	// Send the response.
	return ctx.JSON(http.StatusOK, resp)
}
