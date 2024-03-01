package usecase

import (
	"errors"
	"fmt"

	clientinterface "github.com/sgokul961/echo-hub-post-svc/pkg/client/clientInterface"
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
