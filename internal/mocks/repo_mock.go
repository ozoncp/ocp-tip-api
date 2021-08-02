// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-tip-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-tip-api/internal/models"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddTips mocks base method.
func (m *MockRepo) AddTips(arg0 []models.Tip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTips", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTips indicates an expected call of AddTips.
func (mr *MockRepoMockRecorder) AddTips(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTips", reflect.TypeOf((*MockRepo)(nil).AddTips), arg0)
}

// DescribeTip mocks base method.
func (m *MockRepo) DescribeTip(arg0 uint64) (*models.Tip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTip", arg0)
	ret0, _ := ret[0].(*models.Tip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTip indicates an expected call of DescribeTip.
func (mr *MockRepoMockRecorder) DescribeTip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTip", reflect.TypeOf((*MockRepo)(nil).DescribeTip), arg0)
}

// ListTips mocks base method.
func (m *MockRepo) ListTips(arg0, arg1 uint64) ([]models.Tip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTips", arg0, arg1)
	ret0, _ := ret[0].([]models.Tip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTips indicates an expected call of ListTips.
func (mr *MockRepoMockRecorder) ListTips(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTips", reflect.TypeOf((*MockRepo)(nil).ListTips), arg0, arg1)
}