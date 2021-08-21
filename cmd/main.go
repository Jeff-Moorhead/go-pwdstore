package main

import (
	"fmt"

	"github.com/jeff-moorhead/go-pwdmgr/app"
)

func main() {
	// Must be able to add, remove, get, and modify passwords based on website ID, such as www.facebook.com. ID can be whatever the user wants.
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
