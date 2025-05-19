package middleware

import (
	"net/http"
	"nitro/utilities"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		authHeader := context.Request().Header.Get("Authorization")
		authHeaderSplits := strings.Split(authHeader, " ")

		exception := &utilities.Exception{StatusCode: http.StatusUnauthorized, Message: "Not authenticated", Error: "AUTH_004"}

		if len(authHeaderSplits) != 2 {
			return utilities.ThrowException(context, exception)
		}

		token := authHeaderSplits[1]

		decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !decodedToken.Valid {
			return utilities.ThrowException(context, exception)
		}

		issuer, _ := decodedToken.Claims.GetIssuer()

		if issuer != "carbon" {
			return utilities.ThrowException(context, exception)
		}

		return next(context)
	}
}
