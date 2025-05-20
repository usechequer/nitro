package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/usechequer/utilities"
)

func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		token := context.Get("token").(utilities.Token)
		exception := &utilities.Exception{StatusCode: http.StatusUnauthorized, Message: "Not authenticated", Error: "AUTH_004"}

		if token.Issuer != "carbon" {
			return utilities.ThrowException(context, exception)
		}

		return next(context)
	}
}
