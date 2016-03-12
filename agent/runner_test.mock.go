// Automatically generated by MockGen. DO NOT EDIT!
// Source: runner.go

package agent

import (
	api "github.com/docker/swarm-v2/api"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// Mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockRunnerRecorder
}

// Recorder for MockRunner (not exported)
type _MockRunnerRecorder struct {
	mock *MockRunner
}

func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &_MockRunnerRecorder{mock}
	return mock
}

func (_m *MockRunner) EXPECT() *_MockRunnerRecorder {
	return _m.recorder
}

func (_m *MockRunner) Prepare(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Prepare", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRunnerRecorder) Prepare(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prepare", arg0)
}

func (_m *MockRunner) Start(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRunnerRecorder) Start(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0)
}

func (_m *MockRunner) Wait(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Wait", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRunnerRecorder) Wait(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Wait", arg0)
}

func (_m *MockRunner) Shutdown(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Shutdown", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRunnerRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Shutdown", arg0)
}

func (_m *MockRunner) Terminate(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Terminate", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRunnerRecorder) Terminate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Terminate", arg0)
}

func (_m *MockRunner) Remove(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "Remove", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRunnerRecorder) Remove(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Remove", arg0)
}

func (_m *MockRunner) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRunnerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of Reporter interface
type MockReporter struct {
	ctrl     *gomock.Controller
	recorder *_MockReporterRecorder
}

// Recorder for MockReporter (not exported)
type _MockReporterRecorder struct {
	mock *MockReporter
}

func NewMockReporter(ctrl *gomock.Controller) *MockReporter {
	mock := &MockReporter{ctrl: ctrl}
	mock.recorder = &_MockReporterRecorder{mock}
	return mock
}

func (_m *MockReporter) EXPECT() *_MockReporterRecorder {
	return _m.recorder
}

func (_m *MockReporter) Report(ctx context.Context, state api.TaskState) error {
	ret := _m.ctrl.Call(_m, "Report", ctx, state)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockReporterRecorder) Report(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Report", arg0, arg1)
}
