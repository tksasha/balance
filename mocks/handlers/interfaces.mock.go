// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handlers/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/handlers/interfaces.go -package mocksforhandlers -destination mocks/handlers/interfaces.mock.go
//

// Package mocksforhandlers is a generated GoMock package.
package mocksforhandlers

import (
	context "context"
	reflect "reflect"

	models "github.com/tksasha/balance/internal/models"
	requests "github.com/tksasha/balance/internal/requests"
	gomock "go.uber.org/mock/gomock"
)

// MockItemService is a mock of ItemService interface.
type MockItemService struct {
	ctrl     *gomock.Controller
	recorder *MockItemServiceMockRecorder
	isgomock struct{}
}

// MockItemServiceMockRecorder is the mock recorder for MockItemService.
type MockItemServiceMockRecorder struct {
	mock *MockItemService
}

// NewMockItemService creates a new mock instance.
func NewMockItemService(ctrl *gomock.Controller) *MockItemService {
	mock := &MockItemService{ctrl: ctrl}
	mock.recorder = &MockItemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemService) EXPECT() *MockItemServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockItemService) Create(ctx context.Context, request requests.ItemCreateRequest) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockItemServiceMockRecorder) Create(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemService)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockItemService) Delete(ctx context.Context, input string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockItemServiceMockRecorder) Delete(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItemService)(nil).Delete), ctx, input)
}

// GetItem mocks base method.
func (m *MockItemService) GetItem(ctx context.Context, input string) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", ctx, input)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockItemServiceMockRecorder) GetItem(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockItemService)(nil).GetItem), ctx, input)
}

// GetItems mocks base method.
func (m *MockItemService) GetItems(ctx context.Context) (models.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", ctx)
	ret0, _ := ret[0].(models.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockItemServiceMockRecorder) GetItems(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockItemService)(nil).GetItems), ctx)
}

// Update mocks base method.
func (m *MockItemService) Update(ctx context.Context, request requests.ItemUpdateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockItemServiceMockRecorder) Update(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItemService)(nil).Update), ctx, request)
}

// MockCategoryService is a mock of CategoryService interface.
type MockCategoryService struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryServiceMockRecorder
	isgomock struct{}
}

// MockCategoryServiceMockRecorder is the mock recorder for MockCategoryService.
type MockCategoryServiceMockRecorder struct {
	mock *MockCategoryService
}

// NewMockCategoryService creates a new mock instance.
func NewMockCategoryService(ctrl *gomock.Controller) *MockCategoryService {
	mock := &MockCategoryService{ctrl: ctrl}
	mock.recorder = &MockCategoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryService) EXPECT() *MockCategoryServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryService) Create(ctx context.Context, request requests.CategoryCreateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCategoryServiceMockRecorder) Create(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryService)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockCategoryService) Delete(ctx context.Context, category *models.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryServiceMockRecorder) Delete(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategoryService)(nil).Delete), ctx, category)
}

// FindByID mocks base method.
func (m *MockCategoryService) FindByID(ctx context.Context, id int) (*models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCategoryServiceMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCategoryService)(nil).FindByID), ctx, id)
}

// GetAll mocks base method.
func (m *MockCategoryService) GetAll(ctx context.Context) (models.Categories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(models.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCategoryServiceMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCategoryService)(nil).GetAll), ctx)
}

// Update mocks base method.
func (m *MockCategoryService) Update(ctx context.Context, category *models.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCategoryServiceMockRecorder) Update(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryService)(nil).Update), ctx, category)
}

// MockCashService is a mock of CashService interface.
type MockCashService struct {
	ctrl     *gomock.Controller
	recorder *MockCashServiceMockRecorder
	isgomock struct{}
}

// MockCashServiceMockRecorder is the mock recorder for MockCashService.
type MockCashServiceMockRecorder struct {
	mock *MockCashService
}

// NewMockCashService creates a new mock instance.
func NewMockCashService(ctrl *gomock.Controller) *MockCashService {
	mock := &MockCashService{ctrl: ctrl}
	mock.recorder = &MockCashServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCashService) EXPECT() *MockCashServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCashService) Create(ctx context.Context, request requests.CashCreateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCashServiceMockRecorder) Create(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCashService)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockCashService) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCashServiceMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCashService)(nil).Delete), ctx, id)
}

// FindByID mocks base method.
func (m *MockCashService) FindByID(ctx context.Context, id string) (*models.Cash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*models.Cash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCashServiceMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCashService)(nil).FindByID), ctx, id)
}

// Update mocks base method.
func (m *MockCashService) Update(ctx context.Context, request requests.CashUpdateRequest) (*models.Cash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, request)
	ret0, _ := ret[0].(*models.Cash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCashServiceMockRecorder) Update(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCashService)(nil).Update), ctx, request)
}
