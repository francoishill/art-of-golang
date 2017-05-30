package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/user/:user_id/notifications", NotifyUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func NotifyUser(c echo.Context) error {
	var (
		authUser = c.Get("user").(*User)
		userID   int64
		body     struct {
			Text string `json:"text,omitempty"`
		}
		user         User
		notification Notification
	)
	chainErr := ChainLoad(
		BindRequestInt64Param(c, "user_id", &userID),
		BindRequestBody(c, &body),
		GetUser(&userID, &user),
		AuthorizeSendNotification(authUser, &user),
		CreateNotification(&body.Text, &notification),
	)
	if chainErr != nil {
		return c.JSON(chainErr.Status, struct{ Error string }{chainErr.Error.Error()})
	}
	return c.JSON(http.StatusOK, struct {
		Notification *Notification
	}{
		Notification: &notification,
	})
}
