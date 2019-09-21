package http

import (
	"clean-arch/modules/user"
	"clean-arch/modules/user/model"
	"clean-arch/utils"
	"log"
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
	groupingPath.POST("/users", handlerWithInjection.CreateNewUsers)
	groupingPath.PUT("/users/:id", handlerWithInjection.UpdateUser)

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

func (userHandler *UserHandler) CreateNewUsers(ctx echo.Context) error {
	var modelUser model.User

	errorHandlerBind := ctx.Bind(&modelUser)

	if !utils.GlobalErrorWithBool(errorHandlerBind) {
		log.Printf("Error when access usecase : %s", errorHandlerBind)
		return ctx.JSON(http.StatusBadRequest,
			echo.Map{"error": "Invalid body request"})
	}

	saveUserUsecase, errorHandlerUseCase := userHandler.UserUseCase.SaveUser(modelUser)

	if !utils.GlobalErrorWithBool(errorHandlerUseCase) {
		log.Printf("Error when access usecase : %s", errorHandlerUseCase)
		return ctx.JSON(http.StatusBadRequest,
			echo.Map{"error": "System cannot process. View logs more info"})
	}

	return ctx.JSON(http.StatusCreated,
		echo.Map{"success": http.StatusCreated,
			"created_resource": saveUserUsecase})
}

func (userHandler *UserHandler) UpdateUser(ctx echo.Context) error {
	var (
		modelUser model.User
		id        = ctx.Param("id")
	)

	errorHandlerBindJSON := ctx.Bind(&modelUser)

	if !utils.GlobalErrorWithBool(errorHandlerBindJSON) {
		log.Printf("Error when bind json : %s", errorHandlerBindJSON)
		return ctx.JSON(http.StatusBadRequest,
			echo.Map{"error": "Invalid body request. View logs more info"})
	}

	updateUserUsecase, errorHandlerUseCase := userHandler.UserUseCase.UpdateUser(id, modelUser)

	if !utils.GlobalErrorWithBool(errorHandlerUseCase) {
		log.Printf("Error when bind json : %s", errorHandlerUseCase)
		return ctx.JSON(http.StatusBadRequest,
			echo.Map{"error": "Invalid request. View logs more info"})
	}

	if updateUserUsecase.Username == "" {
		return ctx.JSON(http.StatusBadRequest,
			echo.Map{"error": "Data Not Found. View logs more info"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"success":      http.StatusOK,
		"updated_user": updateUserUsecase})
}
