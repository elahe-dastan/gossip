package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elahe-dastan/gossip/client"
	"github.com/elahe-dastan/gossip/protocol"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"
)

func Register(rootCmd *cobra.Command) {
	c := &cobra.Command{
		Use:   "client",
		Short: "Runs client",
		Run: func(cmd *cobra.Command, args []string) {
			addr, err := cmd.Flags().GetString("server")
			if err != nil {
				log.Fatal(err)
			}

			cli, err := client.New(addr)
			if err != nil {
				log.Fatal(err)
			}

			var id *protocol.ID

			setID, err := cmd.Flags().GetInt32("ID")
			if err != nil {
				log.Fatal(err)
			}

			if setID == -1 {
				id, err = cli.Who(context.Background(), &empty.Empty{})
				if err != nil {
					log.Fatal(err)
				}
			} else {
				id = &protocol.ID{
					Id: setID,
				}
			}
			fmt.Println(id.Id)

			reader := bufio.NewReader(os.Stdin)

			go client.Show(cli, id)

			for {
				fmt.Print(" >> ")
				text, err := reader.ReadString('\n')
				if err != nil {
					log.Println(err)
				}

				if _, err := cli.Send(context.Background(), &protocol.Data{
					Id:   id,
					Text: text,
				}); err != nil {
					log.Println(err)
				}
			}
		},
	}

	c.Flags().StringP("server", "s", "127.0.0.1:1373", "server address")
	c.Flags().Int32P("ID", "i", -1, "client id")

	rootCmd.AddCommand(c)
}
