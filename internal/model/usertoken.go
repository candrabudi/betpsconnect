package model

type UserToken struct {
	ID     int    `bson:"id,omitempty"`
	UserID int    `bson:"user_id,omitempty"`
	Token  string `bson:"token,omitempty"`
}
