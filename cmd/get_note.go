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

type GetNoteOptions struct {
	AccessToken string
	Title       string
}

func GetNote() *cobra.Command {
	opts := new(GetNoteOptions)
	cmd := &cobra.Command{
		Use: "getnote",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("getnote command")
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
			res, err := client.GetNote(ctx, &pb.GetNoteRequest{
				Title: opts.Title,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("getnote title:", res.GetTitle())
			fmt.Println("getnote content:", res.GetContent())
		},
	}

	getNoteFlags(cmd, opts)

	return cmd
}

func getNoteFlags(cmd *cobra.Command, opts *GetNoteOptions) {
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
	cmd.Flags().StringVarP(&opts.Title, "title", "t", "", "")
}
