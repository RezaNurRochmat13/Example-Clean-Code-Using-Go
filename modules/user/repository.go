package user

import "clean-arch/modules/user/model"

// Repository abstraction
type Repository interface {
	FindAll() ([]model.User, error)
	FindByID(id string) ([]model.User, error)
	Save(user model.User) (model.User, error)
	Update(id string, user model.User) (model.User, error)
	Delete(id string) error
}
