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

type GetFilesOptions struct {
	AccessToken string
}

func GetFiles() *cobra.Command {
	opts := new(GetFilesOptions)
	cmd := &cobra.Command{
		Use: "getfiles",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("getfiles command")
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

			files, err := client.GetFiles(ctx, &pb.GetFilesRequest{})
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("getfiles res:", files.GetFiles())
		},
	}

	getFiles(cmd, opts)

	return cmd
}

func getFiles(cmd *cobra.Command, opts *GetFilesOptions) {
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
