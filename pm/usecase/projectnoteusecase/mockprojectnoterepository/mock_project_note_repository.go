// Code generated by MockGen. DO NOT EDIT.
// Source: domain/projectnotedm/project_note_repository.go

// Package mockprojectnoterepository is a generated GoMock package.
package mockprojectnoterepository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	productdm "github.com/onituka/agile-project-management/project-management/domain/productdm"
	projectdm "github.com/onituka/agile-project-management/project-management/domain/projectdm"
	projectnotedm "github.com/onituka/agile-project-management/project-management/domain/projectnotedm"
)

// MockProjectNoteRepository is a mock of ProjectNoteRepository interface.
type MockProjectNoteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProjectNoteRepositoryMockRecorder
}

// MockProjectNoteRepositoryMockRecorder is the mock recorder for MockProjectNoteRepository.
type MockProjectNoteRepositoryMockRecorder struct {
	mock *MockProjectNoteRepository
}

// NewMockProjectNoteRepository creates a new mock instance.
func NewMockProjectNoteRepository(ctrl *gomock.Controller) *MockProjectNoteRepository {
	mock := &MockProjectNoteRepository{ctrl: ctrl}
	mock.recorder = &MockProjectNoteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectNoteRepository) EXPECT() *MockProjectNoteRepositoryMockRecorder {
	return m.recorder
}

// CreateProjectNote mocks base method.
func (m *MockProjectNoteRepository) CreateProjectNote(ctx context.Context, projectNote *projectnotedm.ProjectNote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProjectNote", ctx, projectNote)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProjectNote indicates an expected call of CreateProjectNote.
func (mr *MockProjectNoteRepositoryMockRecorder) CreateProjectNote(ctx, projectNote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProjectNote", reflect.TypeOf((*MockProjectNoteRepository)(nil).CreateProjectNote), ctx, projectNote)
}

// DeleteProjectNote mocks base method.
func (m *MockProjectNoteRepository) DeleteProjectNote(ctx context.Context, id projectnotedm.ProjectNoteID, productID productdm.ProductID, projectID projectdm.ProjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProjectNote", ctx, id, productID, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProjectNote indicates an expected call of DeleteProjectNote.
func (mr *MockProjectNoteRepositoryMockRecorder) DeleteProjectNote(ctx, id, productID, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectNote", reflect.TypeOf((*MockProjectNoteRepository)(nil).DeleteProjectNote), ctx, id, productID, projectID)
}

// FetchProjectNoteByID mocks base method.
func (m *MockProjectNoteRepository) FetchProjectNoteByID(ctx context.Context, id projectnotedm.ProjectNoteID, projectID projectdm.ProjectID) (*projectnotedm.ProjectNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectNoteByID", ctx, id, projectID)
	ret0, _ := ret[0].(*projectnotedm.ProjectNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectNoteByID indicates an expected call of FetchProjectNoteByID.
func (mr *MockProjectNoteRepositoryMockRecorder) FetchProjectNoteByID(ctx, id, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectNoteByID", reflect.TypeOf((*MockProjectNoteRepository)(nil).FetchProjectNoteByID), ctx, id, projectID)
}

// FetchProjectNoteByIDForUpdate mocks base method.
func (m *MockProjectNoteRepository) FetchProjectNoteByIDForUpdate(ctx context.Context, id projectnotedm.ProjectNoteID, projectID projectdm.ProjectID) (*projectnotedm.ProjectNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectNoteByIDForUpdate", ctx, id, projectID)
	ret0, _ := ret[0].(*projectnotedm.ProjectNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectNoteByIDForUpdate indicates an expected call of FetchProjectNoteByIDForUpdate.
func (mr *MockProjectNoteRepositoryMockRecorder) FetchProjectNoteByIDForUpdate(ctx, id, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectNoteByIDForUpdate", reflect.TypeOf((*MockProjectNoteRepository)(nil).FetchProjectNoteByIDForUpdate), ctx, id, projectID)
}

// FetchProjectNoteByProjectIDAndTitle mocks base method.
func (m *MockProjectNoteRepository) FetchProjectNoteByProjectIDAndTitle(ctx context.Context, projectID projectdm.ProjectID, title projectnotedm.Title) (*projectnotedm.ProjectNote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectNoteByProjectIDAndTitle", ctx, projectID, title)
	ret0, _ := ret[0].(*projectnotedm.ProjectNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectNoteByProjectIDAndTitle indicates an expected call of FetchProjectNoteByProjectIDAndTitle.
func (mr *MockProjectNoteRepositoryMockRecorder) FetchProjectNoteByProjectIDAndTitle(ctx, projectID, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectNoteByProjectIDAndTitle", reflect.TypeOf((*MockProjectNoteRepository)(nil).FetchProjectNoteByProjectIDAndTitle), ctx, projectID, title)
}

// UpdateProjectNote mocks base method.
func (m *MockProjectNoteRepository) UpdateProjectNote(ctx context.Context, projectNote *projectnotedm.ProjectNote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectNote", ctx, projectNote)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectNote indicates an expected call of UpdateProjectNote.
func (mr *MockProjectNoteRepositoryMockRecorder) UpdateProjectNote(ctx, projectNote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectNote", reflect.TypeOf((*MockProjectNoteRepository)(nil).UpdateProjectNote), ctx, projectNote)
}
