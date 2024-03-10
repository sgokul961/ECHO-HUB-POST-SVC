package usecase

import (
	"errors"
	"fmt"

	clientinterface "github.com/sgokul961/echo-hub-post-svc/pkg/client/clientInterface"
	"github.com/sgokul961/echo-hub-post-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-post-svc/pkg/models"
	interfacesR "github.com/sgokul961/echo-hub-post-svc/pkg/repository/interface"
	interfacesU "github.com/sgokul961/echo-hub-post-svc/pkg/usecase/usecaseinterface"
)

type postUsecase struct {
	postRepo   interfacesR.PostRepoInterface
	AuthClient clientinterface.AuthServiceClient
}

func NewPostUseCase(postrepo interfacesR.PostRepoInterface, authClient clientinterface.AuthServiceClient) interfacesU.PostUseCaseInterface {
	return &postUsecase{
		postRepo:   postrepo,
		AuthClient: authClient,
	}

}
func (p *postUsecase) FollowUser(following_id, follower_id int64) (bool, error) {

	if following_id == follower_id {
		return false, errors.New("cant follow your account")
	}
	isBlock, err := p.AuthClient.CheckUserBlocked(follower_id)
	if err != nil {
		return false, err
	}
	if !isBlock {
		return false, errors.New("cant follow this user as there is no such user")
	}
	checkifrelationExist := p.postRepo.FollowRelationshipExists(following_id, follower_id)

	if checkifrelationExist {
		return false, errors.New("this relation alredy exist")
	}

	create, err := p.postRepo.CreateFollowing(following_id, follower_id)

	if err != nil {
		return false, err
	}
	fmt.Println("follower", follower_id)
	return create, nil

}
func (p *postUsecase) UnfollowUser(following_id, follower_id int64) (int64, error) {

	chekForRelation := p.postRepo.FollowRelationshipExists(following_id, follower_id)

	if !chekForRelation {
		return 0, errors.New("no relation  exist with this id ")
	}
	err := p.postRepo.Unfollow(following_id, follower_id)

	fmt.Println("follow ,followoer", following_id, follower_id)

	if err != nil {
		return 0, err
	}
	return follower_id, nil

}
func (p *postUsecase) AddPost(upload models.AddPost) (int64, error) {

	post_id, err := p.postRepo.AddPost(domain.Post{
		UserID:    upload.UserID,
		Content:   upload.Content,
		ImageURL:  upload.ImageURL,
		Timestamp: upload.Timestamp.Local(),
	})
	if err != nil {
		return 0, errors.New("databse error ,cant add the post")
	}
	return post_id, nil

}
func (p *postUsecase) DeletePost(post_id, user_id int64) (int64, error) {

	exist := p.postRepo.PostIdExist(post_id, user_id)
	if !exist {
		return 0, errors.New("post does not exist for the given user ,authorization failed")
	}

	deletedPostid, err := p.postRepo.DeletePost(post_id, user_id)

	if err != nil {
		return 0, errors.New("failed to delete post from the database")
	}
	return deletedPostid, nil

}
func (u *postUsecase) LikePost(post_id, user_id int64) (int64, error) {
	//check if the post id exist

	exist := u.postRepo.CheckForPostId(post_id)
	if !exist {
		return 0, errors.New(" invalid post id  ")
	}
	//check if alredy liked

	alreadyLiked := u.postRepo.AlredyLiked(post_id, user_id)
	if alreadyLiked {
		return 0, errors.New("alredy liked ")
	}

	//like post

	like, err := u.postRepo.LikePost(user_id, post_id)

	if err != nil {
		return 0, err
	}

	fmt.Println("post id", post_id)
	updated := u.postRepo.UpdatePost(post_id)
	if !updated {
		return 0, errors.New("unable to update like count")
	}
	return like, nil

}
func (u *postUsecase) DisLikepost(user_id, post_id int64) (bool, error) {

	liked := u.postRepo.ChekIfLikeExist(post_id, user_id)
	if !liked {
		return false, errors.New("cant dislike no relation exist")
	}

	dislike := u.postRepo.DisLikePost(post_id, user_id)

	if !dislike {
		return dislike, errors.New("cant delete this post")
	}
	updated := u.postRepo.UpdatePostDislike(post_id)
	if !updated {
		return false, errors.New("unable to decrement count from post")

	}
	return true, nil

}

//add comment to post

func (u *postUsecase) AddComment(comment models.AddComent) (int64, error) {
	fmt.Println("userid,postid", comment.UserID, comment.PostsID)
	commentid, err := u.postRepo.AddComment(domain.Comment{
		PostsID:   comment.PostsID,
		UserID:    comment.UserID,
		Content:   comment.Content,
		Timestamp: comment.Timestamp,
	})
	if err != nil {
		return 0, errors.New("databse error ,cant add the post")
	}
	return commentid, nil

}
