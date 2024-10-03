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

type AddCardOptions struct {
	CardNumber, Bank, AccessToken string
}

func AddCard() *cobra.Command {
	opts := new(AddCardOptions)
	cmd := &cobra.Command{
		Use: "addcard",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("addcard command")
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
			_, err = client.AddCard(ctx, &pb.AddCardRequest{
				CardNumber: opts.CardNumber,
				Bank:       opts.Bank,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("addnote res: success")
		},
	}

	addCardFlags(cmd, opts)

	return cmd
}

func addCardFlags(cmd *cobra.Command, opts *AddCardOptions) {
	cmd.Flags().StringVarP(&opts.Bank, "bank", "b", "", "")
	cmd.Flags().StringVarP(&opts.CardNumber, "cardnumber", "c", "", "")
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
