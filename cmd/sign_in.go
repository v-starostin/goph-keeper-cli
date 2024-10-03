package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/v-starostin/goph-keeper/pkg/pb"
)

type SignInOptions struct {
	Username string
	Password string
}

func SignIn() *cobra.Command {
	opts := new(SignInOptions)
	cmd := &cobra.Command{
		Use: "signin",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("SignIn command")
			fmt.Printf("opts: %+v\n", opts)

			conn, err := grpc.NewClient(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()

			client := pb.NewAuthClient(conn)

			res, err := client.Login(context.Background(), &pb.LoginRequest{
				Username: opts.Username,
				Password: opts.Password,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("res.GetAccessToken", res.GetAccessToken())
			fmt.Println("res.GetRefreshToken", res.GetRefreshToken())
		},
	}

	addSignInFlags(cmd, opts)

	return cmd
}

func addSignInFlags(cmd *cobra.Command, opts *SignInOptions) {
	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "")
	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "")
}
