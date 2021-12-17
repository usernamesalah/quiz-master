package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/usernamesalah/quiz-master/internal/constants"
	"github.com/usernamesalah/quiz-master/internal/datatransfers"
)

func ReturnError(c echo.Context, errorCode int, httpStatus int) error {
	statusCode := httpStatus

	response := &datatransfers.Response{
		Success: false,
		Error: &datatransfers.ErrorData{
			Message: constants.MessageEnum(errorCode).String(),
			Code:    int(constants.MessageEnum(errorCode)),
			Status:  statusCode,
		},
	}

	return c.JSON(statusCode, response)
}
