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
)

// Code generated by code-generation. DO NOT EDIT.

// S7MessageUserData is the data-structure of this message
type S7MessageUserData struct {
	*S7Message
}

// IS7MessageUserData is the corresponding interface of S7MessageUserData
type IS7MessageUserData interface {
	IS7Message
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *S7MessageUserData) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7MessageUserData) InitializeParent(parent *S7Message, tpduReference uint16, parameter *S7Parameter, payload *S7Payload) {
	m.S7Message.TpduReference = tpduReference
	m.S7Message.Parameter = parameter
	m.S7Message.Payload = payload
}

func (m *S7MessageUserData) GetParent() *S7Message {
	return m.S7Message
}

// NewS7MessageUserData factory function for S7MessageUserData
func NewS7MessageUserData(tpduReference uint16, parameter *S7Parameter, payload *S7Payload) *S7MessageUserData {
	_result := &S7MessageUserData{
		S7Message: NewS7Message(tpduReference, parameter, payload),
	}
	_result.Child = _result
	return _result
}

func CastS7MessageUserData(structType interface{}) *S7MessageUserData {
	if casted, ok := structType.(S7MessageUserData); ok {
		return &casted
	}
	if casted, ok := structType.(*S7MessageUserData); ok {
		return casted
	}
	if casted, ok := structType.(S7Message); ok {
		return CastS7MessageUserData(casted.Child)
	}
	if casted, ok := structType.(*S7Message); ok {
		return CastS7MessageUserData(casted.Child)
	}
	return nil
}

func (m *S7MessageUserData) GetTypeName() string {
	return "S7MessageUserData"
}

func (m *S7MessageUserData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7MessageUserData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *S7MessageUserData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7MessageUserDataParse(readBuffer utils.ReadBuffer) (*S7MessageUserData, error) {
	if pullErr := readBuffer.PullContext("S7MessageUserData"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7MessageUserData"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7MessageUserData{
		S7Message: &S7Message{},
	}
	_child.S7Message.Child = _child
	return _child, nil
}

func (m *S7MessageUserData) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageUserData"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("S7MessageUserData"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7MessageUserData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
