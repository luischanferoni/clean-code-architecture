// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	domain "opensea/internal/domain"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockOpenseaServiceContract is a mock of OpenseaServiceContract interface.
type MockOpenseaServiceContract struct {
	ctrl     *gomock.Controller
	recorder *MockOpenseaServiceContractMockRecorder
}

// MockOpenseaServiceContractMockRecorder is the mock recorder for MockOpenseaServiceContract.
type MockOpenseaServiceContractMockRecorder struct {
	mock *MockOpenseaServiceContract
}

// NewMockOpenseaServiceContract creates a new mock instance.
func NewMockOpenseaServiceContract(ctrl *gomock.Controller) *MockOpenseaServiceContract {
	mock := &MockOpenseaServiceContract{ctrl: ctrl}
	mock.recorder = &MockOpenseaServiceContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenseaServiceContract) EXPECT() *MockOpenseaServiceContractMockRecorder {
	return m.recorder
}

// Buy mocks base method.
func (m *MockOpenseaServiceContract) Buy(ctx *gin.Context, movieId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Buy", ctx, movieId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Buy indicates an expected call of Buy.
func (mr *MockOpenseaServiceContractMockRecorder) Buy(ctx, movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Buy", reflect.TypeOf((*MockOpenseaServiceContract)(nil).Buy), ctx, movieId)
}

// Create mocks base method.
func (m *MockOpenseaServiceContract) Create(ctx *gin.Context, movie *domain.Movie) (*domain.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, movie)
	ret0, _ := ret[0].(*domain.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOpenseaServiceContractMockRecorder) Create(ctx, movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOpenseaServiceContract)(nil).Create), ctx, movie)
}

// Get mocks base method.
func (m *MockOpenseaServiceContract) Get(ctx *gin.Context, movieId int64) (*domain.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, movieId)
	ret0, _ := ret[0].(*domain.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOpenseaServiceContractMockRecorder) Get(ctx, movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOpenseaServiceContract)(nil).Get), ctx, movieId)
}

// GetAll mocks base method.
func (m *MockOpenseaServiceContract) GetAll(ctx *gin.Context, page int) ([]*domain.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, page)
	ret0, _ := ret[0].([]*domain.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockOpenseaServiceContractMockRecorder) GetAll(ctx, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockOpenseaServiceContract)(nil).GetAll), ctx, page)
}