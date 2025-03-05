// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/category/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/app/category/interfaces.go -package mocks -destination internal/app/category/test/mocks/interfaces.mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	category "github.com/tksasha/balance/internal/app/category"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
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

// FindAll mocks base method.
func (m *MockRepository) FindAll(ctx context.Context) (category.Categories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].(category.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRepositoryMockRecorder) FindAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRepository)(nil).FindAll), ctx)
}

// FindAllByFilters mocks base method.
func (m *MockRepository) FindAllByFilters(ctx context.Context, filters category.Filters) (category.Categories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllByFilters", ctx, filters)
	ret0, _ := ret[0].(category.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByFilters indicates an expected call of FindAllByFilters.
func (mr *MockRepositoryMockRecorder) FindAllByFilters(ctx, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByFilters", reflect.TypeOf((*MockRepository)(nil).FindAllByFilters), ctx, filters)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GroupedList mocks base method.
func (m *MockService) GroupedList(ctx context.Context, request category.Request) (category.GroupedCategories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupedList", ctx, request)
	ret0, _ := ret[0].(category.GroupedCategories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupedList indicates an expected call of GroupedList.
func (mr *MockServiceMockRecorder) GroupedList(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupedList", reflect.TypeOf((*MockService)(nil).GroupedList), ctx, request)
}
