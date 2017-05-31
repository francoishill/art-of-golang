package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func BindRequestInt64Param(c echo.Context, paramName string, next func(val int64) *Error) *Error {
	strVal := strings.TrimSpace(c.Param(paramName))
	if len(strVal) == 0 {
		return &Error{fmt.Errorf("Param %s is required", paramName), http.StatusBadRequest}
	}

	val, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		return &Error{fmt.Errorf("Cannot parse %s as int", strVal), http.StatusBadRequest}
	}

	return next(val)
}

func BindRequestBody(c echo.Context, dest interface{}, next func() *Error) *Error {
	if err := c.Bind(dest); err != nil {
		return &Error{err, http.StatusBadRequest}
	}
	return next()
}
