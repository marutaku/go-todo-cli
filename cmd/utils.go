package cmd

import (
	"fmt"
	"os"
)

func ExitWithError(message error) {
	fmt.Println(message)
	os.Exit(1)
}
