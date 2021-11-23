// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/productusecase/create_product_usecase.go

// Package mockproductusecase is a generated GoMock package.
package mockproductusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	productusecase "github.com/onituka/agile-project-management/project-management/usecase/productusecase"
)

// MockCreateProductUsecase is a mock of CreateProductUsecase interface.
type MockCreateProductUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateProductUsecaseMockRecorder
}

// MockCreateProductUsecaseMockRecorder is the mock recorder for MockCreateProductUsecase.
type MockCreateProductUsecaseMockRecorder struct {
	mock *MockCreateProductUsecase
}

// NewMockCreateProductUsecase creates a new mock instance.
func NewMockCreateProductUsecase(ctrl *gomock.Controller) *MockCreateProductUsecase {
	mock := &MockCreateProductUsecase{ctrl: ctrl}
	mock.recorder = &MockCreateProductUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateProductUsecase) EXPECT() *MockCreateProductUsecaseMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockCreateProductUsecase) CreateProduct(ctx context.Context, in *productusecase.CreateProductInput) (*productusecase.CreateProductOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, in)
	ret0, _ := ret[0].(*productusecase.CreateProductOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockCreateProductUsecaseMockRecorder) CreateProduct(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockCreateProductUsecase)(nil).CreateProduct), ctx, in)
}