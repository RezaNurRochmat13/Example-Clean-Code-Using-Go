package usecase

import (
	"clean-arch/modules/user"
	"clean-arch/modules/user/model"
	"clean-arch/utils"
)

type userUseCaseImpl struct {
	userRepository user.Repository
}

func NewUserUseCase(userRepo user.Repository) user.Usecase {
	return &userUseCaseImpl{
		userRepository: userRepo,
	}
}

func (userUsecaseImpl *userUseCaseImpl) FindAllUser() ([]model.User, error) {
	findAllUserResult, errorHandlerUserRepo := userUsecaseImpl.userRepository.FindAll()
	if !utils.GlobalErrorWithBool(errorHandlerUserRepo) {
		return nil, errorHandlerUserRepo
	}

	return findAllUserResult, nil
}
