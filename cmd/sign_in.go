package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func SignIn() *cobra.Command {
	return &cobra.Command{
		Use: "signin",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("SignIn command")
		},
	}
}
