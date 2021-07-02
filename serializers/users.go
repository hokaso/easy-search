package serializers

import "easy-search/models"

type User struct {
	ID    uint   `json:"id"`
	Qq    string `json:"qq"`
	Phone string `json:"phone"`
}

type UserPhoneRequest struct {
	Phone string `form:"phone" binding:"required"`
}

type UserQqRequest struct {
	Qq string `form:"qq" binding:"required"`
}

type UserResponse struct {
	User User `json:"user"`
}

func SerializeUser(user models.SenQq) User {
	return User{
		ID:    user.ID,
		Qq:    user.Qq,
		Phone: user.Phone,
	}
}
