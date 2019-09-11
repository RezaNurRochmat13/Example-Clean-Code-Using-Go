package repository

import (
	"clean-arch/modules/user"
	"clean-arch/modules/user/model"
	"clean-arch/utils"
	"database/sql"
)

// Implement the repository interface
type userRepositoryImpl struct {
	Connection *sql.DB
}

func NewUserRepositoryImpl(Conn *sql.DB) user.Repository {
	return &userRepositoryImpl{Conn}
}

func (userRepoImpl *userRepositoryImpl) FindAll() ([]model.User, error) {
	var (
		queryStmt = "SELECT id_user, user_name, user_address, " +
			"user_phone, user_age FROM user"
		userModel  model.User
		resultUser []model.User
	)

	queryFindUsers, errorHandlerQuery := userRepoImpl.Connection.Query(queryStmt)

	if !utils.GlobalQueryErrorWithBool(errorHandlerQuery) {
		return nil, errorHandlerQuery
	}

	for queryFindUsers.Next() {
		errorHandlerScan := queryFindUsers.Scan(
			&userModel.IDUser,
			&userModel.Username,
			&userModel.UserAddress,
			&userModel.UserPhone,
			&userModel.UserAge)

		if !utils.GlobalQueryErrorWithBool(errorHandlerScan) {
			return nil, errorHandlerScan
		}

		resultUser = append(resultUser, userModel)
	}

	return resultUser, nil
}
