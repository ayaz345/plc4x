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

// HVACRawLevels is the corresponding interface of HVACRawLevels
type HVACRawLevels interface {
	utils.LengthAware
	utils.Serializable
	// GetRawValue returns RawValue (property field)
	GetRawValue() int16
	// GetValueInPercent returns ValueInPercent (virtual field)
	GetValueInPercent() float32
}

// HVACRawLevelsExactly can be used when we want exactly this type and not a type which fulfills HVACRawLevels.
// This is useful for switch cases.
type HVACRawLevelsExactly interface {
	HVACRawLevels
	isHVACRawLevels() bool
}

// _HVACRawLevels is the data-structure of this message
type _HVACRawLevels struct {
	RawValue int16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACRawLevels) GetRawValue() int16 {
	return m.RawValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACRawLevels) GetValueInPercent() float32 {
	return float32(float32(m.GetRawValue()) / float32(float32(32767)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHVACRawLevels factory function for _HVACRawLevels
func NewHVACRawLevels(rawValue int16) *_HVACRawLevels {
	return &_HVACRawLevels{RawValue: rawValue}
}

// Deprecated: use the interface for direct cast
func CastHVACRawLevels(structType interface{}) HVACRawLevels {
	if casted, ok := structType.(HVACRawLevels); ok {
		return casted
	}
	if casted, ok := structType.(*HVACRawLevels); ok {
		return *casted
	}
	return nil
}

func (m *_HVACRawLevels) GetTypeName() string {
	return "HVACRawLevels"
}

func (m *_HVACRawLevels) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_HVACRawLevels) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (rawValue)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_HVACRawLevels) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func HVACRawLevelsParse(theBytes []byte) (HVACRawLevels, error) {
	return HVACRawLevelsParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func HVACRawLevelsParseWithBuffer(readBuffer utils.ReadBuffer) (HVACRawLevels, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACRawLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACRawLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (rawValue)
	_rawValue, _rawValueErr := readBuffer.ReadInt16("rawValue", 16)
	if _rawValueErr != nil {
		return nil, errors.Wrap(_rawValueErr, "Error parsing 'rawValue' field of HVACRawLevels")
	}
	rawValue := _rawValue

	// Virtual field
	_valueInPercent := float32(rawValue) / float32(float32(32767))
	valueInPercent := float32(_valueInPercent)
	_ = valueInPercent

	if closeErr := readBuffer.CloseContext("HVACRawLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACRawLevels")
	}

	// Create the instance
	return &_HVACRawLevels{
		RawValue: rawValue,
	}, nil
}

func (m *_HVACRawLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACRawLevels) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("HVACRawLevels"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACRawLevels")
	}

	// Simple Field (rawValue)
	rawValue := int16(m.GetRawValue())
	_rawValueErr := writeBuffer.WriteInt16("rawValue", 16, (rawValue))
	if _rawValueErr != nil {
		return errors.Wrap(_rawValueErr, "Error serializing 'rawValue' field")
	}
	// Virtual field
	if _valueInPercentErr := writeBuffer.WriteVirtual("valueInPercent", m.GetValueInPercent()); _valueInPercentErr != nil {
		return errors.Wrap(_valueInPercentErr, "Error serializing 'valueInPercent' field")
	}

	if popErr := writeBuffer.PopContext("HVACRawLevels"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACRawLevels")
	}
	return nil
}

func (m *_HVACRawLevels) isHVACRawLevels() bool {
	return true
}

func (m *_HVACRawLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
