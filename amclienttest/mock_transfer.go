// Code generated by MockGen. DO NOT EDIT.
// Source: go.artefactual.dev/amclient (interfaces: TransferService)
//
// Generated by this command:
//
//	mockgen -typed -destination=./amclienttest/mock_transfer.go -package=amclienttest go.artefactual.dev/amclient TransferService
//

// Package amclienttest is a generated GoMock package.
package amclienttest

import (
	context "context"
	reflect "reflect"

	amclient "go.artefactual.dev/amclient"
	gomock "go.uber.org/mock/gomock"
)

// MockTransferService is a mock of TransferService interface.
type MockTransferService struct {
	ctrl     *gomock.Controller
	recorder *MockTransferServiceMockRecorder
}

// MockTransferServiceMockRecorder is the mock recorder for MockTransferService.
type MockTransferServiceMockRecorder struct {
	mock *MockTransferService
}

// NewMockTransferService creates a new mock instance.
func NewMockTransferService(ctrl *gomock.Controller) *MockTransferService {
	mock := &MockTransferService{ctrl: ctrl}
	mock.recorder = &MockTransferServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferService) EXPECT() *MockTransferServiceMockRecorder {
	return m.recorder
}

// Approve mocks base method.
func (m *MockTransferService) Approve(arg0 context.Context, arg1 *amclient.TransferApproveRequest) (*amclient.TransferApproveResponse, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Approve", arg0, arg1)
	ret0, _ := ret[0].(*amclient.TransferApproveResponse)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Approve indicates an expected call of Approve.
func (mr *MockTransferServiceMockRecorder) Approve(arg0, arg1 any) *MockTransferServiceApproveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Approve", reflect.TypeOf((*MockTransferService)(nil).Approve), arg0, arg1)
	return &MockTransferServiceApproveCall{Call: call}
}

// MockTransferServiceApproveCall wrap *gomock.Call
type MockTransferServiceApproveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTransferServiceApproveCall) Return(arg0 *amclient.TransferApproveResponse, arg1 *amclient.Response, arg2 error) *MockTransferServiceApproveCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTransferServiceApproveCall) Do(f func(context.Context, *amclient.TransferApproveRequest) (*amclient.TransferApproveResponse, *amclient.Response, error)) *MockTransferServiceApproveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTransferServiceApproveCall) DoAndReturn(f func(context.Context, *amclient.TransferApproveRequest) (*amclient.TransferApproveResponse, *amclient.Response, error)) *MockTransferServiceApproveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Hide mocks base method.
func (m *MockTransferService) Hide(arg0 context.Context, arg1 string) (*amclient.TransferHideResponse, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hide", arg0, arg1)
	ret0, _ := ret[0].(*amclient.TransferHideResponse)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Hide indicates an expected call of Hide.
func (mr *MockTransferServiceMockRecorder) Hide(arg0, arg1 any) *MockTransferServiceHideCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hide", reflect.TypeOf((*MockTransferService)(nil).Hide), arg0, arg1)
	return &MockTransferServiceHideCall{Call: call}
}

// MockTransferServiceHideCall wrap *gomock.Call
type MockTransferServiceHideCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTransferServiceHideCall) Return(arg0 *amclient.TransferHideResponse, arg1 *amclient.Response, arg2 error) *MockTransferServiceHideCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTransferServiceHideCall) Do(f func(context.Context, string) (*amclient.TransferHideResponse, *amclient.Response, error)) *MockTransferServiceHideCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTransferServiceHideCall) DoAndReturn(f func(context.Context, string) (*amclient.TransferHideResponse, *amclient.Response, error)) *MockTransferServiceHideCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Start mocks base method.
func (m *MockTransferService) Start(arg0 context.Context, arg1 *amclient.TransferStartRequest) (*amclient.TransferStartResponse, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(*amclient.TransferStartResponse)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Start indicates an expected call of Start.
func (mr *MockTransferServiceMockRecorder) Start(arg0, arg1 any) *MockTransferServiceStartCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTransferService)(nil).Start), arg0, arg1)
	return &MockTransferServiceStartCall{Call: call}
}

// MockTransferServiceStartCall wrap *gomock.Call
type MockTransferServiceStartCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTransferServiceStartCall) Return(arg0 *amclient.TransferStartResponse, arg1 *amclient.Response, arg2 error) *MockTransferServiceStartCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTransferServiceStartCall) Do(f func(context.Context, *amclient.TransferStartRequest) (*amclient.TransferStartResponse, *amclient.Response, error)) *MockTransferServiceStartCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTransferServiceStartCall) DoAndReturn(f func(context.Context, *amclient.TransferStartRequest) (*amclient.TransferStartResponse, *amclient.Response, error)) *MockTransferServiceStartCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Status mocks base method.
func (m *MockTransferService) Status(arg0 context.Context, arg1 string) (*amclient.TransferStatusResponse, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*amclient.TransferStatusResponse)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Status indicates an expected call of Status.
func (mr *MockTransferServiceMockRecorder) Status(arg0, arg1 any) *MockTransferServiceStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockTransferService)(nil).Status), arg0, arg1)
	return &MockTransferServiceStatusCall{Call: call}
}

// MockTransferServiceStatusCall wrap *gomock.Call
type MockTransferServiceStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTransferServiceStatusCall) Return(arg0 *amclient.TransferStatusResponse, arg1 *amclient.Response, arg2 error) *MockTransferServiceStatusCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTransferServiceStatusCall) Do(f func(context.Context, string) (*amclient.TransferStatusResponse, *amclient.Response, error)) *MockTransferServiceStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTransferServiceStatusCall) DoAndReturn(f func(context.Context, string) (*amclient.TransferStatusResponse, *amclient.Response, error)) *MockTransferServiceStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Unapproved mocks base method.
func (m *MockTransferService) Unapproved(arg0 context.Context, arg1 *amclient.TransferUnapprovedRequest) (*amclient.TransferUnapprovedResponse, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unapproved", arg0, arg1)
	ret0, _ := ret[0].(*amclient.TransferUnapprovedResponse)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Unapproved indicates an expected call of Unapproved.
func (mr *MockTransferServiceMockRecorder) Unapproved(arg0, arg1 any) *MockTransferServiceUnapprovedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unapproved", reflect.TypeOf((*MockTransferService)(nil).Unapproved), arg0, arg1)
	return &MockTransferServiceUnapprovedCall{Call: call}
}

// MockTransferServiceUnapprovedCall wrap *gomock.Call
type MockTransferServiceUnapprovedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTransferServiceUnapprovedCall) Return(arg0 *amclient.TransferUnapprovedResponse, arg1 *amclient.Response, arg2 error) *MockTransferServiceUnapprovedCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTransferServiceUnapprovedCall) Do(f func(context.Context, *amclient.TransferUnapprovedRequest) (*amclient.TransferUnapprovedResponse, *amclient.Response, error)) *MockTransferServiceUnapprovedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTransferServiceUnapprovedCall) DoAndReturn(f func(context.Context, *amclient.TransferUnapprovedRequest) (*amclient.TransferUnapprovedResponse, *amclient.Response, error)) *MockTransferServiceUnapprovedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
