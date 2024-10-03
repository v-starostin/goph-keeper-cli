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

type AddFileOptions struct {
	AccessToken string
	File        string
}

func AddFile() *cobra.Command {
	opts := new(AddFileOptions)
	cmd := &cobra.Command{
		Use: "addfile",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("addfile command")
			fmt.Printf("opts: %+v\n", opts)

			conn, err := grpc.NewClient(
				":9090",
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()

			file, err := os.Open(opts.File)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			client := pb.NewGophkeeperClient(conn)
			m := make(map[string]string)
			md := metadata.New(m)
			md.Set("authorization", "bearer "+opts.AccessToken)
			md.Set("filename", opts.File)
			ctx := metadata.NewOutgoingContext(context.Background(), md)

			stream, err := client.AddFile(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}

			buf := make([]byte, 5*1024*1024)

			for {
				n, err := file.Read(buf)
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println(err)
					return
				}
				err = stream.Send(&pb.AddFileRequest{
					Filename: opts.File,
					Content:  buf[:n],
				})
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			_, err = stream.CloseAndRecv()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("addfile res: success")
		},
	}

	addFile(cmd, opts)

	return cmd
}

func addFile(cmd *cobra.Command, opts *AddFileOptions) {
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "")
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
