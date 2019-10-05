package user

import "clean-arch/modules/user/model"

// Usecase abstraction
type Usecase interface {
	FindAllUser() ([]model.User, error)
	FindUserByID(id string) ([]model.User, error)
	SaveUser(user model.User) (model.User, error)
	UpdateUser(id string, userUpdate model.User) (model.User, error)
	DeleteUser(id string) error
}
