package Models

type User struct {
	FirstName string `json:"firstName" bson:"first_name" binding:"required"`
	LastName  string `json:"lastName" bson:"last_name" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required"`
}
