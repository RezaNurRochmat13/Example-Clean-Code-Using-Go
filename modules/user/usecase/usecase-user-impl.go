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

func (userUsecaseImpl *userUseCaseImpl) FindUserByID(id string) ([]model.User, error) {
	findUserByIDResult, errorHandlerUserRepo := userUsecaseImpl.userRepository.FindByID(id)

	if !utils.GlobalErrorWithBool(errorHandlerUserRepo) {
		return nil, errorHandlerUserRepo
	}

	return findUserByIDResult, nil
}

func (userUsecaseImpl *userUseCaseImpl) SaveUser(userPayload model.User) (model.User, error) {
	saveUserRepo, errorHandlerRepo := userUsecaseImpl.userRepository.Save(userPayload)

	if !utils.GlobalErrorWithBool(errorHandlerRepo) {
		return model.User{}, errorHandlerRepo
	}

	return saveUserRepo, nil
}

func (userUsecaseImpl *userUseCaseImpl) UpdateUser(id string, userUpdate model.User) (model.User, error) {
	findUserRepo, errorHandlerRepo := userUsecaseImpl.userRepository.FindByID(id)

	if !utils.GlobalErrorWithBool(errorHandlerRepo) {
		return model.User{}, errorHandlerRepo
	}

	if findUserRepo == nil {
		return model.User{}, nil
	}

	updateUserRepo, errorHandlerRepo := userUsecaseImpl.userRepository.Update(id, userUpdate)

	if !utils.GlobalErrorWithBool(errorHandlerRepo) {
		return model.User{}, errorHandlerRepo
	}

	return updateUserRepo, nil

}

func (userUsecaseImpl *userUseCaseImpl) DeleteUser(id string) error {

	errorHandlerDeleteUserRepo := userUsecaseImpl.userRepository.Delete(id)

	if !utils.GlobalErrorWithBool(errorHandlerDeleteUserRepo) {
		return errorHandlerDeleteUserRepo
	}

	return nil
}
