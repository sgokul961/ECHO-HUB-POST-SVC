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

type ChatServiceClient struct {
	Client pb.ChatServiceClient
}

func NewChatServiceClient(c config.Config) clientinterface.ChatServiceClient {

	//if call fails chech for error here grpc communication

	cc, err := grpc.Dial(c.ChatHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("coudnt connect:", err)
	}
	return &ChatServiceClient{Client: pb.NewChatServiceClient(cc)}
}

// CreateChatRoom implements clientinterface.ChatServiceClient.
func (c *ChatServiceClient) CreateChatRoom(user1 int64, user2 int64) error {
	req := &pb.ChatRoomRequest{
		FollowingId: user1,
		FOllowerId:  user2,
	}
	_, err := c.Client.CreateChatRoom(context.Background(), req)
	if err != nil {
		return errors.New("error in getting auth service method")
	}
	return nil
}
