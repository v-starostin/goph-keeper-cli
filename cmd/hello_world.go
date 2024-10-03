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

type HelloWorldOptions struct {
	Username    string
	AccessToken string
}

func HelloWorld() *cobra.Command {
	opts := new(HelloWorldOptions)
	cmd := &cobra.Command{
		Use: "helloworld",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("helloworld command")
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
			res, err := client.HelloWorld(ctx, &pb.HelloWorldRequest{
				Username: opts.Username,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("res", res.GetResponse())
		},
	}

	addHelloWorldFlags(cmd, opts)

	return cmd
}

func addHelloWorldFlags(cmd *cobra.Command, opts *HelloWorldOptions) {
	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "")
	cmd.Flags().StringVarP(&opts.AccessToken, "access", "a", "", "")
}
