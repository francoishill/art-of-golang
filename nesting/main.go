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
		body     struct {
			Text string `json:"text,omitempty"`
		}
		notification *Notification
	)
	err := BindRequestInt64Param(c, "user_id", func(userID int64) *Error {
		return BindRequestBody(c, body, func() *Error {
			return GetUser(userID, func(user *User) *Error {
				return AuthorizeSendNotification(authUser, user, func() *Error {
					return CreateNotification(body.Text, func(notif *Notification) *Error {
						notification = notif
						return nil
					})
				})
			})
		})
	})
	if err != nil {
		return c.JSON(err.Status, struct{ Error string }{err.Error.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Notification *Notification
	}{
		Notification: notification,
	})
}
