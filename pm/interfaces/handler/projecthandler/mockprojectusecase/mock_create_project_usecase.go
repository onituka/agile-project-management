// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/projectusecase/create_project_usecase.go

// Package mockprojectusecase is a generated GoMock package.
package mockprojectusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	projectusecase "github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

// MockCreateProjectUsecase is a mock of CreateProjectUsecase interface.
type MockCreateProjectUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateProjectUsecaseMockRecorder
}

// MockCreateProjectUsecaseMockRecorder is the mock recorder for MockCreateProjectUsecase.
type MockCreateProjectUsecaseMockRecorder struct {
	mock *MockCreateProjectUsecase
}

// NewMockCreateProjectUsecase creates a new mock instance.
func NewMockCreateProjectUsecase(ctrl *gomock.Controller) *MockCreateProjectUsecase {
	mock := &MockCreateProjectUsecase{ctrl: ctrl}
	mock.recorder = &MockCreateProjectUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateProjectUsecase) EXPECT() *MockCreateProjectUsecaseMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockCreateProjectUsecase) CreateProject(ctx context.Context, in *projectusecase.CreateProjectInput) (*projectusecase.CreateProjectOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, in)
	ret0, _ := ret[0].(*projectusecase.CreateProjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockCreateProjectUsecaseMockRecorder) CreateProject(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockCreateProjectUsecase)(nil).CreateProject), ctx, in)
}
