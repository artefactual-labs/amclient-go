// Code generated by MockGen. DO NOT EDIT.
// Source: go.artefactual.dev/amclient (interfaces: ProcessingConfigService)

// Package amclienttest is a generated GoMock package.
package amclienttest

import (
	context "context"
	reflect "reflect"

	amclient "go.artefactual.dev/amclient"
	gomock "go.uber.org/mock/gomock"
)

// MockProcessingConfigService is a mock of ProcessingConfigService interface.
type MockProcessingConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockProcessingConfigServiceMockRecorder
}

// MockProcessingConfigServiceMockRecorder is the mock recorder for MockProcessingConfigService.
type MockProcessingConfigServiceMockRecorder struct {
	mock *MockProcessingConfigService
}

// NewMockProcessingConfigService creates a new mock instance.
func NewMockProcessingConfigService(ctrl *gomock.Controller) *MockProcessingConfigService {
	mock := &MockProcessingConfigService{ctrl: ctrl}
	mock.recorder = &MockProcessingConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessingConfigService) EXPECT() *MockProcessingConfigServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProcessingConfigService) Get(arg0 context.Context, arg1 string) (*amclient.ProcessingConfig, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*amclient.ProcessingConfig)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockProcessingConfigServiceMockRecorder) Get(arg0, arg1 interface{}) *ProcessingConfigServiceGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProcessingConfigService)(nil).Get), arg0, arg1)
	return &ProcessingConfigServiceGetCall{Call: call}
}

// ProcessingConfigServiceGetCall wrap *gomock.Call
type ProcessingConfigServiceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProcessingConfigServiceGetCall) Return(arg0 *amclient.ProcessingConfig, arg1 *amclient.Response, arg2 error) *ProcessingConfigServiceGetCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProcessingConfigServiceGetCall) Do(f func(context.Context, string) (*amclient.ProcessingConfig, *amclient.Response, error)) *ProcessingConfigServiceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProcessingConfigServiceGetCall) DoAndReturn(f func(context.Context, string) (*amclient.ProcessingConfig, *amclient.Response, error)) *ProcessingConfigServiceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockProcessingConfigService) List(arg0 context.Context) ([]string, *amclient.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*amclient.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockProcessingConfigServiceMockRecorder) List(arg0 interface{}) *ProcessingConfigServiceListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProcessingConfigService)(nil).List), arg0)
	return &ProcessingConfigServiceListCall{Call: call}
}

// ProcessingConfigServiceListCall wrap *gomock.Call
type ProcessingConfigServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProcessingConfigServiceListCall) Return(arg0 []string, arg1 *amclient.Response, arg2 error) *ProcessingConfigServiceListCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProcessingConfigServiceListCall) Do(f func(context.Context) ([]string, *amclient.Response, error)) *ProcessingConfigServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProcessingConfigServiceListCall) DoAndReturn(f func(context.Context) ([]string, *amclient.Response, error)) *ProcessingConfigServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
