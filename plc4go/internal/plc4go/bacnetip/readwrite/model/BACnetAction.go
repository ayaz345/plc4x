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

// The data-structure of this message
type BACnetAction struct {
	RawData   *BACnetContextTagEnumerated
	IsDirect  bool
	IsReverse bool
}

// The corresponding interface
type IBACnetAction interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewBACnetAction(rawData *BACnetContextTagEnumerated, isDirect bool, isReverse bool) *BACnetAction {
	return &BACnetAction{RawData: rawData, IsDirect: isDirect, IsReverse: isReverse}
}

func CastBACnetAction(structType interface{}) *BACnetAction {
	castFunc := func(typ interface{}) *BACnetAction {
		if casted, ok := typ.(BACnetAction); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetAction); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetAction) GetTypeName() string {
	return "BACnetAction"
}

func (m *BACnetAction) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetAction) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (rawData)
	if m.RawData != nil {
		lengthInBits += (*m.RawData).LengthInBits()
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetAction) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetActionParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetAction, error) {
	if pullErr := readBuffer.PullContext("BACnetAction"); pullErr != nil {
		return nil, pullErr
	}

	// Optional Field (rawData) (Can be skipped, if a given expression evaluates to false)
	var rawData *BACnetContextTagEnumerated = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("rawData"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, tagNumber, BACnetDataType_ENUMERATED)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'rawData' field")
		case _err == utils.ParseAssertError:
			readBuffer.Reset(currentPos)
		default:
			rawData = CastBACnetContextTagEnumerated(_val)
			if closeErr := readBuffer.CloseContext("rawData"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Virtual field
	_isDirect := bool(bool(bool((rawData) != (nil))) && bool(bool((len((*rawData).Data)) == (1)))) && bool(bool(((*rawData).Data[0]) == (0)))
	isDirect := bool(_isDirect)

	// Virtual field
	_isReverse := bool(bool(bool((rawData) != (nil))) && bool(bool((len((*rawData).Data)) == (1)))) && bool(bool(((*rawData).Data[0]) == (1)))
	isReverse := bool(_isReverse)

	if closeErr := readBuffer.CloseContext("BACnetAction"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetAction(rawData, isDirect, isReverse), nil
}

func (m *BACnetAction) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetAction"); pushErr != nil {
		return pushErr
	}

	// Optional Field (rawData) (Can be skipped, if the value is null)
	var rawData *BACnetContextTagEnumerated = nil
	if m.RawData != nil {
		if pushErr := writeBuffer.PushContext("rawData"); pushErr != nil {
			return pushErr
		}
		rawData = m.RawData
		_rawDataErr := rawData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("rawData"); popErr != nil {
			return popErr
		}
		if _rawDataErr != nil {
			return errors.Wrap(_rawDataErr, "Error serializing 'rawData' field")
		}
	}
	// Virtual field
	if _isDirectErr := writeBuffer.WriteVirtual("isDirect", m.IsDirect); _isDirectErr != nil {
		return errors.Wrap(_isDirectErr, "Error serializing 'isDirect' field")
	}
	// Virtual field
	if _isReverseErr := writeBuffer.WriteVirtual("isReverse", m.IsReverse); _isReverseErr != nil {
		return errors.Wrap(_isReverseErr, "Error serializing 'isReverse' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAction"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetAction) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}