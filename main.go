package main

import (
	"fmt"
	"os"
)

// So stdin can be mocked during testing.
var stdin *os.File

func main() {
	parseFlags(nil, nil)

	email, err := settings()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = sendMail(email)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
