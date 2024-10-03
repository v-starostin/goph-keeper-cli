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

type GetCardOptions struct {
	Bank, AccessToken string
}

func GetCard() *cobra.Command {
	opts := new(GetCardOptions)
	cmd := &cobra.Command{
		Use: "getcard",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("getcard command")
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
			res, err := client.GetCard(ctx, &pb.GetCardRequest{
				Bank: opts.Bank,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("getcard res:", res.GetCards())
		},
	}

	getCardFlags(cmd, opts)

	return cmd
}

func getCardFlags(cmd *cobra.Command, opts *GetCardOptions) {
	cmd.Flags().StringVarP(&opts.Bank, "bank", "b", "", "")
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
