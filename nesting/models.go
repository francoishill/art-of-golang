package main

import (
	"fmt"
	"net/http"
)

var (
	//TODO: obviously no thread-safety here
	users         map[int64]*User         = map[int64]*User{}
	notifications map[int64]*Notification = map[int64]*Notification{}
)

type User struct {
	ID    int64
	Admin bool
}

type Notification struct {
	ID   int64
	Text string
}

func GetUser(id int64, next func(user *User) *Error) *Error {
	user, ok := users[id]
	if !ok {
		return &Error{fmt.Errorf("User not found with ID %d", id), http.StatusBadRequest}
	}
	return next(user)
}

func CreateNotification(text string, next func(notif *Notification) *Error) *Error {
	//TODO: no thread-safety
	id := int64(len(notifications) + 1)
	notif := &Notification{
		ID:   id,
		Text: text,
	}
	notifications[id] = notif
	return next(notif)
}
