package repository

import (
	"github.com/usernamesalah/quiz-master/internal/datatransfers"
	"github.com/usernamesalah/quiz-master/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuestionRepository interface {
	Create(question *models.Question, db *gorm.DB) (err error)
	GetAll(params *datatransfers.ListQueryParams) (questions []*models.Question, cnt int64, err error)
	GetByID(questionID int) (question *models.Question, err error)
	Update(question *models.Question, db *gorm.DB) (err error)
}
type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Create(question *models.Question, db *gorm.DB) (err error) {
	err = db.Omit(clause.Associations).Model(question).Create(question).Error
	return
}

func (r *questionRepository) GetAll(params *datatransfers.ListQueryParams) (questions []*models.Question, cnt int64, err error) {
	qs := r.db.Where("deleted_at IS NULL")

	err = qs.Model(&models.Question{}).Count(&cnt).Error
	if err != nil {
		return
	}

	if params.Limit > 0 {
		qs = qs.Limit(params.Limit)
	}

	if params.Offset > 0 {
		qs = qs.Offset(params.Offset)
	}

	err = qs.Find(&questions).Error
	return
}

func (r *questionRepository) GetByID(questionID int) (question *models.Question, err error) {
	var result models.Question
	err = r.db.Where("id = ? and deleted_at IS NULL", questionID).First(&result).Error
	if err != nil {
		return
	}
	return &result, nil
}

func (r *questionRepository) Update(question *models.Question, db *gorm.DB) (err error) {
	err = db.Model(models.Question{ID: question.ID}).Updates(question).Error
	return
}
