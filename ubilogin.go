package goubi

import (
	"strconv"
)

/*****************************************************************************/

var creds credentials

// Represents the login credentials
type credentials struct {
	Username string
	Password string
}

/*****************************************************************************/

// Sets the login credentials
func setCredentials(clientid int, username string, token string) {
	creds.Username = strconv.Itoa(clientid)
	creds.Password = username + ":" + token
}

// Returns the login credentials
func getCredentials() (string, string) {
	return creds.Username, creds.Password
}
