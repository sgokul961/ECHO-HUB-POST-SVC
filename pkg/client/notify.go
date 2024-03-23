package client

import (
	"fmt"

	clientinterface "github.com/sgokul961/echo-hub-post-svc/pkg/client/clientInterface"
	"github.com/sgokul961/echo-hub-post-svc/pkg/config"
	"github.com/sgokul961/echo-hub-post-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NotifyServiceClient struct {
	Client pb.NotificationServiceClient
}

func NewNotifyServiceClient(c config.Config) clientinterface.NotifyServiceClient {
	cc, err := grpc.Dial(c.AuthHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("coudnt connect:", err)
	}

	return &NotifyServiceClient{Client: pb.NewNotificationServiceClient(cc)}
}
