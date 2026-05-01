package users

import "github.com/fahrurben/geteventgo/common"

type UserSerializer struct {
	Model *UserModel
}

type UserResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Token     string `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	userModel := self.Model

	response := UserResponse{
		Email:     userModel.Email,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Token:     common.GenToken(userModel.ID),
	}

	return response
}
