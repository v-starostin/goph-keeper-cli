package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/v-starostin/goph-keeper/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type GetFileOptions struct {
	AccessToken string
	File        string
}

func GetFile() *cobra.Command {
	opts := new(GetFileOptions)
	cmd := &cobra.Command{
		Use: "getfile",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("getfile command")
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
			md.Set("filename", opts.File)
			ctx := metadata.NewOutgoingContext(context.Background(), md)

			stream, err := client.GetFile(ctx, &pb.GetFileRequest{Filename: opts.File})
			if err != nil {
				fmt.Println(err)
				return
			}

			file, err := os.Create(opts.File)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					fmt.Println("end of file")
					break
				}
				if err != nil {
					fmt.Println(err)
					return
				}
				_, err = file.Write(resp.GetContent())
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			fmt.Println("getfile res: success")
		},
	}

	getFile(cmd, opts)

	return cmd
}

func getFile(cmd *cobra.Command, opts *GetFileOptions) {
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "")
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
