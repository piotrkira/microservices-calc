package client

import (
	"context"
	"fmt"
	"log"

	"github.com/piotrkira/microservices-calc/addsub/endpoints"

	"google.golang.org/grpc"
)

type Client struct {
	cli endpoints.AddSubClient
}

func New(serverAddres string) *Client {
	client := Client{}
	connection, err := grpc.Dial(fmt.Sprintf("%s:7777", serverAddres), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client.cli = endpoints.NewAddSubClient(connection)
	return &client
}

func (c *Client) Add(a, b int64) (string, error) {
	res, err := c.cli.Add(context.Background(), &endpoints.Numbers{A: a, B: b})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.Content, nil
}

func (c *Client) Sub(a, b int64) (string, error) {
	res, err := c.cli.Sub(context.Background(), &endpoints.Numbers{A: a, B: b})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.Content, nil
}
