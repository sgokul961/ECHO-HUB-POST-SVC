package usecase

import (
	"errors"
	"fmt"
	"log"

	clientinterface "github.com/sgokul961/echo-hub-post-svc/pkg/client/clientInterface"
	"github.com/sgokul961/echo-hub-post-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-post-svc/pkg/helper"
	"github.com/sgokul961/echo-hub-post-svc/pkg/models"
	interfacesR "github.com/sgokul961/echo-hub-post-svc/pkg/repository/interface"
	interfacesU "github.com/sgokul961/echo-hub-post-svc/pkg/usecase/usecaseinterface"
)

type postUsecase struct {
	postRepo   interfacesR.PostRepoInterface
	AuthClient clientinterface.AuthServiceClient
	ChatClient clientinterface.ChatServiceClient
}

func NewPostUseCase(postrepo interfacesR.PostRepoInterface, authClient clientinterface.AuthServiceClient, chatClinet clientinterface.ChatServiceClient) interfacesU.PostUseCaseInterface {
	return &postUsecase{
		postRepo:   postrepo,
		AuthClient: authClient,
		ChatClient: chatClinet,
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
	created := p.ChatClient.CreateChatRoom(follower_id, following_id)

	fmt.Println("chat room creation :", follower_id, following_id)
	if created != nil {
		// Handle error
		return false, errors.New("error creating chat room")
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
	//for notification
	// fetch array of userIds of the followers
	//loop the array and make notification part for each time with the topic as that userid
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

	// creating topic for kafka....the topic will be the one who is going to get the likes
	// postUserId, err := u.postRepo.FetchPostedUserId(post_id)

	// if err != nil {
	// 	return 0, err
	// }

	topic := "like_notification"

	//converting the   user id to string and assigning it into a topic.

	// topic := strconv.FormatInt(postUserId, 10)

	notificationMsg := fmt.Sprintf("New like on post %d by user %d", post_id, user_id)
	fmt.Println("post id is", post_id)
	err = helper.PushLikeNotificationToQueue(models.LikeNotification{
		Topic:   topic,
		UserID:  user_id,
		PostsID: post_id,
	}, []byte(notificationMsg))

	if err != nil {
		// Handle error (e.g., log it)
		log.Printf("Failed to push comment notification to Kafka: %v", err)
		// Continue processing even if pushing notification fails
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

	commentid, err := u.postRepo.AddComment(domain.Comment{
		PostsID:   comment.PostsID,
		UserID:    comment.UserID,
		Content:   comment.Content,
		Timestamp: comment.Timestamp,
	})
	if err != nil {
		return 0, errors.New("databse error ,cant add the post")
	}
	update := u.postRepo.UpdateCommentCount(comment.PostsID)
	if !update {
		return 0, errors.New("cant update comment count")

	}
	topic := "comment_notifications"
	//creating an event in kafka to trigger the pushcomment to

	// Create a notification message for the new comment
	notificationMsg := fmt.Sprintf("New comment on post %d by user %d: %s", comment.PostsID, comment.UserID, comment.Content)
	// fetch userid of the post and make it as the topic
	// Push the notification message to Kafka
	// err = helper.PushCommentToQueue("comment_notifications", []byte(notificationMsg))
	// if err != nil {
	// 	// Handle error (e.g., log it)
	// 	log.Printf("Failed to push comment notification to Kafka: %v", err)
	// 	// Continue processing even if pushing notification fails
	// }

	//changed    check

	err = helper.PushcommentNotificationToQueue(models.CommentNotification{

		UserID:  comment.UserID,
		Message: comment.Content,
		PostID:  comment.PostsID,
		Topic:   topic,
		Content: comment.Content,
	}, []byte(notificationMsg))

	if err != nil {
		log.Printf("Failed to push comment notification to Kafka: %v", err)
	}

	// Return the ID of the added comment
	return commentid, nil

}
func (u *postUsecase) GetComment(post_id int64) ([]string, error) {

	idexist := u.postRepo.CheckForPostId(post_id)

	if !idexist {
		return nil, errors.New("post id dosnt exist")
	}

	comments, err := u.postRepo.GetComment(post_id)
	if err != nil {
		return nil, err
	}

	// Now you have the comments, you can use them as needed
	fmt.Println("Comments:", comments)

	return comments, nil

}

//delete comment

func (u *postUsecase) DeleteComment(postID, commentID, UserID int64) (int64, bool) {

	exists := u.postRepo.ChcekCommentExist(postID, commentID, UserID)
	if !exists {
		return 0, false
	}

	commentId, success := u.postRepo.DeleteComment(postID, commentID, UserID)

	if !success {
		return 0, false
	}

	update := u.postRepo.UpdateCommentCountAfterDelete(postID)
	if !update {
		return 0, false
	}
	return commentId, true
}

//getUserIdOfPost  :: for validating kafka

func (p *postUsecase) GetPostOwner(postid int64) (int64, error) {

	user_id, err := p.postRepo.FetchPostedUserId(postid)

	if err != nil {
		return 0, err
	}
	return user_id, nil

}
