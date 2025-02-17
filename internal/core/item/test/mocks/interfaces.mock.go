// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/item/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/core/item/interfaces.go -package mocks -destination internal/core/item/test/mocks/interfaces.mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	item "github.com/tksasha/balance/internal/core/item"
	month "github.com/tksasha/month"
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

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, item *item.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, item)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, id)
}

// FindAllByMonth mocks base method.
func (m *MockRepository) FindAllByMonth(ctx context.Context, month month.Month) (item.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllByMonth", ctx, month)
	ret0, _ := ret[0].(item.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByMonth indicates an expected call of FindAllByMonth.
func (mr *MockRepositoryMockRecorder) FindAllByMonth(ctx, month any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByMonth", reflect.TypeOf((*MockRepository)(nil).FindAllByMonth), ctx, month)
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(ctx context.Context, id int) (*item.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*item.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), ctx, id)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, item *item.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, item)
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

// Create mocks base method.
func (m *MockService) Create(ctx context.Context, request item.CreateRequest) (*item.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(*item.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockService) Delete(ctx context.Context, input string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), ctx, input)
}

// Edit mocks base method.
func (m *MockService) Edit(ctx context.Context, input string) (*item.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", ctx, input)
	ret0, _ := ret[0].(*item.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Edit indicates an expected call of Edit.
func (mr *MockServiceMockRecorder) Edit(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockService)(nil).Edit), ctx, input)
}

// Index mocks base method.
func (m *MockService) Index(ctx context.Context, request item.IndexRequest) (item.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", ctx, request)
	ret0, _ := ret[0].(item.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Index indicates an expected call of Index.
func (mr *MockServiceMockRecorder) Index(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockService)(nil).Index), ctx, request)
}

// Update mocks base method.
func (m *MockService) Update(ctx context.Context, request item.UpdateRequest) (*item.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, request)
	ret0, _ := ret[0].(*item.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), ctx, request)
}
