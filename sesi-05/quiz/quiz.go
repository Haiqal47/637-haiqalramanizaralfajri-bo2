package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	defer catchErr()

	var username string
	fmt.Print("Enter Username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	// fmt.Scanln(&password)
	// password, err := terminal.ReadPassword(0)
	pass, _ := gopass.GetPasswdMasked()

	_, err := validAccount(username, string(pass))

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("username dan password benar")
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validAccount(username string, password string) (string, error) {
	ul := len(username)
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}
	if ul < 5 {
		return "", errors.New("username has to have more than 4 characters")
	}

	return "Valid Account", nil
}
