package models

type User struct {
	ID        string `bson:"_id,omitempty"`
	Name      string `bson:"name,omitempty"`
	Email     string `bson:"email,omitempty"`
	Password  string `bson:"password,omitempty"`
	Address   string `bson:"address,omitempty"`
	Gender    string `bson:"gender,omitempty"`
	CreatedAt string `bson:"created_at,omitempty"`
}

type UserCreateRequest struct {
	Name     string `json:"name" form:"name" binding:"required,max=50"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
	Address  string `json:"address" form:"address" binding:"required"`
	Gender   string `json:"gender" form:"gender" binding:"required"`
}
