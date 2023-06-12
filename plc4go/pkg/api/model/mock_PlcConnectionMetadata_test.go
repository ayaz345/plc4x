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

// MockPlcConnectionMetadata is an autogenerated mock type for the PlcConnectionMetadata type
type MockPlcConnectionMetadata struct {
	mock.Mock
}

type MockPlcConnectionMetadata_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcConnectionMetadata) EXPECT() *MockPlcConnectionMetadata_Expecter {
	return &MockPlcConnectionMetadata_Expecter{mock: &_m.Mock}
}

// CanBrowse provides a mock function with given fields:
func (_m *MockPlcConnectionMetadata) CanBrowse() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcConnectionMetadata_CanBrowse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanBrowse'
type MockPlcConnectionMetadata_CanBrowse_Call struct {
	*mock.Call
}

// CanBrowse is a helper method to define mock.On call
func (_e *MockPlcConnectionMetadata_Expecter) CanBrowse() *MockPlcConnectionMetadata_CanBrowse_Call {
	return &MockPlcConnectionMetadata_CanBrowse_Call{Call: _e.mock.On("CanBrowse")}
}

func (_c *MockPlcConnectionMetadata_CanBrowse_Call) Run(run func()) *MockPlcConnectionMetadata_CanBrowse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionMetadata_CanBrowse_Call) Return(_a0 bool) *MockPlcConnectionMetadata_CanBrowse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionMetadata_CanBrowse_Call) RunAndReturn(run func() bool) *MockPlcConnectionMetadata_CanBrowse_Call {
	_c.Call.Return(run)
	return _c
}

// CanRead provides a mock function with given fields:
func (_m *MockPlcConnectionMetadata) CanRead() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcConnectionMetadata_CanRead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanRead'
type MockPlcConnectionMetadata_CanRead_Call struct {
	*mock.Call
}

// CanRead is a helper method to define mock.On call
func (_e *MockPlcConnectionMetadata_Expecter) CanRead() *MockPlcConnectionMetadata_CanRead_Call {
	return &MockPlcConnectionMetadata_CanRead_Call{Call: _e.mock.On("CanRead")}
}

func (_c *MockPlcConnectionMetadata_CanRead_Call) Run(run func()) *MockPlcConnectionMetadata_CanRead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionMetadata_CanRead_Call) Return(_a0 bool) *MockPlcConnectionMetadata_CanRead_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionMetadata_CanRead_Call) RunAndReturn(run func() bool) *MockPlcConnectionMetadata_CanRead_Call {
	_c.Call.Return(run)
	return _c
}

// CanSubscribe provides a mock function with given fields:
func (_m *MockPlcConnectionMetadata) CanSubscribe() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcConnectionMetadata_CanSubscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanSubscribe'
type MockPlcConnectionMetadata_CanSubscribe_Call struct {
	*mock.Call
}

// CanSubscribe is a helper method to define mock.On call
func (_e *MockPlcConnectionMetadata_Expecter) CanSubscribe() *MockPlcConnectionMetadata_CanSubscribe_Call {
	return &MockPlcConnectionMetadata_CanSubscribe_Call{Call: _e.mock.On("CanSubscribe")}
}

func (_c *MockPlcConnectionMetadata_CanSubscribe_Call) Run(run func()) *MockPlcConnectionMetadata_CanSubscribe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionMetadata_CanSubscribe_Call) Return(_a0 bool) *MockPlcConnectionMetadata_CanSubscribe_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionMetadata_CanSubscribe_Call) RunAndReturn(run func() bool) *MockPlcConnectionMetadata_CanSubscribe_Call {
	_c.Call.Return(run)
	return _c
}

// CanWrite provides a mock function with given fields:
func (_m *MockPlcConnectionMetadata) CanWrite() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcConnectionMetadata_CanWrite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CanWrite'
type MockPlcConnectionMetadata_CanWrite_Call struct {
	*mock.Call
}

// CanWrite is a helper method to define mock.On call
func (_e *MockPlcConnectionMetadata_Expecter) CanWrite() *MockPlcConnectionMetadata_CanWrite_Call {
	return &MockPlcConnectionMetadata_CanWrite_Call{Call: _e.mock.On("CanWrite")}
}

func (_c *MockPlcConnectionMetadata_CanWrite_Call) Run(run func()) *MockPlcConnectionMetadata_CanWrite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionMetadata_CanWrite_Call) Return(_a0 bool) *MockPlcConnectionMetadata_CanWrite_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionMetadata_CanWrite_Call) RunAndReturn(run func() bool) *MockPlcConnectionMetadata_CanWrite_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionAttributes provides a mock function with given fields:
func (_m *MockPlcConnectionMetadata) GetConnectionAttributes() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// MockPlcConnectionMetadata_GetConnectionAttributes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionAttributes'
type MockPlcConnectionMetadata_GetConnectionAttributes_Call struct {
	*mock.Call
}

// GetConnectionAttributes is a helper method to define mock.On call
func (_e *MockPlcConnectionMetadata_Expecter) GetConnectionAttributes() *MockPlcConnectionMetadata_GetConnectionAttributes_Call {
	return &MockPlcConnectionMetadata_GetConnectionAttributes_Call{Call: _e.mock.On("GetConnectionAttributes")}
}

func (_c *MockPlcConnectionMetadata_GetConnectionAttributes_Call) Run(run func()) *MockPlcConnectionMetadata_GetConnectionAttributes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionMetadata_GetConnectionAttributes_Call) Return(_a0 map[string]string) *MockPlcConnectionMetadata_GetConnectionAttributes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionMetadata_GetConnectionAttributes_Call) RunAndReturn(run func() map[string]string) *MockPlcConnectionMetadata_GetConnectionAttributes_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcConnectionMetadata interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcConnectionMetadata creates a new instance of MockPlcConnectionMetadata. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcConnectionMetadata(t mockConstructorTestingTNewMockPlcConnectionMetadata) *MockPlcConnectionMetadata {
	mock := &MockPlcConnectionMetadata{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
