package http

import (
	"clean-arch/modules/user"
	"clean-arch/utils"

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
