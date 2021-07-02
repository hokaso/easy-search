package serializers

import "easy-search/models"

type User struct {
	ID    uint   `json:"id"`
	Qq    string `json:"qq"`
	Phone string `json:"phone"`
}

type UserPhoneRequest struct {
	Phone string `json:"phone" binding:"required"`
}

type UserQqRequest struct {
	Qq string `json:"qq" binding:"required"`
}

type UserResponse struct {
	User User `json:"user"`
}

func SerializeUser(user models.User) User {
	return User{
		ID:    user.ID,
		Qq:    user.Qq,
		Phone: user.Phone,
	}
}
