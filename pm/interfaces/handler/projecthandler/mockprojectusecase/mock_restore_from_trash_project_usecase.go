// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/projectusecase/restore_from_trash_project_usecase.go

// Package mockprojectusecase is a generated GoMock package.
package mockprojectusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	projectinput "github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	projectoutput "github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

// MockRestoreFromTrashProjectUsecase is a mock of RestoreFromTrashProjectUsecase interface.
type MockRestoreFromTrashProjectUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockRestoreFromTrashProjectUsecaseMockRecorder
}

// MockRestoreFromTrashProjectUsecaseMockRecorder is the mock recorder for MockRestoreFromTrashProjectUsecase.
type MockRestoreFromTrashProjectUsecaseMockRecorder struct {
	mock *MockRestoreFromTrashProjectUsecase
}

// NewMockRestoreFromTrashProjectUsecase creates a new mock instance.
func NewMockRestoreFromTrashProjectUsecase(ctrl *gomock.Controller) *MockRestoreFromTrashProjectUsecase {
	mock := &MockRestoreFromTrashProjectUsecase{ctrl: ctrl}
	mock.recorder = &MockRestoreFromTrashProjectUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestoreFromTrashProjectUsecase) EXPECT() *MockRestoreFromTrashProjectUsecaseMockRecorder {
	return m.recorder
}

// RestoreFromTrashProject mocks base method.
func (m *MockRestoreFromTrashProjectUsecase) RestoreFromTrashProject(ctx context.Context, in *projectinput.RestoreFromTrashProjectIDInput) (*projectoutput.RestoreFromTrashProjectIDOutPut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreFromTrashProject", ctx, in)
	ret0, _ := ret[0].(*projectoutput.RestoreFromTrashProjectIDOutPut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreFromTrashProject indicates an expected call of RestoreFromTrashProject.
func (mr *MockRestoreFromTrashProjectUsecaseMockRecorder) RestoreFromTrashProject(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreFromTrashProject", reflect.TypeOf((*MockRestoreFromTrashProjectUsecase)(nil).RestoreFromTrashProject), ctx, in)
}