// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/abdghn/alpha-indo-soft-be-test/repository (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/abdghn/alpha-indo-soft-be-test/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 *models.Article) (*models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0)
}

// GetArticles mocks base method.
func (m *MockRepository) GetArticles(arg0 map[string]interface{}, arg1 string) ([]*models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles", arg0, arg1)
	ret0, _ := ret[0].([]*models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockRepositoryMockRecorder) GetArticles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockRepository)(nil).GetArticles), arg0, arg1)
}
