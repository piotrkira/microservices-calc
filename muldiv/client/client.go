package client

import (
	"context"
	"fmt"
	"log"

	"github.com/piotrkira/microservices-calc/muldiv/endpoints"

	"google.golang.org/grpc"
)

type Client struct {
	cli endpoints.MulDivClient
}

func New(serverAddres string) *Client {
	client := Client{}
	connection, err := grpc.Dial(fmt.Sprintf("%s:7777", serverAddres), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client.cli = endpoints.NewMulDivClient(connection)
	return &client
}

func (c *Client) Mul(a, b int64) (string, error) {
	res, err := c.cli.Mul(context.Background(), &endpoints.Numbers{A: a, B: b})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.Content, nil
}

func (c *Client) Div(a, b int64) (string, error) {
	res, err := c.cli.Div(context.Background(), &endpoints.Numbers{A: a, B: b})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.Content, nil
}
