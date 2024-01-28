// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	models "github.com/RyanTrue/GophKeeper/internal/models"
	gomock "github.com/golang/mock/gomock"
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

// Create mocks base method.
func (m *MockUsers) Create(ctx context.Context, login, password, aesSecret, privateKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, login, password, aesSecret, privateKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUsersMockRecorder) Create(ctx, login, password, aesSecret, privateKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsers)(nil).Create), ctx, login, password, aesSecret, privateKey)
}

// FindByLogin mocks base method.
func (m *MockUsers) FindByLogin(ctx context.Context, login string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLogin", ctx, login)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLogin indicates an expected call of FindByLogin.
func (mr *MockUsersMockRecorder) FindByLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*MockUsers)(nil).FindByLogin), ctx, login)
}

// MockSettings is a mock of Settings interface.
type MockSettings struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsMockRecorder
}

// MockSettingsMockRecorder is the mock recorder for MockSettings.
type MockSettingsMockRecorder struct {
	mock *MockSettings
}

// NewMockSettings creates a new mock instance.
func NewMockSettings(ctrl *gomock.Controller) *MockSettings {
	mock := &MockSettings{ctrl: ctrl}
	mock.recorder = &MockSettingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettings) EXPECT() *MockSettingsMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSettings) Delete(ctx context.Context, key string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockSettingsMockRecorder) Delete(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSettings)(nil).Delete), ctx, key)
}

// Get mocks base method.
func (m *MockSettings) Get(ctx context.Context, key string) (string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockSettingsMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSettings)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockSettings) Set(ctx context.Context, key, value string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockSettingsMockRecorder) Set(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSettings)(nil).Set), ctx, key, value)
}

// Truncate mocks base method.
func (m *MockSettings) Truncate(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Truncate", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Truncate indicates an expected call of Truncate.
func (mr *MockSettingsMockRecorder) Truncate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Truncate", reflect.TypeOf((*MockSettings)(nil).Truncate), ctx)
}

// MockCredsSecrets is a mock of CredsSecrets interface.
type MockCredsSecrets struct {
	ctrl     *gomock.Controller
	recorder *MockCredsSecretsMockRecorder
}

// MockCredsSecretsMockRecorder is the mock recorder for MockCredsSecrets.
type MockCredsSecretsMockRecorder struct {
	mock *MockCredsSecrets
}

// NewMockCredsSecrets creates a new mock instance.
func NewMockCredsSecrets(ctrl *gomock.Controller) *MockCredsSecrets {
	mock := &MockCredsSecrets{ctrl: ctrl}
	mock.recorder = &MockCredsSecretsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredsSecrets) EXPECT() *MockCredsSecretsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCredsSecrets) Create(ctx context.Context, userID int, website, login, encPassword, additionalData string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userID, website, login, encPassword, additionalData)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCredsSecretsMockRecorder) Create(ctx, userID, website, login, encPassword, additionalData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCredsSecrets)(nil).Create), ctx, userID, website, login, encPassword, additionalData)
}

// Delete mocks base method.
func (m *MockCredsSecrets) Delete(ctx context.Context, uid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCredsSecretsMockRecorder) Delete(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCredsSecrets)(nil).Delete), ctx, uid)
}

// GetById mocks base method.
func (m *MockCredsSecrets) GetById(ctx context.Context, uid int64) (*models.CredsSecret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, uid)
	ret0, _ := ret[0].(*models.CredsSecret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCredsSecretsMockRecorder) GetById(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCredsSecrets)(nil).GetById), ctx, uid)
}

// GetList mocks base method.
func (m *MockCredsSecrets) GetList(ctx context.Context, userID int) ([]*models.CredsSecret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, userID)
	ret0, _ := ret[0].([]*models.CredsSecret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockCredsSecretsMockRecorder) GetList(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockCredsSecrets)(nil).GetList), ctx, userID)
}

// SetList mocks base method.
func (m *MockCredsSecrets) SetList(ctx context.Context, list []models.CredsSecret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetList", ctx, list)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetList indicates an expected call of SetList.
func (mr *MockCredsSecretsMockRecorder) SetList(ctx, list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetList", reflect.TypeOf((*MockCredsSecrets)(nil).SetList), ctx, list)
}

// Truncate mocks base method.
func (m *MockCredsSecrets) Truncate(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Truncate", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Truncate indicates an expected call of Truncate.
func (mr *MockCredsSecretsMockRecorder) Truncate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Truncate", reflect.TypeOf((*MockCredsSecrets)(nil).Truncate), ctx)
}
