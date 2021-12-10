// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/productusecase/fetch_product_by_id_usecase.go

// Package mockproductusecase is a generated GoMock package.
package mockproductusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	productusecase "github.com/onituka/agile-project-management/project-management/usecase/productusecase"
)

// MockFetchProductByIDUsecase is a mock of FetchProductByIDUsecase interface.
type MockFetchProductByIDUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockFetchProductByIDUsecaseMockRecorder
}

// MockFetchProductByIDUsecaseMockRecorder is the mock recorder for MockFetchProductByIDUsecase.
type MockFetchProductByIDUsecaseMockRecorder struct {
	mock *MockFetchProductByIDUsecase
}

// NewMockFetchProductByIDUsecase creates a new mock instance.
func NewMockFetchProductByIDUsecase(ctrl *gomock.Controller) *MockFetchProductByIDUsecase {
	mock := &MockFetchProductByIDUsecase{ctrl: ctrl}
	mock.recorder = &MockFetchProductByIDUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFetchProductByIDUsecase) EXPECT() *MockFetchProductByIDUsecaseMockRecorder {
	return m.recorder
}

// FetchProductByID mocks base method.
func (m *MockFetchProductByIDUsecase) FetchProductByID(ctx context.Context, in *productusecase.FetchProductByIDInput) (*productusecase.FetchProductByIDOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProductByID", ctx, in)
	ret0, _ := ret[0].(*productusecase.FetchProductByIDOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProductByID indicates an expected call of FetchProductByID.
func (mr *MockFetchProductByIDUsecaseMockRecorder) FetchProductByID(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProductByID", reflect.TypeOf((*MockFetchProductByIDUsecase)(nil).FetchProductByID), ctx, in)
}
