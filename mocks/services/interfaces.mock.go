// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/services/interfaces.go -package mocksforservices -destination mocks/services/interfaces.mock.go
//

// Package mocksforservices is a generated GoMock package.
package mocksforservices

import (
	context "context"
	reflect "reflect"

	models "github.com/tksasha/balance/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockItemRepository is a mock of ItemRepository interface.
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
	isgomock struct{}
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository.
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance.
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockItemRepository) Create(ctx context.Context, item *models.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockItemRepositoryMockRecorder) Create(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemRepository)(nil).Create), ctx, item)
}

// Delete mocks base method.
func (m *MockItemRepository) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockItemRepositoryMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItemRepository)(nil).Delete), ctx, id)
}

// FindByID mocks base method.
func (m *MockItemRepository) FindByID(ctx context.Context, id int) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockItemRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockItemRepository)(nil).FindByID), ctx, id)
}

// GetItems mocks base method.
func (m *MockItemRepository) GetItems(ctx context.Context) (models.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", ctx)
	ret0, _ := ret[0].(models.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockItemRepositoryMockRecorder) GetItems(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockItemRepository)(nil).GetItems), ctx)
}

// Update mocks base method.
func (m *MockItemRepository) Update(ctx context.Context, item *models.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockItemRepositoryMockRecorder) Update(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItemRepository)(nil).Update), ctx, item)
}
