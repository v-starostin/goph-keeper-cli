package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func SignUp() *cobra.Command {
	return &cobra.Command{
		Use: "signup",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("SignUp command")
		},
	}
}
