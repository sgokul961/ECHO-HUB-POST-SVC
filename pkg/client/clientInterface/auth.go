package clientinterface

type AuthServiceClient interface {
	CheckUserBlocked(id int64) (bool, error)
}
