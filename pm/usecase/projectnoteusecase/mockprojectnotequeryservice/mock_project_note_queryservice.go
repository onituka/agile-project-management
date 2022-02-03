// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/projectnoteusecase/projectnotequeryservice/project_note_query_service.go

// Package mockprojectnotequeryservice is a generated GoMock package.
package mockprojectnotequeryservice

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	productdm "github.com/onituka/agile-project-management/project-management/domain/productdm"
	projectdm "github.com/onituka/agile-project-management/project-management/domain/projectdm"
	projectnoteoutput "github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
)

// MockProjectNoteQueryService is a mock of ProjectNoteQueryService interface.
type MockProjectNoteQueryService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectNoteQueryServiceMockRecorder
}

// MockProjectNoteQueryServiceMockRecorder is the mock recorder for MockProjectNoteQueryService.
type MockProjectNoteQueryServiceMockRecorder struct {
	mock *MockProjectNoteQueryService
}

// NewMockProjectNoteQueryService creates a new mock instance.
func NewMockProjectNoteQueryService(ctrl *gomock.Controller) *MockProjectNoteQueryService {
	mock := &MockProjectNoteQueryService{ctrl: ctrl}
	mock.recorder = &MockProjectNoteQueryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectNoteQueryService) EXPECT() *MockProjectNoteQueryServiceMockRecorder {
	return m.recorder
}

// CountProjectNotesByProductIDAndProjectID mocks base method.
func (m *MockProjectNoteQueryService) CountProjectNotesByProductIDAndProjectID(ctx context.Context, productID productdm.ProductID, projectID projectdm.ProjectID) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProjectNotesByProductIDAndProjectID", ctx, productID, projectID)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProjectNotesByProductIDAndProjectID indicates an expected call of CountProjectNotesByProductIDAndProjectID.
func (mr *MockProjectNoteQueryServiceMockRecorder) CountProjectNotesByProductIDAndProjectID(ctx, productID, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProjectNotesByProductIDAndProjectID", reflect.TypeOf((*MockProjectNoteQueryService)(nil).CountProjectNotesByProductIDAndProjectID), ctx, productID, projectID)
}

// FetchProjectNotes mocks base method.
func (m *MockProjectNoteQueryService) FetchProjectNotes(ctx context.Context, productID productdm.ProductID, projectID projectdm.ProjectID, limit, offset uint32) ([]*projectnoteoutput.ProjectNoteOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectNotes", ctx, productID, projectID, limit, offset)
	ret0, _ := ret[0].([]*projectnoteoutput.ProjectNoteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectNotes indicates an expected call of FetchProjectNotes.
func (mr *MockProjectNoteQueryServiceMockRecorder) FetchProjectNotes(ctx, productID, projectID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectNotes", reflect.TypeOf((*MockProjectNoteQueryService)(nil).FetchProjectNotes), ctx, productID, projectID, limit, offset)
}