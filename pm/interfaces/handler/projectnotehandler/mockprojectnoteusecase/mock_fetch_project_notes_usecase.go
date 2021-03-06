// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/projectnoteusecase/fetch_project_notes_usecase.go

// Package mockprojectnoteusecase is a generated GoMock package.
package mockprojectnoteusecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	projectnoteinput "github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
	projectnoteoutput "github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
)

// MockFetchProjectNotesUsecase is a mock of FetchProjectNotesUsecase interface.
type MockFetchProjectNotesUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockFetchProjectNotesUsecaseMockRecorder
}

// MockFetchProjectNotesUsecaseMockRecorder is the mock recorder for MockFetchProjectNotesUsecase.
type MockFetchProjectNotesUsecaseMockRecorder struct {
	mock *MockFetchProjectNotesUsecase
}

// NewMockFetchProjectNotesUsecase creates a new mock instance.
func NewMockFetchProjectNotesUsecase(ctrl *gomock.Controller) *MockFetchProjectNotesUsecase {
	mock := &MockFetchProjectNotesUsecase{ctrl: ctrl}
	mock.recorder = &MockFetchProjectNotesUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFetchProjectNotesUsecase) EXPECT() *MockFetchProjectNotesUsecaseMockRecorder {
	return m.recorder
}

// FetchProjectNotes mocks base method.
func (m *MockFetchProjectNotesUsecase) FetchProjectNotes(ctx context.Context, in *projectnoteinput.FetchProjectNotesInput) (*projectnoteoutput.FetchProjectNotesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectNotes", ctx, in)
	ret0, _ := ret[0].(*projectnoteoutput.FetchProjectNotesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectNotes indicates an expected call of FetchProjectNotes.
func (mr *MockFetchProjectNotesUsecaseMockRecorder) FetchProjectNotes(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectNotes", reflect.TypeOf((*MockFetchProjectNotesUsecase)(nil).FetchProjectNotes), ctx, in)
}
