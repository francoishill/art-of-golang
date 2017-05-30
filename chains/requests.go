package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func BindRequestInt64Param(c echo.Context, paramName string, val *int64) Loader {
	return func() *Error {
		strVal := strings.TrimSpace(c.Param(paramName))
		if len(strVal) == 0 {
			return &Error{fmt.Errorf("Param %s is required", paramName), http.StatusBadRequest}
		}

		tmpVal, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			return &Error{fmt.Errorf("Cannot parse %s as int", strVal), http.StatusBadRequest}
		}

		*val = tmpVal
		return nil
	}
}

func BindRequestBody(c echo.Context, dest interface{}) Loader {
	return func() *Error {
		if err := c.Bind(dest); err != nil {
			return &Error{err, http.StatusBadRequest}
		}
		return nil
	}
}
