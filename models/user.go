package models

type User struct {
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	Address  string `bson:"address,omitempty"`
	Gender   string `bson:"gender,omitempty"`
}
