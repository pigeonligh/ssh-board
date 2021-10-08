package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/pigeonligh/ssh-board/pkg/auth"
)

type Config struct {
	Auths []auth.Auth
}

func initViper() {
	_ = os.MkdirAll("/etc/sshboard/", os.ModePerm)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/sshboard/")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.SafeWriteConfig()
		}
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}
}

func saveViper() {
	err := viper.WriteConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func getAuthFromViper() []auth.Auth {
	config := &Config{}
	err := viper.Unmarshal(config)
	if err != nil {
		return []auth.Auth{}
	}
	return config.Auths
}

func set() {
	auths := getAuthFromViper()
	if err := auth.ApplyAuth(auths); err != nil {
		fmt.Printf("failed to apply auth, err: %v\n", err)
		return
	}
}

func add(user auth.Auth) {
	auths := getAuthFromViper()

	for _, auth := range auths {
		if auth.PublicKey == user.PublicKey {
			fmt.Println("public key added")
			return
		}
	}
	auths = append(auths, user)

	if err := auth.ApplyAuth(auths); err != nil {
		fmt.Printf("failed to apply auth, err: %v\n", err)
		return
	}
	viper.Set("auths", auths)
	saveViper()
}

func main() {
	initViper()

	if len(os.Args) == 1 {
		set()
		return
	}
	if len(os.Args) == 3 {
		add(auth.Auth{
			User:      os.Args[1],
			PublicKey: os.Args[2],
		})
		return
	}
	fmt.Println("error")
}
