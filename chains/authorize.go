package main

import (
	"errors"
	"net/http"
)

func AuthorizeSendNotification(from, to *User) Loader {
	return func() *Error {
		//TODO: Add own Authorize logic here
		if from.Admin {
			return &Error{errors.New("Only Admin users can send notifications"), http.StatusUnauthorized}
		}
		return nil
	}
}
