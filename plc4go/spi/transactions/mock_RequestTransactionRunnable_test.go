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

package transactions

import mock "github.com/stretchr/testify/mock"

// MockRequestTransactionRunnable is an autogenerated mock type for the RequestTransactionRunnable type
type MockRequestTransactionRunnable struct {
	mock.Mock
}

type MockRequestTransactionRunnable_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRequestTransactionRunnable) EXPECT() *MockRequestTransactionRunnable_Expecter {
	return &MockRequestTransactionRunnable_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockRequestTransactionRunnable) Execute(_a0 RequestTransaction) {
	_m.Called(_a0)
}

// MockRequestTransactionRunnable_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockRequestTransactionRunnable_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 RequestTransaction
func (_e *MockRequestTransactionRunnable_Expecter) Execute(_a0 interface{}) *MockRequestTransactionRunnable_Execute_Call {
	return &MockRequestTransactionRunnable_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockRequestTransactionRunnable_Execute_Call) Run(run func(_a0 RequestTransaction)) *MockRequestTransactionRunnable_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(RequestTransaction))
	})
	return _c
}

func (_c *MockRequestTransactionRunnable_Execute_Call) Return() *MockRequestTransactionRunnable_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRequestTransactionRunnable_Execute_Call) RunAndReturn(run func(RequestTransaction)) *MockRequestTransactionRunnable_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRequestTransactionRunnable interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRequestTransactionRunnable creates a new instance of MockRequestTransactionRunnable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRequestTransactionRunnable(t mockConstructorTestingTNewMockRequestTransactionRunnable) *MockRequestTransactionRunnable {
	mock := &MockRequestTransactionRunnable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
