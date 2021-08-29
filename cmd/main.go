package main

import (
	"fmt"

	"github.com/jeff-moorhead/go-pwdstore/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		fmt.Printf("Unable to run password manager: %v\n", err)
		return
	}

	err = app.Run()
	if err != nil {
		fmt.Printf("Unable to run password manager: %v\n", err)
		return
	}
}
