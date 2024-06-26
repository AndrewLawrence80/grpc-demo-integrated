// Code generated by MockGen. DO NOT EDIT.
// Source: api/calculator/calculator_grpc.pb.go

// Package calculator is a generated GoMock package.
package calculator

import (
	context "context"
	reflect "reflect"

	calculator "github.com/andrewlawrence80/grpc-demo-integrated/api/calculator"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCalculatorClient is a mock of CalculatorClient interface.
type MockCalculatorClient struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorClientMockRecorder
}

// MockCalculatorClientMockRecorder is the mock recorder for MockCalculatorClient.
type MockCalculatorClientMockRecorder struct {
	mock *MockCalculatorClient
}

// NewMockCalculatorClient creates a new mock instance.
func NewMockCalculatorClient(ctrl *gomock.Controller) *MockCalculatorClient {
	mock := &MockCalculatorClient{ctrl: ctrl}
	mock.recorder = &MockCalculatorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculatorClient) EXPECT() *MockCalculatorClientMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCalculatorClient) Add(ctx context.Context, in *calculator.CalcRequest, opts ...grpc.CallOption) (*calculator.CalcResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*calculator.CalcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockCalculatorClientMockRecorder) Add(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCalculatorClient)(nil).Add), varargs...)
}

// Divide mocks base method.
func (m *MockCalculatorClient) Divide(ctx context.Context, in *calculator.CalcRequest, opts ...grpc.CallOption) (*calculator.CalcResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Divide", varargs...)
	ret0, _ := ret[0].(*calculator.CalcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Divide indicates an expected call of Divide.
func (mr *MockCalculatorClientMockRecorder) Divide(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Divide", reflect.TypeOf((*MockCalculatorClient)(nil).Divide), varargs...)
}

// Multiply mocks base method.
func (m *MockCalculatorClient) Multiply(ctx context.Context, in *calculator.CalcRequest, opts ...grpc.CallOption) (*calculator.CalcResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Multiply", varargs...)
	ret0, _ := ret[0].(*calculator.CalcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Multiply indicates an expected call of Multiply.
func (mr *MockCalculatorClientMockRecorder) Multiply(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Multiply", reflect.TypeOf((*MockCalculatorClient)(nil).Multiply), varargs...)
}

// Subtract mocks base method.
func (m *MockCalculatorClient) Subtract(ctx context.Context, in *calculator.CalcRequest, opts ...grpc.CallOption) (*calculator.CalcResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subtract", varargs...)
	ret0, _ := ret[0].(*calculator.CalcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subtract indicates an expected call of Subtract.
func (mr *MockCalculatorClientMockRecorder) Subtract(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subtract", reflect.TypeOf((*MockCalculatorClient)(nil).Subtract), varargs...)
}

// MockCalculatorServer is a mock of CalculatorServer interface.
type MockCalculatorServer struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorServerMockRecorder
}

// MockCalculatorServerMockRecorder is the mock recorder for MockCalculatorServer.
type MockCalculatorServerMockRecorder struct {
	mock *MockCalculatorServer
}

// NewMockCalculatorServer creates a new mock instance.
func NewMockCalculatorServer(ctrl *gomock.Controller) *MockCalculatorServer {
	mock := &MockCalculatorServer{ctrl: ctrl}
	mock.recorder = &MockCalculatorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculatorServer) EXPECT() *MockCalculatorServerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCalculatorServer) Add(arg0 context.Context, arg1 *calculator.CalcRequest) (*calculator.CalcResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(*calculator.CalcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockCalculatorServerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCalculatorServer)(nil).Add), arg0, arg1)
}

// Divide mocks base method.
func (m *MockCalculatorServer) Divide(arg0 context.Context, arg1 *calculator.CalcRequest) (*calculator.CalcResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Divide", arg0, arg1)
	ret0, _ := ret[0].(*calculator.CalcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Divide indicates an expected call of Divide.
func (mr *MockCalculatorServerMockRecorder) Divide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Divide", reflect.TypeOf((*MockCalculatorServer)(nil).Divide), arg0, arg1)
}

// Multiply mocks base method.
func (m *MockCalculatorServer) Multiply(arg0 context.Context, arg1 *calculator.CalcRequest) (*calculator.CalcResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Multiply", arg0, arg1)
	ret0, _ := ret[0].(*calculator.CalcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Multiply indicates an expected call of Multiply.
func (mr *MockCalculatorServerMockRecorder) Multiply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Multiply", reflect.TypeOf((*MockCalculatorServer)(nil).Multiply), arg0, arg1)
}

// Subtract mocks base method.
func (m *MockCalculatorServer) Subtract(arg0 context.Context, arg1 *calculator.CalcRequest) (*calculator.CalcResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subtract", arg0, arg1)
	ret0, _ := ret[0].(*calculator.CalcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subtract indicates an expected call of Subtract.
func (mr *MockCalculatorServerMockRecorder) Subtract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subtract", reflect.TypeOf((*MockCalculatorServer)(nil).Subtract), arg0, arg1)
}

// mustEmbedUnimplementedCalculatorServer mocks base method.
func (m *MockCalculatorServer) mustEmbedUnimplementedCalculatorServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedCalculatorServer")
}

// mustEmbedUnimplementedCalculatorServer indicates an expected call of mustEmbedUnimplementedCalculatorServer.
func (mr *MockCalculatorServerMockRecorder) mustEmbedUnimplementedCalculatorServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedCalculatorServer", reflect.TypeOf((*MockCalculatorServer)(nil).mustEmbedUnimplementedCalculatorServer))
}

// MockUnsafeCalculatorServer is a mock of UnsafeCalculatorServer interface.
type MockUnsafeCalculatorServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeCalculatorServerMockRecorder
}

// MockUnsafeCalculatorServerMockRecorder is the mock recorder for MockUnsafeCalculatorServer.
type MockUnsafeCalculatorServerMockRecorder struct {
	mock *MockUnsafeCalculatorServer
}

// NewMockUnsafeCalculatorServer creates a new mock instance.
func NewMockUnsafeCalculatorServer(ctrl *gomock.Controller) *MockUnsafeCalculatorServer {
	mock := &MockUnsafeCalculatorServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeCalculatorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeCalculatorServer) EXPECT() *MockUnsafeCalculatorServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedCalculatorServer mocks base method.
func (m *MockUnsafeCalculatorServer) mustEmbedUnimplementedCalculatorServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedCalculatorServer")
}

// mustEmbedUnimplementedCalculatorServer indicates an expected call of mustEmbedUnimplementedCalculatorServer.
func (mr *MockUnsafeCalculatorServerMockRecorder) mustEmbedUnimplementedCalculatorServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedCalculatorServer", reflect.TypeOf((*MockUnsafeCalculatorServer)(nil).mustEmbedUnimplementedCalculatorServer))
}
