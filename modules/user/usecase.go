package user

import "clean-arch/modules/user/model"

// Usecase abstraction
type Usecase interface {
	FindAllUser() ([]model.User, error)
	FindUserByID(id string) ([]model.User, error)
}
