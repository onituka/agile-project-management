// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/projectusecase/projectqueryservice/project_query_service.go

// Package mockprojectqueryservice is a generated GoMock package.
package mockprojectqueryservice

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	productdm "github.com/onituka/agile-project-management/project-management/domain/productdm"
	projectoutput "github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

// MockProjectQueryService is a mock of ProjectQueryService interface.
type MockProjectQueryService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectQueryServiceMockRecorder
}

// MockProjectQueryServiceMockRecorder is the mock recorder for MockProjectQueryService.
type MockProjectQueryServiceMockRecorder struct {
	mock *MockProjectQueryService
}

// NewMockProjectQueryService creates a new mock instance.
func NewMockProjectQueryService(ctrl *gomock.Controller) *MockProjectQueryService {
	mock := &MockProjectQueryService{ctrl: ctrl}
	mock.recorder = &MockProjectQueryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectQueryService) EXPECT() *MockProjectQueryServiceMockRecorder {
	return m.recorder
}

// CountProjectsByKeyNameAndName mocks base method.
func (m *MockProjectQueryService) CountProjectsByKeyNameAndName(ctx context.Context, productID productdm.ProductID, keyword string) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProjectsByKeyNameAndName", ctx, productID, keyword)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProjectsByKeyNameAndName indicates an expected call of CountProjectsByKeyNameAndName.
func (mr *MockProjectQueryServiceMockRecorder) CountProjectsByKeyNameAndName(ctx, productID, keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProjectsByKeyNameAndName", reflect.TypeOf((*MockProjectQueryService)(nil).CountProjectsByKeyNameAndName), ctx, productID, keyword)
}

// CountProjectsByProductID mocks base method.
func (m *MockProjectQueryService) CountProjectsByProductID(ctx context.Context, productID productdm.ProductID) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProjectsByProductID", ctx, productID)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProjectsByProductID indicates an expected call of CountProjectsByProductID.
func (mr *MockProjectQueryServiceMockRecorder) CountProjectsByProductID(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProjectsByProductID", reflect.TypeOf((*MockProjectQueryService)(nil).CountProjectsByProductID), ctx, productID)
}

// CountTrashedProjectsByProductID mocks base method.
func (m *MockProjectQueryService) CountTrashedProjectsByProductID(ctx context.Context, productID productdm.ProductID) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountTrashedProjectsByProductID", ctx, productID)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountTrashedProjectsByProductID indicates an expected call of CountTrashedProjectsByProductID.
func (mr *MockProjectQueryServiceMockRecorder) CountTrashedProjectsByProductID(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountTrashedProjectsByProductID", reflect.TypeOf((*MockProjectQueryService)(nil).CountTrashedProjectsByProductID), ctx, productID)
}

// FetchProjects mocks base method.
func (m *MockProjectQueryService) FetchProjects(ctx context.Context, productID productdm.ProductID, limit, offset uint32) ([]*projectoutput.ProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjects", ctx, productID, limit, offset)
	ret0, _ := ret[0].([]*projectoutput.ProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjects indicates an expected call of FetchProjects.
func (mr *MockProjectQueryServiceMockRecorder) FetchProjects(ctx, productID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjects", reflect.TypeOf((*MockProjectQueryService)(nil).FetchProjects), ctx, productID, limit, offset)
}

// FetchTrashedProjects mocks base method.
func (m *MockProjectQueryService) FetchTrashedProjects(ctx context.Context, productID productdm.ProductID, limit, offset uint32) ([]*projectoutput.FetchTrashedProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTrashedProjects", ctx, productID, limit, offset)
	ret0, _ := ret[0].([]*projectoutput.FetchTrashedProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTrashedProjects indicates an expected call of FetchTrashedProjects.
func (mr *MockProjectQueryServiceMockRecorder) FetchTrashedProjects(ctx, productID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTrashedProjects", reflect.TypeOf((*MockProjectQueryService)(nil).FetchTrashedProjects), ctx, productID, limit, offset)
}

// SearchProjects mocks base method.
func (m *MockProjectQueryService) SearchProjects(ctx context.Context, productID productdm.ProductID, keyword string, limit, offset uint32) ([]*projectoutput.SearchProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProjects", ctx, productID, keyword, limit, offset)
	ret0, _ := ret[0].([]*projectoutput.SearchProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProjects indicates an expected call of SearchProjects.
func (mr *MockProjectQueryServiceMockRecorder) SearchProjects(ctx, productID, keyword, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProjects", reflect.TypeOf((*MockProjectQueryService)(nil).SearchProjects), ctx, productID, keyword, limit, offset)
}
