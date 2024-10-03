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

type GetNotesOptions struct {
	AccessToken string
}

func GetNotes() *cobra.Command {
	opts := new(GetNotesOptions)
	cmd := &cobra.Command{
		Use: "getnotes",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("getnotes command")
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
			res, err := client.GetNotes(ctx, &pb.GetNotesRequest{})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("getnotes res:", res.GetTitles())
		},
	}

	getNotesFlags(cmd, opts)

	return cmd
}

func getNotesFlags(cmd *cobra.Command, opts *GetNotesOptions) {
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
