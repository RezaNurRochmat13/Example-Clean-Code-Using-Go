package main

import (
	"clean-arch/config/database"
	"clean-arch/utils"
	"fmt"
	"log"
	"net/http"

	_userHandler "clean-arch/modules/user/delivery/http"
	_userRepo "clean-arch/modules/user/repository"
	_userUsecase "clean-arch/modules/user/usecase"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Halo server :)")

	databaseConn, errorHandlerDatabaseConn := database.MysqlDevelopmentConfiguration()

	if !utils.GlobalErrorWithBool(errorHandlerDatabaseConn) {
		log.Printf("Error open database conn : %s", errorHandlerDatabaseConn)
	}

	echoRouter := echo.New()
	echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	userRepository := _userRepo.NewUserRepositoryImpl(databaseConn)
	userUsecase := _userUsecase.NewUserUseCase(userRepository)
	_userHandler.NewUserHandler(echoRouter, userUsecase)

	echoRouter.Logger.Fatal(echoRouter.Start(":8081"))
}
