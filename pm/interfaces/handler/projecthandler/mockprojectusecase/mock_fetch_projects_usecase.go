// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/projectusecase/fetch_projects_usecase.go

// Package mockprojectusecase is a generated GoMock package.
package mockprojectusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	projectinput "github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	projectoutput "github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

// MockFetchProjectsUsecase is a mock of FetchProjectsUsecase interface.
type MockFetchProjectsUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockFetchProjectsUsecaseMockRecorder
}

// MockFetchProjectsUsecaseMockRecorder is the mock recorder for MockFetchProjectsUsecase.
type MockFetchProjectsUsecaseMockRecorder struct {
	mock *MockFetchProjectsUsecase
}

// NewMockFetchProjectsUsecase creates a new mock instance.
func NewMockFetchProjectsUsecase(ctrl *gomock.Controller) *MockFetchProjectsUsecase {
	mock := &MockFetchProjectsUsecase{ctrl: ctrl}
	mock.recorder = &MockFetchProjectsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFetchProjectsUsecase) EXPECT() *MockFetchProjectsUsecaseMockRecorder {
	return m.recorder
}

// FetchProjects mocks base method.
func (m *MockFetchProjectsUsecase) FetchProjects(ctx context.Context, in *projectinput.FetchProjectsInput) (*projectoutput.FetchProjectsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjects", ctx, in)
	ret0, _ := ret[0].(*projectoutput.FetchProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjects indicates an expected call of FetchProjects.
func (mr *MockFetchProjectsUsecaseMockRecorder) FetchProjects(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjects", reflect.TypeOf((*MockFetchProjectsUsecase)(nil).FetchProjects), ctx, in)
}
