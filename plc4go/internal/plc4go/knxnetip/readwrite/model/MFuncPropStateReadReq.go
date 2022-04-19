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

// MFuncPropStateReadReq is the data-structure of this message
type MFuncPropStateReadReq struct {
	*CEMI

	// Arguments.
	Size uint16
}

// IMFuncPropStateReadReq is the corresponding interface of MFuncPropStateReadReq
type IMFuncPropStateReadReq interface {
	ICEMI
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

func (m *MFuncPropStateReadReq) GetMessageCode() uint8 {
	return 0xF9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *MFuncPropStateReadReq) InitializeParent(parent *CEMI) {}

func (m *MFuncPropStateReadReq) GetParent() *CEMI {
	return m.CEMI
}

// NewMFuncPropStateReadReq factory function for MFuncPropStateReadReq
func NewMFuncPropStateReadReq(size uint16) *MFuncPropStateReadReq {
	_result := &MFuncPropStateReadReq{
		CEMI: NewCEMI(size),
	}
	_result.Child = _result
	return _result
}

func CastMFuncPropStateReadReq(structType interface{}) *MFuncPropStateReadReq {
	if casted, ok := structType.(MFuncPropStateReadReq); ok {
		return &casted
	}
	if casted, ok := structType.(*MFuncPropStateReadReq); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastMFuncPropStateReadReq(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastMFuncPropStateReadReq(casted.Child)
	}
	return nil
}

func (m *MFuncPropStateReadReq) GetTypeName() string {
	return "MFuncPropStateReadReq"
}

func (m *MFuncPropStateReadReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *MFuncPropStateReadReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *MFuncPropStateReadReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MFuncPropStateReadReqParse(readBuffer utils.ReadBuffer, size uint16) (*MFuncPropStateReadReq, error) {
	if pullErr := readBuffer.PullContext("MFuncPropStateReadReq"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropStateReadReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MFuncPropStateReadReq{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child, nil
}

func (m *MFuncPropStateReadReq) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropStateReadReq"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("MFuncPropStateReadReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MFuncPropStateReadReq) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
