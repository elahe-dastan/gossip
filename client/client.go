package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elahe-dastan/gossip/protocol"

	"google.golang.org/grpc"
)

func New(addr string) (protocol.ChatClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return protocol.NewChatClient(conn), nil
}

func Show(cli protocol.ChatClient, id *protocol.ID) {
	const connPeriod = 5 * time.Second

	for {
		res, err := cli.Receive(context.Background(), id)
		if err != nil {
			log.Println(err)
		}

		for {
			m, err := res.Recv()
			if err != nil {
				break
			}

			fmt.Printf("id:%d > %s", m.Id.Id, m.Text)
		}

		<-time.NewTicker(connPeriod).C
	}
}
