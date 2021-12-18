// Code generated by MockGen. DO NOT EDIT.
// Source: domain/projectdm/project_repository.go

// Package mockprojectrepository is a generated GoMock package.
package mockprojectrepository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	groupdm "github.com/onituka/agile-project-management/project-management/domain/groupdm"
	projectdm "github.com/onituka/agile-project-management/project-management/domain/projectdm"
)

// MockProjectRepository is a mock of ProjectRepository interface.
type MockProjectRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProjectRepositoryMockRecorder
}

// MockProjectRepositoryMockRecorder is the mock recorder for MockProjectRepository.
type MockProjectRepositoryMockRecorder struct {
	mock *MockProjectRepository
}

// NewMockProjectRepository creates a new mock instance.
func NewMockProjectRepository(ctrl *gomock.Controller) *MockProjectRepository {
	mock := &MockProjectRepository{ctrl: ctrl}
	mock.recorder = &MockProjectRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectRepository) EXPECT() *MockProjectRepositoryMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockProjectRepository) CreateProject(ctx context.Context, project *projectdm.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, project)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectRepositoryMockRecorder) CreateProject(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectRepository)(nil).CreateProject), ctx, project)
}

// FetchProjectByGroupIDAndKeyName mocks base method.
func (m *MockProjectRepository) FetchProjectByGroupIDAndKeyName(ctx context.Context, groupID groupdm.GroupID, keyName projectdm.KeyName) (*projectdm.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectByGroupIDAndKeyName", ctx, groupID, keyName)
	ret0, _ := ret[0].(*projectdm.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectByGroupIDAndKeyName indicates an expected call of FetchProjectByGroupIDAndKeyName.
func (mr *MockProjectRepositoryMockRecorder) FetchProjectByGroupIDAndKeyName(ctx, groupID, keyName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectByGroupIDAndKeyName", reflect.TypeOf((*MockProjectRepository)(nil).FetchProjectByGroupIDAndKeyName), ctx, groupID, keyName)
}

// FetchProjectByGroupIDAndName mocks base method.
func (m *MockProjectRepository) FetchProjectByGroupIDAndName(ctx context.Context, groupID groupdm.GroupID, name projectdm.Name) (*projectdm.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectByGroupIDAndName", ctx, groupID, name)
	ret0, _ := ret[0].(*projectdm.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectByGroupIDAndName indicates an expected call of FetchProjectByGroupIDAndName.
func (mr *MockProjectRepositoryMockRecorder) FetchProjectByGroupIDAndName(ctx, groupID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectByGroupIDAndName", reflect.TypeOf((*MockProjectRepository)(nil).FetchProjectByGroupIDAndName), ctx, groupID, name)
}

// FetchProjectByID mocks base method.
func (m *MockProjectRepository) FetchProjectByID(ctx context.Context, id projectdm.ProjectID) (*projectdm.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectByID", ctx, id)
	ret0, _ := ret[0].(*projectdm.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectByID indicates an expected call of FetchProjectByID.
func (mr *MockProjectRepositoryMockRecorder) FetchProjectByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectByID", reflect.TypeOf((*MockProjectRepository)(nil).FetchProjectByID), ctx, id)
}

// FetchProjectByIDForUpdate mocks base method.
func (m *MockProjectRepository) FetchProjectByIDForUpdate(ctx context.Context, id projectdm.ProjectID) (*projectdm.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjectByIDForUpdate", ctx, id)
	ret0, _ := ret[0].(*projectdm.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjectByIDForUpdate indicates an expected call of FetchProjectByIDForUpdate.
func (mr *MockProjectRepositoryMockRecorder) FetchProjectByIDForUpdate(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjectByIDForUpdate", reflect.TypeOf((*MockProjectRepository)(nil).FetchProjectByIDForUpdate), ctx, id)
}

// FetchProjects mocks base method.
func (m *MockProjectRepository) FetchProjects(ctx context.Context) ([]*projectdm.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProjects", ctx)
	ret0, _ := ret[0].([]*projectdm.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProjects indicates an expected call of FetchProjects.
func (mr *MockProjectRepositoryMockRecorder) FetchProjects(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProjects", reflect.TypeOf((*MockProjectRepository)(nil).FetchProjects), ctx)
}

// UpdateProject mocks base method.
func (m *MockProjectRepository) UpdateProject(ctx context.Context, project *projectdm.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", ctx, project)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockProjectRepositoryMockRecorder) UpdateProject(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectRepository)(nil).UpdateProject), ctx, project)
}
