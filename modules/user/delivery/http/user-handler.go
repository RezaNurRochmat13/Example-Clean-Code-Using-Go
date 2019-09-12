package http

import (
	"clean-arch/modules/user"
	"clean-arch/utils"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUseCase user.Usecase
}

func NewUserHandler(e *echo.Echo, usecase user.Usecase) {
	handlerWithInjection := &UserHandler{
		UserUseCase: usecase,
	}

	groupingPath := e.Group("/api/v1")
	groupingPath.GET("/users", handlerWithInjection.GetAllUsersHandler)
	groupingPath.GET("/users/:id", handlerWithInjection.GetDetailUsers)

}

func (userHandler *UserHandler) GetAllUsersHandler(ctx echo.Context) error {
	findAllUserUsecase, errorHandlerUsecase := userHandler.UserUseCase.FindAllUser()

	if !utils.GlobalErrorWithBool(errorHandlerUsecase) {
		return ctx.JSON(400, echo.Map{"message": "Cannot process service"})
	}

	if findAllUserUsecase != nil {
		return ctx.JSON(200, echo.Map{
			"total": len(findAllUserUsecase),
			"count": len(findAllUserUsecase),
			"data":  findAllUserUsecase})
	}

	return ctx.JSON(200, echo.Map{
		"total": len(findAllUserUsecase),
		"count": len(findAllUserUsecase),
		"data":  findAllUserUsecase})
}

func (userHandler *UserHandler) GetDetailUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	findUserByIdUsecase, errorHandlerUserUsecase := userHandler.UserUseCase.FindUserByID(id)

	if !utils.GlobalErrorWithBool(errorHandlerUserUsecase) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Cannot get usecase",
		})
	}

	if findUserByIdUsecase == nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Data not found",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"data": findUserByIdUsecase,
	})
}
