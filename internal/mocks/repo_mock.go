// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-tip-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
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

// AddTip mocks base method.
func (m *MockRepo) AddTip(arg0 context.Context, arg1 models.Tip) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTip", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTip indicates an expected call of AddTip.
func (mr *MockRepoMockRecorder) AddTip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTip", reflect.TypeOf((*MockRepo)(nil).AddTip), arg0, arg1)
}

// AddTips mocks base method.
func (m *MockRepo) AddTips(arg0 context.Context, arg1 []models.Tip) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTips", arg0, arg1)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTips indicates an expected call of AddTips.
func (mr *MockRepoMockRecorder) AddTips(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTips", reflect.TypeOf((*MockRepo)(nil).AddTips), arg0, arg1)
}

// DescribeTip mocks base method.
func (m *MockRepo) DescribeTip(arg0 context.Context, arg1 uint64) (*models.Tip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTip", arg0, arg1)
	ret0, _ := ret[0].(*models.Tip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTip indicates an expected call of DescribeTip.
func (mr *MockRepoMockRecorder) DescribeTip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTip", reflect.TypeOf((*MockRepo)(nil).DescribeTip), arg0, arg1)
}

// ListTips mocks base method.
func (m *MockRepo) ListTips(arg0 context.Context, arg1, arg2 uint64) ([]models.Tip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTips", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Tip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTips indicates an expected call of ListTips.
func (mr *MockRepoMockRecorder) ListTips(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTips", reflect.TypeOf((*MockRepo)(nil).ListTips), arg0, arg1, arg2)
}

// RemoveTip mocks base method.
func (m *MockRepo) RemoveTip(arg0 context.Context, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTip", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTip indicates an expected call of RemoveTip.
func (mr *MockRepoMockRecorder) RemoveTip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTip", reflect.TypeOf((*MockRepo)(nil).RemoveTip), arg0, arg1)
}

// UpdateTip mocks base method.
func (m *MockRepo) UpdateTip(arg0 context.Context, arg1 models.Tip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTip", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTip indicates an expected call of UpdateTip.
func (mr *MockRepoMockRecorder) UpdateTip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTip", reflect.TypeOf((*MockRepo)(nil).UpdateTip), arg0, arg1)
}
