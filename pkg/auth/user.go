package auth

import "os"

func GetUser() (string, bool) {
	username := os.Getenv("REMOTEUSER")
	if username == "noname" || username == "" {
		return "", false
	}
	return username, true
}
