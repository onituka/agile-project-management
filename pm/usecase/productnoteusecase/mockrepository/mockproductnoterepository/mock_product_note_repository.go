// Code generated by MockGen. DO NOT EDIT.
// Source: domain/productnotedm/product_note_repository.go

// Package mockproductnoterepository is a generated GoMock package.
package mockproductnoterepository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	productdm "github.com/onituka/agile-project-management/project-management/domain/productdm"
	productnotedm "github.com/onituka/agile-project-management/project-management/domain/productnotedm"
)

// MockProductNoteRepository is a mock of ProductNoteRepository interface.
type MockProductNoteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductNoteRepositoryMockRecorder
}

// MockProductNoteRepositoryMockRecorder is the mock recorder for MockProductNoteRepository.
type MockProductNoteRepositoryMockRecorder struct {
	mock *MockProductNoteRepository
}

// NewMockProductNoteRepository creates a new mock instance.
func NewMockProductNoteRepository(ctrl *gomock.Controller) *MockProductNoteRepository {
	mock := &MockProductNoteRepository{ctrl: ctrl}
	mock.recorder = &MockProductNoteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductNoteRepository) EXPECT() *MockProductNoteRepositoryMockRecorder {
	return m.recorder
}

// CreateProductNote mocks base method.
func (m *MockProductNoteRepository) CreateProductNote(ctx context.Context, productNote *productnotedm.ProductNote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductNote", ctx, productNote)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProductNote indicates an expected call of CreateProductNote.
func (mr *MockProductNoteRepositoryMockRecorder) CreateProductNote(ctx, productNote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductNote", reflect.TypeOf((*MockProductNoteRepository)(nil).CreateProductNote), ctx, productNote)
}

// DeleteProductNote mocks base method.
func (m *MockProductNoteRepository) DeleteProductNote(ctx context.Context, id productnotedm.ProductNoteID, productID productdm.ProductID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductNote", ctx, id, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductNote indicates an expected call of DeleteProductNote.
func (mr *MockProductNoteRepositoryMockRecorder) DeleteProductNote(ctx, id, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductNote", reflect.TypeOf((*MockProductNoteRepository)(nil).DeleteProductNote), ctx, id, productID)
}

// FetchProductNoteByID mocks base method.
func (m *MockProductNoteRepository) FetchProductNoteByID(ctx context.Context, id productnotedm.ProductNoteID, productID productdm.ProductID) (*productnotedm.ProductNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProductNoteByID", ctx, id, productID)
	ret0, _ := ret[0].(*productnotedm.ProductNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProductNoteByID indicates an expected call of FetchProductNoteByID.
func (mr *MockProductNoteRepositoryMockRecorder) FetchProductNoteByID(ctx, id, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProductNoteByID", reflect.TypeOf((*MockProductNoteRepository)(nil).FetchProductNoteByID), ctx, id, productID)
}

// FetchProductNoteByIDForUpdate mocks base method.
func (m *MockProductNoteRepository) FetchProductNoteByIDForUpdate(ctx context.Context, id productnotedm.ProductNoteID, productID productdm.ProductID) (*productnotedm.ProductNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProductNoteByIDForUpdate", ctx, id, productID)
	ret0, _ := ret[0].(*productnotedm.ProductNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProductNoteByIDForUpdate indicates an expected call of FetchProductNoteByIDForUpdate.
func (mr *MockProductNoteRepositoryMockRecorder) FetchProductNoteByIDForUpdate(ctx, id, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProductNoteByIDForUpdate", reflect.TypeOf((*MockProductNoteRepository)(nil).FetchProductNoteByIDForUpdate), ctx, id, productID)
}

// FetchProductNoteByProductIDAndTitle mocks base method.
func (m *MockProductNoteRepository) FetchProductNoteByProductIDAndTitle(ctx context.Context, productID productdm.ProductID, Title productnotedm.Title) (*productnotedm.ProductNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProductNoteByProductIDAndTitle", ctx, productID, Title)
	ret0, _ := ret[0].(*productnotedm.ProductNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProductNoteByProductIDAndTitle indicates an expected call of FetchProductNoteByProductIDAndTitle.
func (mr *MockProductNoteRepositoryMockRecorder) FetchProductNoteByProductIDAndTitle(ctx, productID, Title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProductNoteByProductIDAndTitle", reflect.TypeOf((*MockProductNoteRepository)(nil).FetchProductNoteByProductIDAndTitle), ctx, productID, Title)
}

// UpdateProductNote mocks base method.
func (m *MockProductNoteRepository) UpdateProductNote(ctx context.Context, productNote *productnotedm.ProductNote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductNote", ctx, productNote)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductNote indicates an expected call of UpdateProductNote.
func (mr *MockProductNoteRepositoryMockRecorder) UpdateProductNote(ctx, productNote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductNote", reflect.TypeOf((*MockProductNoteRepository)(nil).UpdateProductNote), ctx, productNote)
}
