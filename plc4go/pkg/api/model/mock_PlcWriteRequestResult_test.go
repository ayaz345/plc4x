/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.29.0. DO NOT EDIT.

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcWriteRequestResult is an autogenerated mock type for the PlcWriteRequestResult type
type MockPlcWriteRequestResult struct {
	mock.Mock
}

type MockPlcWriteRequestResult_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcWriteRequestResult) EXPECT() *MockPlcWriteRequestResult_Expecter {
	return &MockPlcWriteRequestResult_Expecter{mock: &_m.Mock}
}

// GetErr provides a mock function with given fields:
func (_m *MockPlcWriteRequestResult) GetErr() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcWriteRequestResult_GetErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetErr'
type MockPlcWriteRequestResult_GetErr_Call struct {
	*mock.Call
}

// GetErr is a helper method to define mock.On call
func (_e *MockPlcWriteRequestResult_Expecter) GetErr() *MockPlcWriteRequestResult_GetErr_Call {
	return &MockPlcWriteRequestResult_GetErr_Call{Call: _e.mock.On("GetErr")}
}

func (_c *MockPlcWriteRequestResult_GetErr_Call) Run(run func()) *MockPlcWriteRequestResult_GetErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequestResult_GetErr_Call) Return(_a0 error) *MockPlcWriteRequestResult_GetErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequestResult_GetErr_Call) RunAndReturn(run func() error) *MockPlcWriteRequestResult_GetErr_Call {
	_c.Call.Return(run)
	return _c
}

// GetRequest provides a mock function with given fields:
func (_m *MockPlcWriteRequestResult) GetRequest() PlcWriteRequest {
	ret := _m.Called()

	var r0 PlcWriteRequest
	if rf, ok := ret.Get(0).(func() PlcWriteRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcWriteRequest)
		}
	}

	return r0
}

// MockPlcWriteRequestResult_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type MockPlcWriteRequestResult_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
func (_e *MockPlcWriteRequestResult_Expecter) GetRequest() *MockPlcWriteRequestResult_GetRequest_Call {
	return &MockPlcWriteRequestResult_GetRequest_Call{Call: _e.mock.On("GetRequest")}
}

func (_c *MockPlcWriteRequestResult_GetRequest_Call) Run(run func()) *MockPlcWriteRequestResult_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequestResult_GetRequest_Call) Return(_a0 PlcWriteRequest) *MockPlcWriteRequestResult_GetRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequestResult_GetRequest_Call) RunAndReturn(run func() PlcWriteRequest) *MockPlcWriteRequestResult_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetResponse provides a mock function with given fields:
func (_m *MockPlcWriteRequestResult) GetResponse() PlcWriteResponse {
	ret := _m.Called()

	var r0 PlcWriteResponse
	if rf, ok := ret.Get(0).(func() PlcWriteResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcWriteResponse)
		}
	}

	return r0
}

// MockPlcWriteRequestResult_GetResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResponse'
type MockPlcWriteRequestResult_GetResponse_Call struct {
	*mock.Call
}

// GetResponse is a helper method to define mock.On call
func (_e *MockPlcWriteRequestResult_Expecter) GetResponse() *MockPlcWriteRequestResult_GetResponse_Call {
	return &MockPlcWriteRequestResult_GetResponse_Call{Call: _e.mock.On("GetResponse")}
}

func (_c *MockPlcWriteRequestResult_GetResponse_Call) Run(run func()) *MockPlcWriteRequestResult_GetResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequestResult_GetResponse_Call) Return(_a0 PlcWriteResponse) *MockPlcWriteRequestResult_GetResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequestResult_GetResponse_Call) RunAndReturn(run func() PlcWriteResponse) *MockPlcWriteRequestResult_GetResponse_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcWriteRequestResult interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcWriteRequestResult creates a new instance of MockPlcWriteRequestResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcWriteRequestResult(t mockConstructorTestingTNewMockPlcWriteRequestResult) *MockPlcWriteRequestResult {
	mock := &MockPlcWriteRequestResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
