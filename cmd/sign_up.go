package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/v-starostin/goph-keeper/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SignUpOptions struct {
	Username string
	Password string
}

func SignUp() *cobra.Command {
	opts := new(SignUpOptions)
	cmd := &cobra.Command{
		Use: "signup",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println("SignUp command")
			fmt.Printf("opts: %+v\n", opts)

			conn, err := grpc.NewClient(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()

			client := pb.NewAuthClient(conn)

			res, err := client.Register(context.Background(), &pb.RegisterRequest{
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

	addSignUpFlags(cmd, opts)

	return cmd
}

func addSignUpFlags(cmd *cobra.Command, opts *SignUpOptions) {
	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "")
	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "")
}
