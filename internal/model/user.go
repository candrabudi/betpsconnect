package model

type User struct {
	ID       int    `bson:"id,omitempty"`
	FullName string `bson:"full_name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	Regency  string `bson:"regency,empty"`
	Role     string `bson:"role, omitempty"`
	Status   string `bson:"status, omitempty"`
}
