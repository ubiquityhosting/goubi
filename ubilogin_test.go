package goubi

import (
	"strconv"
	"testing"
)

func TestCredentials(t *testing.T) {
	id := 123
	user := "ubic-123"
	token := "asd65f4asd65f46s5adf4"
	setCredentials(id, user, token)
	username, password := getCredentials()
	if username != strconv.Itoa(id) {
		t.Error("Username did not convert")
	}
	if password != user+":"+token {
		t.Error("Password format error")
	}
}
