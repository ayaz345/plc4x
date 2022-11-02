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

package model

import (
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// EnableControlData is the corresponding interface of EnableControlData
type EnableControlData interface {
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() EnableControlCommandTypeContainer
	// GetEnableNetworkVariable returns EnableNetworkVariable (property field)
	GetEnableNetworkVariable() byte
	// GetValue returns Value (property field)
	GetValue() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() EnableControlCommandType
}

// EnableControlDataExactly can be used when we want exactly this type and not a type which fulfills EnableControlData.
// This is useful for switch cases.
type EnableControlDataExactly interface {
	EnableControlData
	isEnableControlData() bool
}

// _EnableControlData is the data-structure of this message
type _EnableControlData struct {
	CommandTypeContainer  EnableControlCommandTypeContainer
	EnableNetworkVariable byte
	Value                 byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EnableControlData) GetCommandTypeContainer() EnableControlCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_EnableControlData) GetEnableNetworkVariable() byte {
	return m.EnableNetworkVariable
}

func (m *_EnableControlData) GetValue() byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_EnableControlData) GetCommandType() EnableControlCommandType {
	return CastEnableControlCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEnableControlData factory function for _EnableControlData
func NewEnableControlData(commandTypeContainer EnableControlCommandTypeContainer, enableNetworkVariable byte, value byte) *_EnableControlData {
	return &_EnableControlData{CommandTypeContainer: commandTypeContainer, EnableNetworkVariable: enableNetworkVariable, Value: value}
}

// Deprecated: use the interface for direct cast
func CastEnableControlData(structType interface{}) EnableControlData {
	if casted, ok := structType.(EnableControlData); ok {
		return casted
	}
	if casted, ok := structType.(*EnableControlData); ok {
		return *casted
	}
	return nil
}

func (m *_EnableControlData) GetTypeName() string {
	return "EnableControlData"
}

func (m *_EnableControlData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_EnableControlData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (enableNetworkVariable)
	lengthInBits += 8

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *_EnableControlData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func EnableControlDataParse(theBytes []byte) (EnableControlData, error) {
	return EnableControlDataParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func EnableControlDataParseWithBuffer(readBuffer utils.ReadBuffer) (EnableControlData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EnableControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EnableControlData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsEnableControlCommandTypeContainer(readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := EnableControlCommandTypeContainerParseWithBuffer(readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of EnableControlData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := EnableControlCommandType(_commandType)
	_ = commandType

	// Simple Field (enableNetworkVariable)
	_enableNetworkVariable, _enableNetworkVariableErr := readBuffer.ReadByte("enableNetworkVariable")
	if _enableNetworkVariableErr != nil {
		return nil, errors.Wrap(_enableNetworkVariableErr, "Error parsing 'enableNetworkVariable' field of EnableControlData")
	}
	enableNetworkVariable := _enableNetworkVariable

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadByte("value")
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of EnableControlData")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("EnableControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EnableControlData")
	}

	// Create the instance
	return &_EnableControlData{
		CommandTypeContainer:  commandTypeContainer,
		EnableNetworkVariable: enableNetworkVariable,
		Value:                 value,
	}, nil
}

func (m *_EnableControlData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EnableControlData) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("EnableControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EnableControlData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	if _commandTypeErr := writeBuffer.WriteVirtual("commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Simple Field (enableNetworkVariable)
	enableNetworkVariable := byte(m.GetEnableNetworkVariable())
	_enableNetworkVariableErr := writeBuffer.WriteByte("enableNetworkVariable", (enableNetworkVariable))
	if _enableNetworkVariableErr != nil {
		return errors.Wrap(_enableNetworkVariableErr, "Error serializing 'enableNetworkVariable' field")
	}

	// Simple Field (value)
	value := byte(m.GetValue())
	_valueErr := writeBuffer.WriteByte("value", (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("EnableControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EnableControlData")
	}
	return nil
}

func (m *_EnableControlData) isEnableControlData() bool {
	return true
}

func (m *_EnableControlData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
