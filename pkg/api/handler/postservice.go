package handler

import (
	"context"

	"github.com/sgokul961/echo-hub-post-svc/pkg/pb"
	interfacesU "github.com/sgokul961/echo-hub-post-svc/pkg/usecase/usecaseinterface"
)

type PostHandler struct {
	postusecase interfacesU.PostUseCaseInterface

	pb.UnimplementedPostServiceServer
}

func NewPostHandler(postUse interfacesU.PostUseCaseInterface) *PostHandler {
	return &PostHandler{
		postusecase: postUse,
	}
}
func (u *PostHandler) FollowUser(ctx context.Context, follow *pb.FollowUserRequest) (*pb.FollowUserResponse, error) {

	response, err := u.postusecase.FollowUser(follow.FollowUserId, follow.FollowerUserId)
	if err != nil {
		return &pb.FollowUserResponse{
			Success: false,
		}, err
	}
	return &pb.FollowUserResponse{
		Success: response,
	}, nil

}
