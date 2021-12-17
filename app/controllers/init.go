package controllers

import (
	"github.com/labstack/echo/v4"
	usecase "github.com/usernamesalah/quiz-master/app/usecases"
	"gorm.io/gorm"
)

func InitAll(e *echo.Echo, dbConnection *gorm.DB) {
	// init use case
	questionUcase := usecase.NewQuestionUsecase(dbConnection)

	// init controllers
	NewQuestionHandler(e, questionUcase)
}
