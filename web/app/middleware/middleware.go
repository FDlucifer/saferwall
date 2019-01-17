// Copyright 2018 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func requireJSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("content-type")
		fmt.Println(c.Request().Header)
		if contentType != "application/json" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return next(c)
	}
}

// ValidateAPIKey will check if API key is allowed
func ValidateAPIKey(key string, ctx echo.Context) (bool, error) {

	return true, nil
}

// Init middlewares
func Init(e *echo.Echo) {

	// logging
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// require JSON
	// e.Use(requireJSON)

	// authorization
	// e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	// 	KeyLookup: "query:api-key",
	// 	Validator: middleware.KeyAuthValidator(ValidateAPIKey),
	// }))

}