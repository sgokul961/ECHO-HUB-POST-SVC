package client

import (
	"context"
	"errors"
	"fmt"

	clientinterface "github.com/sgokul961/echo-hub-post-svc/pkg/client/clientInterface"
	"github.com/sgokul961/echo-hub-post-svc/pkg/config"
	"github.com/sgokul961/echo-hub-post-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func NewAuthServiceClient(c config.Config) clientinterface.AuthServiceClient {

	//if call fails chech for error here grpc communication

	cc, err := grpc.Dial(c.AuthHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("coudnt connect:", err)
	}
	return &AuthServiceClient{Client: pb.NewAuthServiceClient(cc)}
}

func (a *AuthServiceClient) CheckUserBlocked(id int64) (bool, error) {
	req := &pb.CheckUserBlockedRequest{
		Id: id,
	}
	res, err := a.Client.CheckUserBlocked(context.Background(), req)
	if err != nil {
		return false, errors.New("error in getting auth service method")
	}
	return res.IsBlock, nil
}
