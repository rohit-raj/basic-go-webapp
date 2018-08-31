package model

const (
	CollectionUsers = "users"
)

// User model
type User struct {
	Id        	string
	Username    string
	Password    string
	CreatedOn 	int64
	UpdatedOn 	int64
}
