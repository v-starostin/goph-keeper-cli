package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/v-starostin/goph-keeper/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type AddCredsOptions struct {
	Username, Password, ResourceName, Description, AccessToken string
}

func AddCreds() *cobra.Command {
	opts := new(AddCredsOptions)
	cmd := &cobra.Command{
		Use: "addcreds",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("addcreds command")
			fmt.Printf("opts: %+v\n", opts)

			conn, err := grpc.NewClient(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()

			client := pb.NewGophkeeperClient(conn)
			m := make(map[string]string)
			md := metadata.New(m)
			md.Set("authorization", "bearer "+opts.AccessToken)
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			_, err = client.AddCredentials(ctx, &pb.AddCredentialsRequest{
				Username:     opts.Username,
				Password:     opts.Password,
				Description:  opts.Description,
				ResourceName: opts.ResourceName,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("addnote res: success")
		},
	}

	addcredsFlags(cmd, opts)

	return cmd
}

func addcredsFlags(cmd *cobra.Command, opts *AddCredsOptions) {
	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "")
	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "")
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
	cmd.Flags().StringVarP(&opts.Description, "description", "d", "", "")
	cmd.Flags().StringVarP(&opts.ResourceName, "resource", "r", "", "")
}
