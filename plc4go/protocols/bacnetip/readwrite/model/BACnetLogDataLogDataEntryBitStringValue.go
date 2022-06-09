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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLogDataLogDataEntryBitStringValue is the data-structure of this message
type BACnetLogDataLogDataEntryBitStringValue struct {
	*BACnetLogDataLogDataEntry
	BitStringValue *BACnetContextTagBitString
}

// IBACnetLogDataLogDataEntryBitStringValue is the corresponding interface of BACnetLogDataLogDataEntryBitStringValue
type IBACnetLogDataLogDataEntryBitStringValue interface {
	IBACnetLogDataLogDataEntry
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() *BACnetContextTagBitString
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetLogDataLogDataEntryBitStringValue) InitializeParent(parent *BACnetLogDataLogDataEntry, peekedTagHeader *BACnetTagHeader) {
	m.BACnetLogDataLogDataEntry.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetLogDataLogDataEntryBitStringValue) GetParent() *BACnetLogDataLogDataEntry {
	return m.BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLogDataLogDataEntryBitStringValue) GetBitStringValue() *BACnetContextTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryBitStringValue factory function for BACnetLogDataLogDataEntryBitStringValue
func NewBACnetLogDataLogDataEntryBitStringValue(bitStringValue *BACnetContextTagBitString, peekedTagHeader *BACnetTagHeader) *BACnetLogDataLogDataEntryBitStringValue {
	_result := &BACnetLogDataLogDataEntryBitStringValue{
		BitStringValue:            bitStringValue,
		BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetLogDataLogDataEntryBitStringValue(structType interface{}) *BACnetLogDataLogDataEntryBitStringValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryBitStringValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryBitStringValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetLogDataLogDataEntry); ok {
		return CastBACnetLogDataLogDataEntryBitStringValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntry); ok {
		return CastBACnetLogDataLogDataEntryBitStringValue(casted.Child)
	}
	return nil
}

func (m *BACnetLogDataLogDataEntryBitStringValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryBitStringValue"
}

func (m *BACnetLogDataLogDataEntryBitStringValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLogDataLogDataEntryBitStringValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLogDataLogDataEntryBitStringValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataEntryBitStringValueParse(readBuffer utils.ReadBuffer) (*BACnetLogDataLogDataEntryBitStringValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryBitStringValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bitStringValue)
	if pullErr := readBuffer.PullContext("bitStringValue"); pullErr != nil {
		return nil, pullErr
	}
	_bitStringValue, _bitStringValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(5)), BACnetDataType(BACnetDataType_BIT_STRING))
	if _bitStringValueErr != nil {
		return nil, errors.Wrap(_bitStringValueErr, "Error parsing 'bitStringValue' field")
	}
	bitStringValue := CastBACnetContextTagBitString(_bitStringValue)
	if closeErr := readBuffer.CloseContext("bitStringValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryBitStringValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetLogDataLogDataEntryBitStringValue{
		BitStringValue:            CastBACnetContextTagBitString(bitStringValue),
		BACnetLogDataLogDataEntry: &BACnetLogDataLogDataEntry{},
	}
	_child.BACnetLogDataLogDataEntry.Child = _child
	return _child, nil
}

func (m *BACnetLogDataLogDataEntryBitStringValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryBitStringValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (bitStringValue)
		if pushErr := writeBuffer.PushContext("bitStringValue"); pushErr != nil {
			return pushErr
		}
		_bitStringValueErr := m.BitStringValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("bitStringValue"); popErr != nil {
			return popErr
		}
		if _bitStringValueErr != nil {
			return errors.Wrap(_bitStringValueErr, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryBitStringValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetLogDataLogDataEntryBitStringValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
