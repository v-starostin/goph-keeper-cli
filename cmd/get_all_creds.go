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

type GetAllCredsOptions struct {
	AccessToken string
}

func GetAllCreds() *cobra.Command {
	opts := new(GetAllCredsOptions)
	cmd := &cobra.Command{
		Use: "getallcreds",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("getallcreds command")
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
			res, err := client.GetAllCredentials(ctx, &pb.GetAllCredentialsRequest{})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("getallcreds res: credentials,", res.GetCredentials())
		},
	}

	getAllCredsFlags(cmd, opts)

	return cmd
}

func getAllCredsFlags(cmd *cobra.Command, opts *GetAllCredsOptions) {
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
