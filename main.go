package main

import (
	"fmt"
	"os"
)

func main() {
	setupInputs(nil, nil)

	email, err := settings()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = sendMail(email)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
