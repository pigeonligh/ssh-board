package main

import (
	"fmt"
	"os"

	"github.com/pigeonligh/ssh-board/pkg/auth"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("error")
		return
	}
	user := auth.Auth{
		User:      os.Args[1],
		PublicKey: os.Args[2],
	}

	if err := auth.ApplyAuth([]auth.Auth{user}, true); err != nil {
		fmt.Printf("failed to apply auth, err: %v\n", err)
		return
	}
}
