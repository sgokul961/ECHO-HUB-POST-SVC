package handler

import (
	"context"
	"errors"
	"fmt"
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

func (u *PostHandler) DislikePost(ctx context.Context, dislike *pb.DislikePostRequest) (*pb.DislikePostResponse, error) {
	res, err := u.postusecase.DisLikepost(dislike.UserId, dislike.PostId)

	if err != nil {
		return &pb.DislikePostResponse{
			Success: res,
		}, err
	}
	return &pb.DislikePostResponse{
		Success: res,
	}, nil

}
func (u *PostHandler) CommentPost(ctx context.Context, comment *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {

	response, err := u.postusecase.AddComment(models.AddComent{
		PostsID: comment.PostId,
		UserID:  comment.UserId,
		Content: comment.Content,
	})
	fmt.Println("content", comment.Content)
	if err != nil {
		// Handle error
		return nil, err
	}
	return &pb.CommentPostResponse{
		CommentId: response,
	}, nil

}
func (u *PostHandler) GetComments(ctx context.Context, getcomment *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {

	comments, err := u.postusecase.GetComment(getcomment.PostId)

	if err != nil {
		return nil, err
	}
	// Convert []string comments to []*pb.Comment
	pbComments := make([]*pb.Comment, len(comments))
	for i, comment := range comments {
		pbComments[i] = &pb.Comment{Content: comment} // Assuming pb.Comment has a Content field to store comment content
	}

	// Create response with converted comments
	response := &pb.GetCommentsResponse{
		Comments: pbComments,
	}

	return response, nil // Return the response along with nil error if no error occurred

}
func (u *PostHandler) DeleteComments(ctx context.Context, deleteComment *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	// Call the DeleteComment method from the usecase

	deletedCommentID, success := u.postusecase.DeleteComment(deleteComment.PostId, deleteComment.CommentId, deleteComment.UserId)
	fmt.Println("success ", success)
	if !success {
		// Return error if deletion fails

		return nil, errors.New("cant delete the comment,invalid comment id")
	}
	// Construct and return the response message

	response := &pb.DeleteCommentResponse{
		CommentId: deletedCommentID,
		Success:   true,
	}
	return response, nil

}
func (u *PostHandler) GetUserId(ctx context.Context, p *pb.GetUserIdRequest) (*pb.GetUserIdResponse, error) {

	user_id, err := u.postusecase.GetPostOwner(p.PostId)

	if err != nil {
		return nil, err
	}
	response := &pb.GetUserIdResponse{
		UserId:  user_id,
		Success: true,
	}
	return response, nil

}
