package main

import (
	"fmt"
	"os"

	"github.com/fuji97/oauth2c/cmd"
)

func main() {
	if err := cmd.NewOAuth2Cmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
