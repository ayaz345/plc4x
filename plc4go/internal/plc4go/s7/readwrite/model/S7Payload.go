/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7Payload is the data-structure of this message
type S7Payload struct {

	// Arguments.
	Parameter S7Parameter
	Child     IS7PayloadChild
}

// IS7Payload is the corresponding interface of S7Payload
type IS7Payload interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetParameterParameterType returns ParameterParameterType (discriminator field)
	GetParameterParameterType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7PayloadParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7Payload, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7PayloadChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7Payload)
	GetParent() *S7Payload

	GetTypeName() string
	IS7Payload
}

// NewS7Payload factory function for S7Payload
func NewS7Payload(parameter S7Parameter) *S7Payload {
	return &S7Payload{Parameter: parameter}
}

func CastS7Payload(structType interface{}) *S7Payload {
	if casted, ok := structType.(S7Payload); ok {
		return &casted
	}
	if casted, ok := structType.(*S7Payload); ok {
		return casted
	}
	if casted, ok := structType.(IS7PayloadChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *S7Payload) GetTypeName() string {
	return "S7Payload"
}

func (m *S7Payload) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7Payload) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *S7Payload) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *S7Payload) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadParse(readBuffer utils.ReadBuffer, messageType uint8, parameter *S7Parameter) (*S7Payload, error) {
	if pullErr := readBuffer.PullContext("S7Payload"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7PayloadChild interface {
		InitializeParent(*S7Payload)
		GetParent() *S7Payload
	}
	var _child S7PayloadChild
	var typeSwitchError error
	switch {
	case CastS7Parameter(parameter).Child.GetParameterType() == 0x04 && messageType == 0x03: // S7PayloadReadVarResponse
		_child, typeSwitchError = S7PayloadReadVarResponseParse(readBuffer, messageType, parameter)
	case CastS7Parameter(parameter).Child.GetParameterType() == 0x05 && messageType == 0x01: // S7PayloadWriteVarRequest
		_child, typeSwitchError = S7PayloadWriteVarRequestParse(readBuffer, messageType, parameter)
	case CastS7Parameter(parameter).Child.GetParameterType() == 0x05 && messageType == 0x03: // S7PayloadWriteVarResponse
		_child, typeSwitchError = S7PayloadWriteVarResponseParse(readBuffer, messageType, parameter)
	case CastS7Parameter(parameter).Child.GetParameterType() == 0x00 && messageType == 0x07: // S7PayloadUserData
		_child, typeSwitchError = S7PayloadUserDataParse(readBuffer, messageType, parameter)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7Payload"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *S7Payload) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7Payload) SerializeParent(writeBuffer utils.WriteBuffer, child IS7Payload, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7Payload"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Payload"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7Payload) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
