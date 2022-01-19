// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/projectusecase/search_projects_usecase.go

// Package mockprojectusecase is a generated GoMock package.
package mockprojectusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	projectinput "github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	projectoutput "github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

// MockSearchProjectsUsecase is a mock of SearchProjectsUsecase interface.
type MockSearchProjectsUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockSearchProjectsUsecaseMockRecorder
}

// MockSearchProjectsUsecaseMockRecorder is the mock recorder for MockSearchProjectsUsecase.
type MockSearchProjectsUsecaseMockRecorder struct {
	mock *MockSearchProjectsUsecase
}

// NewMockSearchProjectsUsecase creates a new mock instance.
func NewMockSearchProjectsUsecase(ctrl *gomock.Controller) *MockSearchProjectsUsecase {
	mock := &MockSearchProjectsUsecase{ctrl: ctrl}
	mock.recorder = &MockSearchProjectsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchProjectsUsecase) EXPECT() *MockSearchProjectsUsecaseMockRecorder {
	return m.recorder
}

// SearchProjects mocks base method.
func (m *MockSearchProjectsUsecase) SearchProjects(ctx context.Context, in *projectinput.SearchProjectsInput) (*projectoutput.SearchProjectsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProjects", ctx, in)
	ret0, _ := ret[0].(*projectoutput.SearchProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProjects indicates an expected call of SearchProjects.
func (mr *MockSearchProjectsUsecaseMockRecorder) SearchProjects(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProjects", reflect.TypeOf((*MockSearchProjectsUsecase)(nil).SearchProjects), ctx, in)
}
