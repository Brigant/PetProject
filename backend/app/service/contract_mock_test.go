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

// DeleteSesions mocks base method.
func (m *MockAccountStorage) DeleteSesions(accountID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSesions", accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSesions indicates an expected call of DeleteSesions.
func (mr *MockAccountStorageMockRecorder) DeleteSesions(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSesions", reflect.TypeOf((*MockAccountStorage)(nil).DeleteSesions), accountID)
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
func (m *MockAccountStorage) RefreshSession(session core.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshSession", session)
	ret0, _ := ret[0].(error)
	return ret0
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

// SelectSession mocks base method.
func (m *MockAccountStorage) SelectSession(session core.Session) (core.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectSession", session)
	ret0, _ := ret[0].(core.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectSession indicates an expected call of SelectSession.
func (mr *MockAccountStorageMockRecorder) SelectSession(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectSession", reflect.TypeOf((*MockAccountStorage)(nil).SelectSession), session)
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
func (m *MockDirectorStorage) InsertDirector(director core.Director) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDirector", director)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertDirector indicates an expected call of InsertDirector.
func (mr *MockDirectorStorageMockRecorder) InsertDirector(director interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDirector", reflect.TypeOf((*MockDirectorStorage)(nil).InsertDirector), director)
}

// SelectDirectorByID mocks base method.
func (m *MockDirectorStorage) SelectDirectorByID(directorID string) (core.Director, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectDirectorByID", directorID)
	ret0, _ := ret[0].(core.Director)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectDirectorByID indicates an expected call of SelectDirectorByID.
func (mr *MockDirectorStorageMockRecorder) SelectDirectorByID(directorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectDirectorByID", reflect.TypeOf((*MockDirectorStorage)(nil).SelectDirectorByID), directorID)
}

// SelectDirectorList mocks base method.
func (m *MockDirectorStorage) SelectDirectorList() ([]core.Director, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectDirectorList")
	ret0, _ := ret[0].([]core.Director)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectDirectorList indicates an expected call of SelectDirectorList.
func (mr *MockDirectorStorageMockRecorder) SelectDirectorList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectDirectorList", reflect.TypeOf((*MockDirectorStorage)(nil).SelectDirectorList))
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
func (m *MockMovieStorage) InsertMovie(movie core.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMovie", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertMovie indicates an expected call of InsertMovie.
func (mr *MockMovieStorageMockRecorder) InsertMovie(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMovie", reflect.TypeOf((*MockMovieStorage)(nil).InsertMovie), movie)
}

// SelectAllMovies mocks base method.
func (m *MockMovieStorage) SelectAllMovies(ord string) ([]core.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAllMovies", ord)
	ret0, _ := ret[0].([]core.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAllMovies indicates an expected call of SelectAllMovies.
func (mr *MockMovieStorageMockRecorder) SelectAllMovies(ord interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAllMovies", reflect.TypeOf((*MockMovieStorage)(nil).SelectAllMovies), ord)
}

// SelectMovieByID mocks base method.
func (m *MockMovieStorage) SelectMovieByID(movieID string) (core.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMovieByID", movieID)
	ret0, _ := ret[0].(core.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMovieByID indicates an expected call of SelectMovieByID.
func (mr *MockMovieStorageMockRecorder) SelectMovieByID(movieID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMovieByID", reflect.TypeOf((*MockMovieStorage)(nil).SelectMovieByID), movieID)
}
