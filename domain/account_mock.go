// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/da-n/portfolio-app/domain (interfaces: AccountRepository)

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	errs "github.com/da-n/portfolio-app/errs"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// FindAccountById mocks base method.
func (m *MockAccountRepository) FindAccountById(arg0 int) (*Account, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAccountById", arg0)
	ret0, _ := ret[0].(*Account)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindAccountById indicates an expected call of FindAccountById.
func (mr *MockAccountRepositoryMockRecorder) FindAccountById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAccountById", reflect.TypeOf((*MockAccountRepository)(nil).FindAccountById), arg0)
}

// FindAllAccounts mocks base method.
func (m *MockAccountRepository) FindAllAccounts(arg0 int) ([]Account, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllAccounts", arg0)
	ret0, _ := ret[0].([]Account)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindAllAccounts indicates an expected call of FindAllAccounts.
func (mr *MockAccountRepositoryMockRecorder) FindAllAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllAccounts", reflect.TypeOf((*MockAccountRepository)(nil).FindAllAccounts), arg0)
}

// FindOrderSheetById mocks base method.
func (m *MockAccountRepository) FindOrderSheetById(arg0 int) (*OrderSheet, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrderSheetById", arg0)
	ret0, _ := ret[0].(*OrderSheet)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindOrderSheetById indicates an expected call of FindOrderSheetById.
func (mr *MockAccountRepositoryMockRecorder) FindOrderSheetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrderSheetById", reflect.TypeOf((*MockAccountRepository)(nil).FindOrderSheetById), arg0)
}

// FindPortfolioById mocks base method.
func (m *MockAccountRepository) FindPortfolioById(arg0 int) (*Portfolio, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPortfolioById", arg0)
	ret0, _ := ret[0].(*Portfolio)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindPortfolioById indicates an expected call of FindPortfolioById.
func (mr *MockAccountRepositoryMockRecorder) FindPortfolioById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPortfolioById", reflect.TypeOf((*MockAccountRepository)(nil).FindPortfolioById), arg0)
}

// SaveInstruction mocks base method.
func (m *MockAccountRepository) SaveInstruction(arg0 Instruction) (*Instruction, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveInstruction", arg0)
	ret0, _ := ret[0].(*Instruction)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// SaveInstruction indicates an expected call of SaveInstruction.
func (mr *MockAccountRepositoryMockRecorder) SaveInstruction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveInstruction", reflect.TypeOf((*MockAccountRepository)(nil).SaveInstruction), arg0)
}

// SaveOrderSheet mocks base method.
func (m *MockAccountRepository) SaveOrderSheet(arg0 OrderSheet) (*OrderSheet, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrderSheet", arg0)
	ret0, _ := ret[0].(*OrderSheet)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// SaveOrderSheet indicates an expected call of SaveOrderSheet.
func (mr *MockAccountRepositoryMockRecorder) SaveOrderSheet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrderSheet", reflect.TypeOf((*MockAccountRepository)(nil).SaveOrderSheet), arg0)
}

// SaveWithdrawalRequest mocks base method.
func (m *MockAccountRepository) SaveWithdrawalRequest(arg0 WithdrawalRequest) (*WithdrawalRequest, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWithdrawalRequest", arg0)
	ret0, _ := ret[0].(*WithdrawalRequest)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// SaveWithdrawalRequest indicates an expected call of SaveWithdrawalRequest.
func (mr *MockAccountRepositoryMockRecorder) SaveWithdrawalRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWithdrawalRequest", reflect.TypeOf((*MockAccountRepository)(nil).SaveWithdrawalRequest), arg0)
}
