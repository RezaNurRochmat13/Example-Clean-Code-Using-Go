package user

import "clean-arch/modules/user/model"

// Repository abstraction
type Repository interface {
	FindAll() ([]model.User, error)
	FindByID(id string) ([]model.User, error)
}
