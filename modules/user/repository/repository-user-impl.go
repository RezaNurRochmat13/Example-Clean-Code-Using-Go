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

func (userRepoImpl *userRepositoryImpl) FindByID(id string) ([]model.User, error) {
	var (
		sql = "SELECT id_user, user_name, user_address, " +
			"user_phone, user_age FROM user WHERE id_user = ?"
		modelUser      model.User
		resultUserById []model.User
	)

	queryUserById, errorHandlerQuery := userRepoImpl.Connection.Query(sql, id)

	if !utils.GlobalQueryErrorWithBool(errorHandlerQuery) {
		return nil, errorHandlerQuery
	}

	for queryUserById.Next() {
		errorHandlerScan := queryUserById.Scan(
			&modelUser.IDUser,
			&modelUser.Username,
			&modelUser.UserAddress,
			&modelUser.UserPhone,
			&modelUser.UserAge)

		if !utils.GlobalQueryErrorWithBool(errorHandlerScan) {
			return nil, errorHandlerScan
		}

		resultUserById = append(resultUserById, modelUser)
	}

	return resultUserById, nil
}

func (userRepoImpl *userRepositoryImpl) Save(user model.User) (model.User, error) {
	var (
		sql = `INSERT INTO user(user_name, user_address, user_phone, user_age) ` +
			`VALUES(?, ?, ?, ?)`
	)

	initTransaction, errorHandlerTransaction := userRepoImpl.Connection.Begin()

	if !utils.GlobalQueryErrorWithBool(errorHandlerTransaction) {
		return model.User{}, errorHandlerTransaction
	}

	defer initTransaction.Rollback()

	querySaveUser, errorHandlerQuery := userRepoImpl.Connection.Prepare(sql)

	if !utils.GlobalQueryErrorWithBool(errorHandlerQuery) {
		return model.User{}, errorHandlerQuery
	}

	_, errorHandlerExecStmt := querySaveUser.Exec(user.Username,
		user.UserAddress, user.UserPhone, user.UserAge)

	if !utils.GlobalQueryErrorWithBool(errorHandlerExecStmt) {
		return model.User{}, errorHandlerExecStmt
	}

	errorHandlerCommitTrans := initTransaction.Commit()

	if !utils.GlobalQueryErrorWithBool(errorHandlerCommitTrans) {
		return model.User{}, errorHandlerCommitTrans
	}

	return user, nil
}

func (userRepoImpl *userRepositoryImpl) Update(id string, userUpdate model.User) (model.User, error) {
	sql := `UPDATE user SET user_name = ?, ` +
		`user_address = ?, user_phone = ?, user_age = ? ` +
		`WHERE id_user = ?`

	initTransaction, errorHandlerTransaction := userRepoImpl.Connection.Begin()

	if !utils.GlobalQueryErrorWithBool(errorHandlerTransaction) {
		return model.User{}, errorHandlerTransaction
	}

	defer initTransaction.Rollback()

	queryUpdateUser, errorHandlerQuery := userRepoImpl.Connection.Prepare(sql)

	if !utils.GlobalQueryErrorWithBool(errorHandlerQuery) {
		return model.User{}, errorHandlerQuery
	}

	_, errorHandlerExecStmt := queryUpdateUser.Exec(
		userUpdate.Username, userUpdate.UserAddress,
		userUpdate.UserPhone, userUpdate.UserAge, id)

	if !utils.GlobalQueryErrorWithBool(errorHandlerExecStmt) {
		return model.User{}, errorHandlerExecStmt
	}

	errorHandlerCommitTrans := initTransaction.Commit()

	if !utils.GlobalQueryErrorWithBool(errorHandlerCommitTrans) {
		return model.User{}, errorHandlerCommitTrans
	}

	return userUpdate, nil

}

func (userRepoImpl *userRepositoryImpl) Delete(id string) error {
	sql := "DELETE FROM user WHERE id_user = ?"

	initTransaction, errorHandlerTransaction := userRepoImpl.Connection.Begin()

	if !utils.GlobalQueryErrorWithBool(errorHandlerTransaction) {
		initTransaction.Rollback()
		return errorHandlerTransaction
	}

	defer initTransaction.Rollback()

	preparedStmt, errorHandlerPreparedStmt := initTransaction.Prepare(sql)

	if !utils.GlobalQueryErrorWithBool(errorHandlerPreparedStmt) {
		initTransaction.Rollback()
		return errorHandlerPreparedStmt
	}

	_, errorHandlerExecStmt := preparedStmt.Exec(id)

	if !utils.GlobalQueryErrorWithBool(errorHandlerExecStmt) {
		initTransaction.Rollback()
		return errorHandlerExecStmt
	}

	errorHandlerCommitTransaction := initTransaction.Commit()

	if !utils.GlobalQueryErrorWithBool(errorHandlerCommitTransaction) {
		initTransaction.Rollback()
		return errorHandlerCommitTransaction
	}

	return nil

}
