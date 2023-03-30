// Code generated by MockGen. DO NOT EDIT.
// Source: ./contract.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	core "github.com/Brigant/PetPorject/backend/app/core"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountStorage is a mock of AccountStorage interface.
type MockAccountStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAccountStorageMockRecorder
}

// MockAccountStorageMockRecorder is the mock recorder for MockAccountStorage.
type MockAccountStorageMockRecorder struct {
	mock *MockAccountStorage
}

// NewMockAccountStorage creates a new mock instance.
func NewMockAccountStorage(ctrl *gomock.Controller) *MockAccountStorage {
	mock := &MockAccountStorage{ctrl: ctrl}
	mock.recorder = &MockAccountStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountStorage) EXPECT() *MockAccountStorageMockRecorder {
	return m.recorder
}

// InsertAccount mocks base method.
func (m *MockAccountStorage) InsertAccount(account core.Account) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAccount", account)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertAccount indicates an expected call of InsertAccount.
func (mr *MockAccountStorageMockRecorder) InsertAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAccount", reflect.TypeOf((*MockAccountStorage)(nil).InsertAccount), account)
}

// InsertSession mocks base method.
func (m *MockAccountStorage) InsertSession(session core.Session) (core.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertSession", session)
	ret0, _ := ret[0].(core.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertSession indicates an expected call of InsertSession.
func (mr *MockAccountStorageMockRecorder) InsertSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertSession", reflect.TypeOf((*MockAccountStorage)(nil).InsertSession), session)
}

// RefreshSession mocks base method.
func (m *MockAccountStorage) RefreshSession(session core.Session) (core.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshSession", session)
	ret0, _ := ret[0].(core.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshSession indicates an expected call of RefreshSession.
func (mr *MockAccountStorageMockRecorder) RefreshSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshSession", reflect.TypeOf((*MockAccountStorage)(nil).RefreshSession), session)
}

// SelectAccountByID mocks base method.
func (m *MockAccountStorage) SelectAccountByID(accountID string) (core.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAccountByID", accountID)
	ret0, _ := ret[0].(core.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAccountByID indicates an expected call of SelectAccountByID.
func (mr *MockAccountStorageMockRecorder) SelectAccountByID(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAccountByID", reflect.TypeOf((*MockAccountStorage)(nil).SelectAccountByID), accountID)
}

// SelectAccountByPhone mocks base method.
func (m *MockAccountStorage) SelectAccountByPhone(phone string) (core.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAccountByPhone", phone)
	ret0, _ := ret[0].(core.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAccountByPhone indicates an expected call of SelectAccountByPhone.
func (mr *MockAccountStorageMockRecorder) SelectAccountByPhone(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAccountByPhone", reflect.TypeOf((*MockAccountStorage)(nil).SelectAccountByPhone), phone)
}

// MockDirectorStorage is a mock of DirectorStorage interface.
type MockDirectorStorage struct {
	ctrl     *gomock.Controller
	recorder *MockDirectorStorageMockRecorder
}

// MockDirectorStorageMockRecorder is the mock recorder for MockDirectorStorage.
type MockDirectorStorageMockRecorder struct {
	mock *MockDirectorStorage
}

// NewMockDirectorStorage creates a new mock instance.
func NewMockDirectorStorage(ctrl *gomock.Controller) *MockDirectorStorage {
	mock := &MockDirectorStorage{ctrl: ctrl}
	mock.recorder = &MockDirectorStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectorStorage) EXPECT() *MockDirectorStorageMockRecorder {
	return m.recorder
}

// InsertDirector mocks base method.
func (m *MockDirectorStorage) InsertDirector() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDirector")
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertDirector indicates an expected call of InsertDirector.
func (mr *MockDirectorStorageMockRecorder) InsertDirector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDirector", reflect.TypeOf((*MockDirectorStorage)(nil).InsertDirector))
}

// SelectDirector mocks base method.
func (m *MockDirectorStorage) SelectDirector() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectDirector")
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectDirector indicates an expected call of SelectDirector.
func (mr *MockDirectorStorageMockRecorder) SelectDirector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectDirector", reflect.TypeOf((*MockDirectorStorage)(nil).SelectDirector))
}

// UpdateDirector mocks base method.
func (m *MockDirectorStorage) UpdateDirector() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDirector")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDirector indicates an expected call of UpdateDirector.
func (mr *MockDirectorStorageMockRecorder) UpdateDirector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDirector", reflect.TypeOf((*MockDirectorStorage)(nil).UpdateDirector))
}

// MockMovieStorage is a mock of MovieStorage interface.
type MockMovieStorage struct {
	ctrl     *gomock.Controller
	recorder *MockMovieStorageMockRecorder
}

// MockMovieStorageMockRecorder is the mock recorder for MockMovieStorage.
type MockMovieStorageMockRecorder struct {
	mock *MockMovieStorage
}

// NewMockMovieStorage creates a new mock instance.
func NewMockMovieStorage(ctrl *gomock.Controller) *MockMovieStorage {
	mock := &MockMovieStorage{ctrl: ctrl}
	mock.recorder = &MockMovieStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieStorage) EXPECT() *MockMovieStorageMockRecorder {
	return m.recorder
}

// InsertMovie mocks base method.
func (m *MockMovieStorage) InsertMovie() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMovie")
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertMovie indicates an expected call of InsertMovie.
func (mr *MockMovieStorageMockRecorder) InsertMovie() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMovie", reflect.TypeOf((*MockMovieStorage)(nil).InsertMovie))
}

// SelectAllMovies mocks base method.
func (m *MockMovieStorage) SelectAllMovies() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAllMovies")
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectAllMovies indicates an expected call of SelectAllMovies.
func (mr *MockMovieStorageMockRecorder) SelectAllMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAllMovies", reflect.TypeOf((*MockMovieStorage)(nil).SelectAllMovies))
}

// SelectMovieByID mocks base method.
func (m *MockMovieStorage) SelectMovieByID() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMovieByID")
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectMovieByID indicates an expected call of SelectMovieByID.
func (mr *MockMovieStorageMockRecorder) SelectMovieByID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMovieByID", reflect.TypeOf((*MockMovieStorage)(nil).SelectMovieByID))
}
