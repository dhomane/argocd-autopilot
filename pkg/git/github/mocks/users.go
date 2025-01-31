// Code generated by MockGen. DO NOT EDIT.
// Source: ./github/users.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v35/github"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// AcceptInvitation mocks base method.
func (m *MockUsers) AcceptInvitation(arg0 context.Context, arg1 int64) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptInvitation", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptInvitation indicates an expected call of AcceptInvitation.
func (mr *MockUsersMockRecorder) AcceptInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptInvitation", reflect.TypeOf((*MockUsers)(nil).AcceptInvitation), arg0, arg1)
}

// AddEmails mocks base method.
func (m *MockUsers) AddEmails(arg0 context.Context, arg1 []string) ([]*github.UserEmail, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEmails", arg0, arg1)
	ret0, _ := ret[0].([]*github.UserEmail)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddEmails indicates an expected call of AddEmails.
func (mr *MockUsersMockRecorder) AddEmails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEmails", reflect.TypeOf((*MockUsers)(nil).AddEmails), arg0, arg1)
}

// BlockUser mocks base method.
func (m *MockUsers) BlockUser(arg0 context.Context, arg1 string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockUser", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockUser indicates an expected call of BlockUser.
func (mr *MockUsersMockRecorder) BlockUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockUser", reflect.TypeOf((*MockUsers)(nil).BlockUser), arg0, arg1)
}

// CreateGPGKey mocks base method.
func (m *MockUsers) CreateGPGKey(arg0 context.Context, arg1 string) (*github.GPGKey, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGPGKey", arg0, arg1)
	ret0, _ := ret[0].(*github.GPGKey)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateGPGKey indicates an expected call of CreateGPGKey.
func (mr *MockUsersMockRecorder) CreateGPGKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGPGKey", reflect.TypeOf((*MockUsers)(nil).CreateGPGKey), arg0, arg1)
}

// CreateKey mocks base method.
func (m *MockUsers) CreateKey(arg0 context.Context, arg1 *github.Key) (*github.Key, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", arg0, arg1)
	ret0, _ := ret[0].(*github.Key)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockUsersMockRecorder) CreateKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockUsers)(nil).CreateKey), arg0, arg1)
}

// CreateProject mocks base method.
func (m *MockUsers) CreateProject(arg0 context.Context, arg1 *github.CreateUserProjectOptions) (*github.Project, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(*github.Project)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockUsersMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockUsers)(nil).CreateProject), arg0, arg1)
}

// DeclineInvitation mocks base method.
func (m *MockUsers) DeclineInvitation(arg0 context.Context, arg1 int64) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclineInvitation", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeclineInvitation indicates an expected call of DeclineInvitation.
func (mr *MockUsersMockRecorder) DeclineInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclineInvitation", reflect.TypeOf((*MockUsers)(nil).DeclineInvitation), arg0, arg1)
}

// DeleteEmails mocks base method.
func (m *MockUsers) DeleteEmails(arg0 context.Context, arg1 []string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmails", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEmails indicates an expected call of DeleteEmails.
func (mr *MockUsersMockRecorder) DeleteEmails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmails", reflect.TypeOf((*MockUsers)(nil).DeleteEmails), arg0, arg1)
}

// DeleteGPGKey mocks base method.
func (m *MockUsers) DeleteGPGKey(arg0 context.Context, arg1 int64) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGPGKey", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGPGKey indicates an expected call of DeleteGPGKey.
func (mr *MockUsersMockRecorder) DeleteGPGKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGPGKey", reflect.TypeOf((*MockUsers)(nil).DeleteGPGKey), arg0, arg1)
}

// DeleteKey mocks base method.
func (m *MockUsers) DeleteKey(arg0 context.Context, arg1 int64) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKey", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKey indicates an expected call of DeleteKey.
func (mr *MockUsersMockRecorder) DeleteKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockUsers)(nil).DeleteKey), arg0, arg1)
}

// DemoteSiteAdmin mocks base method.
func (m *MockUsers) DemoteSiteAdmin(arg0 context.Context, arg1 string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DemoteSiteAdmin", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DemoteSiteAdmin indicates an expected call of DemoteSiteAdmin.
func (mr *MockUsersMockRecorder) DemoteSiteAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DemoteSiteAdmin", reflect.TypeOf((*MockUsers)(nil).DemoteSiteAdmin), arg0, arg1)
}

// Edit mocks base method.
func (m *MockUsers) Edit(arg0 context.Context, arg1 *github.User) (*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", arg0, arg1)
	ret0, _ := ret[0].(*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Edit indicates an expected call of Edit.
func (mr *MockUsersMockRecorder) Edit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockUsers)(nil).Edit), arg0, arg1)
}

// Follow mocks base method.
func (m *MockUsers) Follow(arg0 context.Context, arg1 string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Follow indicates an expected call of Follow.
func (mr *MockUsersMockRecorder) Follow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockUsers)(nil).Follow), arg0, arg1)
}

// Get mocks base method.
func (m *MockUsers) Get(arg0 context.Context, arg1 string) (*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockUsersMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsers)(nil).Get), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockUsers) GetByID(arg0 context.Context, arg1 int64) (*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUsersMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUsers)(nil).GetByID), arg0, arg1)
}

// GetGPGKey mocks base method.
func (m *MockUsers) GetGPGKey(arg0 context.Context, arg1 int64) (*github.GPGKey, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGPGKey", arg0, arg1)
	ret0, _ := ret[0].(*github.GPGKey)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGPGKey indicates an expected call of GetGPGKey.
func (mr *MockUsersMockRecorder) GetGPGKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGPGKey", reflect.TypeOf((*MockUsers)(nil).GetGPGKey), arg0, arg1)
}

// GetHovercard mocks base method.
func (m *MockUsers) GetHovercard(arg0 context.Context, arg1 string, arg2 *github.HovercardOptions) (*github.Hovercard, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHovercard", arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.Hovercard)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHovercard indicates an expected call of GetHovercard.
func (mr *MockUsersMockRecorder) GetHovercard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHovercard", reflect.TypeOf((*MockUsers)(nil).GetHovercard), arg0, arg1, arg2)
}

// GetKey mocks base method.
func (m *MockUsers) GetKey(arg0 context.Context, arg1 int64) (*github.Key, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", arg0, arg1)
	ret0, _ := ret[0].(*github.Key)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKey indicates an expected call of GetKey.
func (mr *MockUsersMockRecorder) GetKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockUsers)(nil).GetKey), arg0, arg1)
}

// IsBlocked mocks base method.
func (m *MockUsers) IsBlocked(arg0 context.Context, arg1 string) (bool, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBlocked", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsBlocked indicates an expected call of IsBlocked.
func (mr *MockUsersMockRecorder) IsBlocked(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBlocked", reflect.TypeOf((*MockUsers)(nil).IsBlocked), arg0, arg1)
}

// IsFollowing mocks base method.
func (m *MockUsers) IsFollowing(arg0 context.Context, arg1, arg2 string) (bool, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFollowing", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsFollowing indicates an expected call of IsFollowing.
func (mr *MockUsersMockRecorder) IsFollowing(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFollowing", reflect.TypeOf((*MockUsers)(nil).IsFollowing), arg0, arg1, arg2)
}

// ListAll mocks base method.
func (m *MockUsers) ListAll(arg0 context.Context, arg1 *github.UserListOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0, arg1)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAll indicates an expected call of ListAll.
func (mr *MockUsersMockRecorder) ListAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockUsers)(nil).ListAll), arg0, arg1)
}

// ListBlockedUsers mocks base method.
func (m *MockUsers) ListBlockedUsers(arg0 context.Context, arg1 *github.ListOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBlockedUsers", arg0, arg1)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBlockedUsers indicates an expected call of ListBlockedUsers.
func (mr *MockUsersMockRecorder) ListBlockedUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBlockedUsers", reflect.TypeOf((*MockUsers)(nil).ListBlockedUsers), arg0, arg1)
}

// ListEmails mocks base method.
func (m *MockUsers) ListEmails(arg0 context.Context, arg1 *github.ListOptions) ([]*github.UserEmail, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEmails", arg0, arg1)
	ret0, _ := ret[0].([]*github.UserEmail)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListEmails indicates an expected call of ListEmails.
func (mr *MockUsersMockRecorder) ListEmails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEmails", reflect.TypeOf((*MockUsers)(nil).ListEmails), arg0, arg1)
}

// ListFollowers mocks base method.
func (m *MockUsers) ListFollowers(arg0 context.Context, arg1 string, arg2 *github.ListOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollowers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFollowers indicates an expected call of ListFollowers.
func (mr *MockUsersMockRecorder) ListFollowers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollowers", reflect.TypeOf((*MockUsers)(nil).ListFollowers), arg0, arg1, arg2)
}

// ListFollowing mocks base method.
func (m *MockUsers) ListFollowing(arg0 context.Context, arg1 string, arg2 *github.ListOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFollowing", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFollowing indicates an expected call of ListFollowing.
func (mr *MockUsersMockRecorder) ListFollowing(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFollowing", reflect.TypeOf((*MockUsers)(nil).ListFollowing), arg0, arg1, arg2)
}

// ListGPGKeys mocks base method.
func (m *MockUsers) ListGPGKeys(arg0 context.Context, arg1 string, arg2 *github.ListOptions) ([]*github.GPGKey, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGPGKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.GPGKey)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGPGKeys indicates an expected call of ListGPGKeys.
func (mr *MockUsersMockRecorder) ListGPGKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGPGKeys", reflect.TypeOf((*MockUsers)(nil).ListGPGKeys), arg0, arg1, arg2)
}

// ListInvitations mocks base method.
func (m *MockUsers) ListInvitations(arg0 context.Context, arg1 *github.ListOptions) ([]*github.RepositoryInvitation, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInvitations", arg0, arg1)
	ret0, _ := ret[0].([]*github.RepositoryInvitation)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListInvitations indicates an expected call of ListInvitations.
func (mr *MockUsersMockRecorder) ListInvitations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvitations", reflect.TypeOf((*MockUsers)(nil).ListInvitations), arg0, arg1)
}

// ListKeys mocks base method.
func (m *MockUsers) ListKeys(arg0 context.Context, arg1 string, arg2 *github.ListOptions) ([]*github.Key, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Key)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockUsersMockRecorder) ListKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockUsers)(nil).ListKeys), arg0, arg1, arg2)
}

// ListProjects mocks base method.
func (m *MockUsers) ListProjects(arg0 context.Context, arg1 string, arg2 *github.ProjectListOptions) ([]*github.Project, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Project)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockUsersMockRecorder) ListProjects(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockUsers)(nil).ListProjects), arg0, arg1, arg2)
}

// PromoteSiteAdmin mocks base method.
func (m *MockUsers) PromoteSiteAdmin(arg0 context.Context, arg1 string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PromoteSiteAdmin", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PromoteSiteAdmin indicates an expected call of PromoteSiteAdmin.
func (mr *MockUsersMockRecorder) PromoteSiteAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PromoteSiteAdmin", reflect.TypeOf((*MockUsers)(nil).PromoteSiteAdmin), arg0, arg1)
}

// Suspend mocks base method.
func (m *MockUsers) Suspend(arg0 context.Context, arg1 string, arg2 *github.UserSuspendOptions) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suspend", arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Suspend indicates an expected call of Suspend.
func (mr *MockUsersMockRecorder) Suspend(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suspend", reflect.TypeOf((*MockUsers)(nil).Suspend), arg0, arg1, arg2)
}

// UnblockUser mocks base method.
func (m *MockUsers) UnblockUser(arg0 context.Context, arg1 string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnblockUser", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnblockUser indicates an expected call of UnblockUser.
func (mr *MockUsersMockRecorder) UnblockUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnblockUser", reflect.TypeOf((*MockUsers)(nil).UnblockUser), arg0, arg1)
}

// Unfollow mocks base method.
func (m *MockUsers) Unfollow(arg0 context.Context, arg1 string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unfollow indicates an expected call of Unfollow.
func (mr *MockUsersMockRecorder) Unfollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockUsers)(nil).Unfollow), arg0, arg1)
}

// Unsuspend mocks base method.
func (m *MockUsers) Unsuspend(arg0 context.Context, arg1 string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsuspend", arg0, arg1)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unsuspend indicates an expected call of Unsuspend.
func (mr *MockUsersMockRecorder) Unsuspend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsuspend", reflect.TypeOf((*MockUsers)(nil).Unsuspend), arg0, arg1)
}
