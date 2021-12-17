package usecase

import (
	"strconv"
	"strings"

	"github.com/divan/num2words"
	"github.com/usernamesalah/quiz-master/internal/datatransfers"
	"github.com/usernamesalah/quiz-master/internal/utils"
	"github.com/usernamesalah/quiz-master/pkg/models"
	"github.com/usernamesalah/quiz-master/pkg/repository"
	"gorm.io/gorm"
)

type QuestionUsecase interface {
	AnswerQuestion(answer *datatransfers.Answer) (value string, err error)
	Create(question *models.Question) (err error)
	Delete(question *models.Question) (err error)
	GetAll(param *datatransfers.ListQueryParams) (questions []*models.Question, cnt int64, err error)
	GetByID(questionID int) (question *models.Question, err error)
	Update(question *models.Question) (err error)
}

type questionUsecase struct {
	db           *gorm.DB
	questionRepo repository.QuestionRepository
}

func NewQuestionUsecase(db *gorm.DB) QuestionUsecase {
	questionRepo := repository.NewQuestionRepository(db)
	return &questionUsecase{
		questionRepo: questionRepo,
		db:           db,
	}
}

func (u *questionUsecase) GetAll(param *datatransfers.ListQueryParams) (questions []*models.Question, cnt int64, err error) {
	questions, cnt, err = u.questionRepo.GetAll(param)
	return
}

func (u *questionUsecase) GetByID(questionID int) (question *models.Question, err error) {
	question, err = u.questionRepo.GetByID(questionID)
	return
}

func (u *questionUsecase) Create(question *models.Question) (err error) {
	err = u.questionRepo.Create(question, u.db)
	return
}

func (u *questionUsecase) Update(question *models.Question) (err error) {
	err = u.questionRepo.Update(question, u.db)
	return
}

func (u *questionUsecase) Delete(question *models.Question) (err error) {
	question.DeletedAt = utils.Now()
	err = u.questionRepo.Update(question, u.db)
	return
}

func (u *questionUsecase) AnswerQuestion(answer *datatransfers.Answer) (value string, err error) {
	value = "Incorrect"

	question, err := u.questionRepo.GetByID(answer.QuestionID)
	if err != nil {
		return
	}

	if strings.ToLower(answer.Value) == question.Answer {
		value = "Correct"
		return
	} else {
		answerINT, _ := strconv.Atoi(question.Answer)
		if strNumber := num2words.Convert(answerINT); strNumber == strings.ToLower(answer.Value) {
			value = "Correct"
			return
		}
	}

	return
}
