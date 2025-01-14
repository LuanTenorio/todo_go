package requestError

import (
	"github.com/labstack/echo/v4"
)

const (
	IncompatibleBody = "The body is incompatible with the desired"
)

type requestErrorBody[M string | []string] struct {
	statusCode int
	message    M
}

func NewError(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, requestErrorBody[string]{
		statusCode: statusCode,
		message:    message,
	})
}
