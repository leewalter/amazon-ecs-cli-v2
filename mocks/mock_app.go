// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/archer/app.go

// Package mocks is a generated GoMock package.
package mocks

import (
	archer "github.com/aws/amazon-ecs-cli-v2/internal/pkg/archer"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApplicationStore is a mock of ApplicationStore interface
type MockApplicationStore struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationStoreMockRecorder
}

// MockApplicationStoreMockRecorder is the mock recorder for MockApplicationStore
type MockApplicationStoreMockRecorder struct {
	mock *MockApplicationStore
}

// NewMockApplicationStore creates a new mock instance
func NewMockApplicationStore(ctrl *gomock.Controller) *MockApplicationStore {
	mock := &MockApplicationStore{ctrl: ctrl}
	mock.recorder = &MockApplicationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationStore) EXPECT() *MockApplicationStoreMockRecorder {
	return m.recorder
}

// ListApplications mocks base method
func (m *MockApplicationStore) ListApplications(projectName string) ([]*archer.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", projectName)
	ret0, _ := ret[0].([]*archer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications
func (mr *MockApplicationStoreMockRecorder) ListApplications(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockApplicationStore)(nil).ListApplications), projectName)
}

// GetApplication mocks base method
func (m *MockApplicationStore) GetApplication(projectName, applicationName string) (*archer.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", projectName, applicationName)
	ret0, _ := ret[0].(*archer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication
func (mr *MockApplicationStoreMockRecorder) GetApplication(projectName, applicationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockApplicationStore)(nil).GetApplication), projectName, applicationName)
}

// CreateApplication mocks base method
func (m *MockApplicationStore) CreateApplication(app *archer.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", app)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateApplication indicates an expected call of CreateApplication
func (mr *MockApplicationStoreMockRecorder) CreateApplication(app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationStore)(nil).CreateApplication), app)
}

// MockApplicationLister is a mock of ApplicationLister interface
type MockApplicationLister struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationListerMockRecorder
}

// MockApplicationListerMockRecorder is the mock recorder for MockApplicationLister
type MockApplicationListerMockRecorder struct {
	mock *MockApplicationLister
}

// NewMockApplicationLister creates a new mock instance
func NewMockApplicationLister(ctrl *gomock.Controller) *MockApplicationLister {
	mock := &MockApplicationLister{ctrl: ctrl}
	mock.recorder = &MockApplicationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationLister) EXPECT() *MockApplicationListerMockRecorder {
	return m.recorder
}

// ListApplications mocks base method
func (m *MockApplicationLister) ListApplications(projectName string) ([]*archer.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", projectName)
	ret0, _ := ret[0].([]*archer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications
func (mr *MockApplicationListerMockRecorder) ListApplications(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockApplicationLister)(nil).ListApplications), projectName)
}

// MockApplicationGetter is a mock of ApplicationGetter interface
type MockApplicationGetter struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationGetterMockRecorder
}

// MockApplicationGetterMockRecorder is the mock recorder for MockApplicationGetter
type MockApplicationGetterMockRecorder struct {
	mock *MockApplicationGetter
}

// NewMockApplicationGetter creates a new mock instance
func NewMockApplicationGetter(ctrl *gomock.Controller) *MockApplicationGetter {
	mock := &MockApplicationGetter{ctrl: ctrl}
	mock.recorder = &MockApplicationGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationGetter) EXPECT() *MockApplicationGetterMockRecorder {
	return m.recorder
}

// GetApplication mocks base method
func (m *MockApplicationGetter) GetApplication(projectName, applicationName string) (*archer.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", projectName, applicationName)
	ret0, _ := ret[0].(*archer.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication
func (mr *MockApplicationGetterMockRecorder) GetApplication(projectName, applicationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockApplicationGetter)(nil).GetApplication), projectName, applicationName)
}

// MockApplicationCreator is a mock of ApplicationCreator interface
type MockApplicationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationCreatorMockRecorder
}

// MockApplicationCreatorMockRecorder is the mock recorder for MockApplicationCreator
type MockApplicationCreatorMockRecorder struct {
	mock *MockApplicationCreator
}

// NewMockApplicationCreator creates a new mock instance
func NewMockApplicationCreator(ctrl *gomock.Controller) *MockApplicationCreator {
	mock := &MockApplicationCreator{ctrl: ctrl}
	mock.recorder = &MockApplicationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationCreator) EXPECT() *MockApplicationCreatorMockRecorder {
	return m.recorder
}

// CreateApplication mocks base method
func (m *MockApplicationCreator) CreateApplication(app *archer.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", app)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateApplication indicates an expected call of CreateApplication
func (mr *MockApplicationCreatorMockRecorder) CreateApplication(app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationCreator)(nil).CreateApplication), app)
}