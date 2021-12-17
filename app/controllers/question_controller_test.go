package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/usernamesalah/quiz-master/app/usecases/mocks"
	"github.com/usernamesalah/quiz-master/pkg/models"
)

type mockRequestValidator struct{}

func (m *mockRequestValidator) Validate(i interface{}) error {
	return nil
}

func TestAPI_listQuestions(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/questions", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	mockQuestionUsecase := &mocks.QuestionUsecase{}
	mockQuestionUsecase.On("GetAll", mock.Anything).Return([]models.Question{}, nil)

	api := &QuestionHandler{questionUseCase: mockQuestionUsecase}
	if assert.NoError(t, api.GetAllQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())
	}
}

func TestAPI_getQuestion(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/questions/1", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/questions/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockQuestionUsecase := &mocks.QuestionUsecase{}
	mockQuestionUsecase.On("GetQuestionByID", mock.Anything, int64(1)).Return(models.Question{}, nil)

	api := &QuestionHandler{questionUseCase: mockQuestionUsecase}
	if assert.NoError(t, api.GetQuestionByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":0,\"question\":\"\",\"answer\":\"\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":\"0001-01-01T00:00:00Z\"}\n", rec.Body.String())
	}
}

func TestAPI_createQuestion(t *testing.T) {
	question := models.Question{
		ID:       1,
		Question: "How many letters are there in the English alphabet?",
		Answer:   "5",
	}
	questionJSON, _ := json.Marshal(question)

	req := httptest.NewRequest(http.MethodPost, "/questions", bytes.NewReader(questionJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.Validator = &mockRequestValidator{}
	c := e.NewContext(req, rec)

	mockQuestionUsecase := &mocks.QuestionUsecase{}
	mockQuestionUsecase.On("CreateQuestion", mock.Anything, question).Return(question, nil)

	api := &QuestionHandler{questionUseCase: mockQuestionUsecase}
	if assert.NoError(t, api.CreateQuestion(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(questionJSON)+"\n", rec.Body.String())
	}
}

func TestAPI_updateQuestion(t *testing.T) {
	question := models.Question{
		ID:       1,
		Question: "How many letters are there in the English alphabet?",
		Answer:   "6",
	}
	questionJSON, _ := json.Marshal(question)

	req := httptest.NewRequest(http.MethodPut, "/questions/1", bytes.NewReader(questionJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.Validator = &mockRequestValidator{}
	c := e.NewContext(req, rec)
	c.SetPath("/questions/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockQuestionUsecase := &mocks.QuestionUsecase{}
	mockQuestionUsecase.On("UpdateQuestion", req.Context(), question).Return(question, nil)

	api := &QuestionHandler{questionUseCase: mockQuestionUsecase}
	if assert.NoError(t, api.UpdateQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(questionJSON)+"\n", rec.Body.String())
	}
}

func TestAPI_deleteQuestion(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/questions/1", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/questions/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockQuestionUsecase := &mocks.QuestionUsecase{}
	mockQuestionUsecase.On("DeleteQuestion", mock.Anything, int64(1)).Return(nil)

	api := &QuestionHandler{questionUseCase: mockQuestionUsecase}
	if assert.NoError(t, api.DeleteQuestion(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestAPI_getAnswerQuestion(t *testing.T) {
	question := models.Question{
		Answer: "6",
	}
	questionJSON, _ := json.Marshal(question)

	req := httptest.NewRequest(http.MethodPut, "/questions/answers/1", bytes.NewReader(questionJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	e.Validator = &mockRequestValidator{}
	c := e.NewContext(req, rec)
	c.SetPath("/questions/answers/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockQuestionUsecase := &mocks.QuestionUsecase{}
	mockQuestionUsecase.On("AnswerQuestion", req.Context(), question).Return(question, nil)

	api := &QuestionHandler{questionUseCase: mockQuestionUsecase}
	if assert.NoError(t, api.AnswerQuestion(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Incorrect", rec.Body.String())
	}
}
