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

type AddNoteOptions struct {
	Title       string
	Content     string
	AccessToken string
}

func AddNote() *cobra.Command {
	opts := new(AddNoteOptions)
	cmd := &cobra.Command{
		Use: "addnote",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("addnote command")
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
			_, err = client.AddNote(ctx, &pb.AddNoteRequest{
				Title:   opts.Title,
				Content: opts.Content,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("addnote res: success")
		},
	}

	addNoteFlags(cmd, opts)

	return cmd
}

func addNoteFlags(cmd *cobra.Command, opts *AddNoteOptions) {
	cmd.Flags().StringVarP(&opts.Title, "title", "t", "", "")
	cmd.Flags().StringVarP(&opts.Content, "content", "c", "", "")
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
