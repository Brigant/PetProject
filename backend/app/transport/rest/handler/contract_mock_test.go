// Code generated by MockGen. DO NOT EDIT.
// Source: ./contract.go

// Package handler is a generated GoMock package.
package handler

import (
	reflect "reflect"

	core "github.com/Brigant/PetPorject/backend/app/core"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountService is a mock of AccountService interface.
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService.
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance.
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAccountService) CreateUser(account core.Account) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", account)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAccountServiceMockRecorder) CreateUser(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAccountService)(nil).CreateUser), account)
}

// Login mocks base method.
func (m *MockAccountService) Login(login, password string, session core.Session) (core.TokenPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", login, password, session)
	ret0, _ := ret[0].(core.TokenPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAccountServiceMockRecorder) Login(login, password, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAccountService)(nil).Login), login, password, session)
}

// Logout mocks base method.
func (m *MockAccountService) Logout(accountID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAccountServiceMockRecorder) Logout(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAccountService)(nil).Logout), accountID)
}

// ParseToken mocks base method.
func (m *MockAccountService) ParseToken(arg0 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAccountServiceMockRecorder) ParseToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAccountService)(nil).ParseToken), arg0)
}

// RefreshTokenpair mocks base method.
func (m *MockAccountService) RefreshTokenpair(session core.Session) (core.TokenPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokenpair", session)
	ret0, _ := ret[0].(core.TokenPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokenpair indicates an expected call of RefreshTokenpair.
func (mr *MockAccountServiceMockRecorder) RefreshTokenpair(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokenpair", reflect.TypeOf((*MockAccountService)(nil).RefreshTokenpair), session)
}

// MockDirectorService is a mock of DirectorService interface.
type MockDirectorService struct {
	ctrl     *gomock.Controller
	recorder *MockDirectorServiceMockRecorder
}

// MockDirectorServiceMockRecorder is the mock recorder for MockDirectorService.
type MockDirectorServiceMockRecorder struct {
	mock *MockDirectorService
}

// NewMockDirectorService creates a new mock instance.
func NewMockDirectorService(ctrl *gomock.Controller) *MockDirectorService {
	mock := &MockDirectorService{ctrl: ctrl}
	mock.recorder = &MockDirectorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDirectorService) EXPECT() *MockDirectorServiceMockRecorder {
	return m.recorder
}

// CreatDirector mocks base method.
func (m *MockDirectorService) CreatDirector() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatDirector")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatDirector indicates an expected call of CreatDirector.
func (mr *MockDirectorServiceMockRecorder) CreatDirector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatDirector", reflect.TypeOf((*MockDirectorService)(nil).CreatDirector))
}

// GetDirector mocks base method.
func (m *MockDirectorService) GetDirector() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirector")
	ret0, _ := ret[0].(error)
	return ret0
}

// GetDirector indicates an expected call of GetDirector.
func (mr *MockDirectorServiceMockRecorder) GetDirector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirector", reflect.TypeOf((*MockDirectorService)(nil).GetDirector))
}

// UpdateDirector mocks base method.
func (m *MockDirectorService) UpdateDirector() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDirector")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDirector indicates an expected call of UpdateDirector.
func (mr *MockDirectorServiceMockRecorder) UpdateDirector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDirector", reflect.TypeOf((*MockDirectorService)(nil).UpdateDirector))
}

// MockMovieService is a mock of MovieService interface.
type MockMovieService struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceMockRecorder
}

// MockMovieServiceMockRecorder is the mock recorder for MockMovieService.
type MockMovieServiceMockRecorder struct {
	mock *MockMovieService
}

// NewMockMovieService creates a new mock instance.
func NewMockMovieService(ctrl *gomock.Controller) *MockMovieService {
	mock := &MockMovieService{ctrl: ctrl}
	mock.recorder = &MockMovieServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieService) EXPECT() *MockMovieServiceMockRecorder {
	return m.recorder
}

// GetMovie mocks base method.
func (m *MockMovieService) GetMovie() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie")
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockMovieServiceMockRecorder) GetMovie() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockMovieService)(nil).GetMovie))
}

// MockListsService is a mock of ListsService interface.
type MockListsService struct {
	ctrl     *gomock.Controller
	recorder *MockListsServiceMockRecorder
}

// MockListsServiceMockRecorder is the mock recorder for MockListsService.
type MockListsServiceMockRecorder struct {
	mock *MockListsService
}

// NewMockListsService creates a new mock instance.
func NewMockListsService(ctrl *gomock.Controller) *MockListsService {
	mock := &MockListsService{ctrl: ctrl}
	mock.recorder = &MockListsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListsService) EXPECT() *MockListsServiceMockRecorder {
	return m.recorder
}

// GetList mocks base method.
func (m *MockListsService) GetList() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList")
	ret0, _ := ret[0].(error)
	return ret0
}

// GetList indicates an expected call of GetList.
func (mr *MockListsServiceMockRecorder) GetList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockListsService)(nil).GetList))
}
