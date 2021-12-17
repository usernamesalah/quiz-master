package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	usecase "github.com/usernamesalah/quiz-master/app/usecases"
	"github.com/usernamesalah/quiz-master/internal/datatransfers"
	"github.com/usernamesalah/quiz-master/internal/utils"
	"github.com/usernamesalah/quiz-master/pkg/models"
	"gorm.io/gorm"
)

// QuestionHandler  represent the httphandler
type QuestionHandler struct {
	questionUseCase usecase.QuestionUsecase
}

// NewQuestionHandler will initialize the questionUseCase resources endpoint
func NewQuestionHandler(e *echo.Echo,
	questionUseCase usecase.QuestionUsecase,
) {

	handler := &QuestionHandler{
		questionUseCase: questionUseCase,
	}

	questions := e.Group("/questions")
	{
		questions.GET("", handler.GetAllQuestion)
		questions.GET("/:id", handler.GetQuestionByID)
		questions.DELETE("/:id", handler.GetQuestionByID)
		questions.PUT("/:id", handler.UpdateQuestion)
		questions.POST("", handler.CreateQuestion)
	}
}

// Create a new questions
// @Summary Create a new questions
// @Description Create a new questions
// @Tags questions
// @ID create-questions
// @Produce json
// @Param questions body models.Question true "Create questions"
// @Success 201 {object} models.Question
// @Router /questions [post]
func (h *QuestionHandler) CreateQuestion(c echo.Context) error {
	questions := new(models.Question)
	if err := c.Bind(questions); err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(questions); err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}

	err := h.questionUseCase.Create(questions)
	if err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, questions)
}

// Delete an questions
// @Summary Delete an questions
// @Description Delete an questions by id
// @Tags questions
// @ID delete-questions
// @Produce plain
// @Param id path int true "Question ID"
// @Success 204 {string} string ""
// @Router /questions/{id} [delete]
func (h *QuestionHandler) DeleteQuestion(c echo.Context) error {

	idString := c.Param("id")
	questionID, _ := strconv.Atoi(idString)

	question := &models.Question{ID: questionID}
	err := h.questionUseCase.Delete(question)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ReturnInvalidResponse(http.StatusInternalServerError, err.Error())
		}
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}

// Update an questions
// @Summary Update an questions
// @Description Update an questions by id
// @Tags questions
// @ID update-questions
// @Produce json
// @Param questions body models.Question true "Update questions"
// @Param id path int true "Question ID"
// @Success 200 {object} models.Question"
// @Router /questions/{id} [put]
func (h *QuestionHandler) UpdateQuestion(c echo.Context) error {

	idString := c.Param("id")
	questionID, _ := strconv.Atoi(idString)

	question := new(models.Question)
	if err := c.Bind(question); err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(question); err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}

	question.ID = questionID
	err := h.questionUseCase.Update(question)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ReturnInvalidResponse(http.StatusInternalServerError, err.Error())
		}
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, question)
}

// List questions
// @Summary List questions
// @Description Get the list of questions
// @Tags questions
// @ID list-questions
// @Produce json
// @Param page query integer false "page number" default(1)
// @Param limit query integer false "number of questions in single page" default(10)
// @Success 200 {array} models.Question
// @Router /questions [get]
func (h *QuestionHandler) GetAllQuestion(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	questions, cnt, err := h.questionUseCase.GetAll(&datatransfers.ListQueryParams{
		Limit:  limit,
		Offset: utils.CalculateOffset(limit, page),
	})
	if err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}

	response := &datatransfers.Response{
		Success: true,
		Data:    questions,
		Paging:  utils.SetPaginator(limit, page, cnt, c),
	}
	return c.JSON(http.StatusOK, response)
}

// Get a question by id
// @Summary Get a question by id
// @Description Get a question by id
// @Tags questions
// @ID get-question-id
// @Produce json
// @Param id path int true "id Question"
// @Success 200 {object} models.Question
// @Router /questions/{id} [get]
func (h *QuestionHandler) GetQuestionByID(c echo.Context) error {
	idString := c.Param("id")
	questionID, _ := strconv.Atoi(idString)

	// check by id
	question, err := h.questionUseCase.GetByID(questionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ReturnInvalidResponse(http.StatusInternalServerError, err.Error())
		}
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}

	response := &datatransfers.Response{
		Success: true,
		Data:    question,
	}

	return c.JSON(http.StatusOK, response)
}

// Answer a new questions
// @Summary Answer a new questions
// @Description Answer a new questions
// @Tags questions
// @ID answer-questions
// @Produce json
// @Param questions body datatransfers.Answer true "Answer questions"
// @Success 200
// @Router /questions/answer/{id} [post]
func (h *QuestionHandler) AnswerQuestion(c echo.Context) error {
	idString := c.Param("id")
	questionID, _ := strconv.Atoi(idString)

	answer := new(datatransfers.Answer)
	if err := c.Bind(answer); err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(answer); err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, err.Error())
	}

	answer.QuestionID = questionID
	value, err := h.questionUseCase.AnswerQuestion(answer)
	if err != nil {
		return utils.ReturnInvalidResponse(http.StatusBadRequest, "Incorrect")
	}

	return c.JSON(http.StatusOK, value)
}
