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

type GetCredsOptions struct {
	ResourceName, AccessToken string
}

func GetCreds() *cobra.Command {
	opts := new(GetCredsOptions)
	cmd := &cobra.Command{
		Use: "getcreds",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("getcreds command")
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
			res, err := client.GetCredentials(ctx, &pb.GetCredentialsRequest{
				ResourceName: opts.ResourceName,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("getcreds res: username,", res.GetUsername())
			fmt.Println("getcreds res: passwprd,", res.GetPassword())
		},
	}

	GetCredsFlags(cmd, opts)

	return cmd
}

func GetCredsFlags(cmd *cobra.Command, opts *GetCredsOptions) {
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
	cmd.Flags().StringVarP(&opts.ResourceName, "resource", "r", "", "")
}
