// Code generated by MockGen. DO NOT EDIT.
// Source: domain/productdm/product_repository.go

// Package mockproductrepository is a generated GoMock package.
package mockproductrepository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	groupdm "github.com/onituka/agile-project-management/project-management/domain/groupdm"
	productdm "github.com/onituka/agile-project-management/project-management/domain/productdm"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductRepository) CreateProduct(ctx context.Context, product *productdm.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductRepositoryMockRecorder) CreateProduct(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductRepository)(nil).CreateProduct), ctx, product)
}

// FetchProductByGroupIDAndName mocks base method.
func (m *MockProductRepository) FetchProductByGroupIDAndName(ctx context.Context, groupID groupdm.GroupID, Name productdm.Name) (*productdm.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProductByGroupIDAndName", ctx, groupID, Name)
	ret0, _ := ret[0].(*productdm.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProductByGroupIDAndName indicates an expected call of FetchProductByGroupIDAndName.
func (mr *MockProductRepositoryMockRecorder) FetchProductByGroupIDAndName(ctx, groupID, Name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProductByGroupIDAndName", reflect.TypeOf((*MockProductRepository)(nil).FetchProductByGroupIDAndName), ctx, groupID, Name)
}

// FetchProductByID mocks base method.
func (m *MockProductRepository) FetchProductByID(ctx context.Context, id productdm.ProductID) (*productdm.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProductByID", ctx, id)
	ret0, _ := ret[0].(*productdm.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProductByID indicates an expected call of FetchProductByID.
func (mr *MockProductRepositoryMockRecorder) FetchProductByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProductByID", reflect.TypeOf((*MockProductRepository)(nil).FetchProductByID), ctx, id)
}

// FetchProductByIDForUpdate mocks base method.
func (m *MockProductRepository) FetchProductByIDForUpdate(ctx context.Context, id productdm.ProductID) (*productdm.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProductByIDForUpdate", ctx, id)
	ret0, _ := ret[0].(*productdm.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProductByIDForUpdate indicates an expected call of FetchProductByIDForUpdate.
func (mr *MockProductRepositoryMockRecorder) FetchProductByIDForUpdate(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProductByIDForUpdate", reflect.TypeOf((*MockProductRepository)(nil).FetchProductByIDForUpdate), ctx, id)
}

// UpdateProduct mocks base method.
func (m *MockProductRepository) UpdateProduct(ctx context.Context, product *productdm.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductRepositoryMockRecorder) UpdateProduct(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductRepository)(nil).UpdateProduct), ctx, product)
}