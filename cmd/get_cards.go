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

type GetCardsOptions struct {
	AccessToken string
}

func GetCards() *cobra.Command {
	opts := new(GetCardsOptions)
	cmd := &cobra.Command{
		Use: "getcards",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("getcards command")
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
			res, err := client.GetCards(ctx, &pb.GetCardsRequest{})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("getcards res:", res.GetCards())
		},
	}

	getCardsFlags(cmd, opts)

	return cmd
}

func getCardsFlags(cmd *cobra.Command, opts *GetCardsOptions) {
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
