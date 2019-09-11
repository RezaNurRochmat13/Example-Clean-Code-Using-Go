package user

import "clean-arch/modules/user/model"

// Usecase abstraction
type Usecase interface {
	FindAllUser() ([]model.User, error)
}
