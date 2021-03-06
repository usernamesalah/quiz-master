// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	datatransfers "github.com/usernamesalah/quiz-master/internal/datatransfers"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "github.com/usernamesalah/quiz-master/pkg/models"
)

// QuestionRepository is an autogenerated mock type for the QuestionRepository type
type QuestionRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: question, db
func (_m *QuestionRepository) Create(question *models.Question, db *gorm.DB) error {
	ret := _m.Called(question, db)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Question, *gorm.DB) error); ok {
		r0 = rf(question, db)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: params
func (_m *QuestionRepository) GetAll(params *datatransfers.ListQueryParams) ([]*models.Question, int64, error) {
	ret := _m.Called(params)

	var r0 []*models.Question
	if rf, ok := ret.Get(0).(func(*datatransfers.ListQueryParams) []*models.Question); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Question)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(*datatransfers.ListQueryParams) int64); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*datatransfers.ListQueryParams) error); ok {
		r2 = rf(params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: questionID
func (_m *QuestionRepository) GetByID(questionID int) (*models.Question, error) {
	ret := _m.Called(questionID)

	var r0 *models.Question
	if rf, ok := ret.Get(0).(func(int) *models.Question); ok {
		r0 = rf(questionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Question)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(questionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: question, db
func (_m *QuestionRepository) Update(question *models.Question, db *gorm.DB) error {
	ret := _m.Called(question, db)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Question, *gorm.DB) error); ok {
		r0 = rf(question, db)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
