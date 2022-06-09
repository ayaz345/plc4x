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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLogDataLogDataEntryAnyValue is the data-structure of this message
type BACnetLogDataLogDataEntryAnyValue struct {
	*BACnetLogDataLogDataEntry
	AnyValue *BACnetConstructedData
}

// IBACnetLogDataLogDataEntryAnyValue is the corresponding interface of BACnetLogDataLogDataEntryAnyValue
type IBACnetLogDataLogDataEntryAnyValue interface {
	IBACnetLogDataLogDataEntry
	// GetAnyValue returns AnyValue (property field)
	GetAnyValue() *BACnetConstructedData
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

func (m *BACnetLogDataLogDataEntryAnyValue) InitializeParent(parent *BACnetLogDataLogDataEntry, peekedTagHeader *BACnetTagHeader) {
	m.BACnetLogDataLogDataEntry.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetLogDataLogDataEntryAnyValue) GetParent() *BACnetLogDataLogDataEntry {
	return m.BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLogDataLogDataEntryAnyValue) GetAnyValue() *BACnetConstructedData {
	return m.AnyValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryAnyValue factory function for BACnetLogDataLogDataEntryAnyValue
func NewBACnetLogDataLogDataEntryAnyValue(anyValue *BACnetConstructedData, peekedTagHeader *BACnetTagHeader) *BACnetLogDataLogDataEntryAnyValue {
	_result := &BACnetLogDataLogDataEntryAnyValue{
		AnyValue:                  anyValue,
		BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetLogDataLogDataEntryAnyValue(structType interface{}) *BACnetLogDataLogDataEntryAnyValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryAnyValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryAnyValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetLogDataLogDataEntry); ok {
		return CastBACnetLogDataLogDataEntryAnyValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntry); ok {
		return CastBACnetLogDataLogDataEntryAnyValue(casted.Child)
	}
	return nil
}

func (m *BACnetLogDataLogDataEntryAnyValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryAnyValue"
}

func (m *BACnetLogDataLogDataEntryAnyValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLogDataLogDataEntryAnyValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (anyValue)
	if m.AnyValue != nil {
		lengthInBits += (*m.AnyValue).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetLogDataLogDataEntryAnyValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataEntryAnyValueParse(readBuffer utils.ReadBuffer) (*BACnetLogDataLogDataEntryAnyValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryAnyValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (anyValue) (Can be skipped, if a given expression evaluates to false)
	var anyValue *BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("anyValue"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(8), BACnetObjectType_VENDOR_PROPRIETARY_VALUE, BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'anyValue' field")
		default:
			anyValue = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("anyValue"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryAnyValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetLogDataLogDataEntryAnyValue{
		AnyValue:                  CastBACnetConstructedData(anyValue),
		BACnetLogDataLogDataEntry: &BACnetLogDataLogDataEntry{},
	}
	_child.BACnetLogDataLogDataEntry.Child = _child
	return _child, nil
}

func (m *BACnetLogDataLogDataEntryAnyValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryAnyValue"); pushErr != nil {
			return pushErr
		}

		// Optional Field (anyValue) (Can be skipped, if the value is null)
		var anyValue *BACnetConstructedData = nil
		if m.AnyValue != nil {
			if pushErr := writeBuffer.PushContext("anyValue"); pushErr != nil {
				return pushErr
			}
			anyValue = m.AnyValue
			_anyValueErr := anyValue.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("anyValue"); popErr != nil {
				return popErr
			}
			if _anyValueErr != nil {
				return errors.Wrap(_anyValueErr, "Error serializing 'anyValue' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryAnyValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetLogDataLogDataEntryAnyValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
