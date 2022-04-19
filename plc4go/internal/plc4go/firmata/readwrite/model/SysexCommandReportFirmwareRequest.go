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

// SysexCommandReportFirmwareRequest is the data-structure of this message
type SysexCommandReportFirmwareRequest struct {
	*SysexCommand
}

// ISysexCommandReportFirmwareRequest is the corresponding interface of SysexCommandReportFirmwareRequest
type ISysexCommandReportFirmwareRequest interface {
	ISysexCommand
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

func (m *SysexCommandReportFirmwareRequest) GetCommandType() uint8 {
	return 0x79
}

func (m *SysexCommandReportFirmwareRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *SysexCommandReportFirmwareRequest) InitializeParent(parent *SysexCommand) {}

func (m *SysexCommandReportFirmwareRequest) GetParent() *SysexCommand {
	return m.SysexCommand
}

// NewSysexCommandReportFirmwareRequest factory function for SysexCommandReportFirmwareRequest
func NewSysexCommandReportFirmwareRequest() *SysexCommandReportFirmwareRequest {
	_result := &SysexCommandReportFirmwareRequest{
		SysexCommand: NewSysexCommand(),
	}
	_result.Child = _result
	return _result
}

func CastSysexCommandReportFirmwareRequest(structType interface{}) *SysexCommandReportFirmwareRequest {
	if casted, ok := structType.(SysexCommandReportFirmwareRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*SysexCommandReportFirmwareRequest); ok {
		return casted
	}
	if casted, ok := structType.(SysexCommand); ok {
		return CastSysexCommandReportFirmwareRequest(casted.Child)
	}
	if casted, ok := structType.(*SysexCommand); ok {
		return CastSysexCommandReportFirmwareRequest(casted.Child)
	}
	return nil
}

func (m *SysexCommandReportFirmwareRequest) GetTypeName() string {
	return "SysexCommandReportFirmwareRequest"
}

func (m *SysexCommandReportFirmwareRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SysexCommandReportFirmwareRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandReportFirmwareRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandReportFirmwareRequestParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommandReportFirmwareRequest, error) {
	if pullErr := readBuffer.PullContext("SysexCommandReportFirmwareRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandReportFirmwareRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandReportFirmwareRequest{
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child, nil
}

func (m *SysexCommandReportFirmwareRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandReportFirmwareRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandReportFirmwareRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandReportFirmwareRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
