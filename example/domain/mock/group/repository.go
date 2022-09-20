// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	group "github.com/sivchari/go-rookie-gym/domain/group"
	infrastructure "github.com/sivchari/go-rookie-gym/infrastructure"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Group mocks base method.
func (m *MockRepository) Group(ctx context.Context, group *group.Group) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", ctx, group)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Group indicates an expected call of Group.
func (mr *MockRepositoryMockRecorder) Group(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockRepository)(nil).Group), ctx, group)
}

// Groups mocks base method.
func (m *MockRepository) Groups(ctx context.Context, id int) ([]*group.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Groups", ctx, id)
	ret0, _ := ret[0].([]*group.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Groups indicates an expected call of Groups.
func (mr *MockRepositoryMockRecorder) Groups(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Groups", reflect.TypeOf((*MockRepository)(nil).Groups), ctx, id)
}

// Transaction mocks base method.
func (m *MockRepository) Transaction(txm infrastructure.TxManager) group.Repository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", txm)
	ret0, _ := ret[0].(group.Repository)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockRepositoryMockRecorder) Transaction(txm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockRepository)(nil).Transaction), txm)
}