package clientinterface

type ChatServiceClient interface {
	CreateChatRoom(user1, user2 int64) error
}
