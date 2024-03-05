package handler

import (
	"context"
	"strings"

	"github.com/sgokul961/echo-hub-post-svc/pkg/models"
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
func (u *PostHandler) UnfollowUser(ctx context.Context, unfollow *pb.UnfollowUserRequest) (*pb.UnfollowUserResponse, error) {

	response, err := u.postusecase.UnfollowUser(unfollow.FollowUserId, unfollow.FollowerUserId)
	if err != nil {
		return &pb.UnfollowUserResponse{
			FollowerUserId: unfollow.FollowerUserId,
		}, err
	}
	return &pb.UnfollowUserResponse{
		FollowerUserId: response,
	}, nil

}
func (u *PostHandler) UploadPost(ctx context.Context, upload *pb.UploadPostRequest) (*pb.UploadPostResponse, error) {
	mediaURLs := make([]string, len(upload.MediaUrls))
	copy(mediaURLs, upload.MediaUrls)

	response, err := u.postusecase.AddPost(models.AddPost{
		UserID:   upload.UserId,
		Content:  upload.Content,
		ImageURL: strings.Join(mediaURLs, ""),
	})
	if err != nil {
		// Handle error
		return nil, err
	}
	return &pb.UploadPostResponse{
		UserId: response,
	}, nil
}
func (u *PostHandler) DeletePost(ctx context.Context, delete *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {

	res, err := u.postusecase.DeletePost(delete.PostId, delete.UserId)
	if err != nil {
		return &pb.DeletePostResponse{
			Success: false,
			PostId:  res,
		}, err
	}
	return &pb.DeletePostResponse{
		Success: true,
		PostId:  delete.PostId,
	}, nil

}
func (u *PostHandler) LikePost(ctx context.Context, like *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	res, err := u.postusecase.LikePost(like.PostId, like.UserId)
	if err != nil {
		return &pb.LikePostResponse{
			PostId: like.PostId,
		}, err
	}
	return &pb.LikePostResponse{
		PostId: res,
	}, nil

}
