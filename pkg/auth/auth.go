package auth

import (
	"fmt"
	"os"
	"strings"
)

type Auth struct {
	User      string
	PublicKey string
}

func ApplyAuth(auths []Auth) error {
	flag := os.O_WRONLY | os.O_CREATE
	file, err := os.OpenFile(AuthPath, flag, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	lines := make([]string, 0)
	for _, auth := range auths {
		line := fmt.Sprintf("environment=\"REMOTEUSER=%s\" %s", auth.User, auth.PublicKey)
		lines = append(lines, line)
	}
	lines = append(lines, "")
	_, err = file.WriteString(strings.Join(lines, "\n"))
	return err
}
