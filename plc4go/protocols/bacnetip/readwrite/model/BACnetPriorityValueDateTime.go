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

// BACnetPriorityValueDateTime is the corresponding interface of BACnetPriorityValueDateTime
type BACnetPriorityValueDateTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() BACnetDateTimeEnclosed
}

// BACnetPriorityValueDateTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueDateTime.
// This is useful for switch cases.
type BACnetPriorityValueDateTimeExactly interface {
	BACnetPriorityValueDateTime
	isBACnetPriorityValueDateTime() bool
}

// _BACnetPriorityValueDateTime is the data-structure of this message
type _BACnetPriorityValueDateTime struct {
	*_BACnetPriorityValue
	DateTimeValue BACnetDateTimeEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueDateTime) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueDateTime) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueDateTime) GetDateTimeValue() BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueDateTime factory function for _BACnetPriorityValueDateTime
func NewBACnetPriorityValueDateTime(dateTimeValue BACnetDateTimeEnclosed, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueDateTime {
	_result := &_BACnetPriorityValueDateTime{
		DateTimeValue:        dateTimeValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueDateTime(structType interface{}) BACnetPriorityValueDateTime {
	if casted, ok := structType.(BACnetPriorityValueDateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueDateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueDateTime) GetTypeName() string {
	return "BACnetPriorityValueDateTime"
}

func (m *_BACnetPriorityValueDateTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueDateTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPriorityValueDateTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueDateTimeParse(theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetPriorityValueDateTime, error) {
	return BACnetPriorityValueDateTimeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), objectTypeArgument) // TODO: get endianness from mspec
}

func BACnetPriorityValueDateTimeParseWithBuffer(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueDateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateTimeValue")
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeEnclosedParseWithBuffer(readBuffer, uint8(uint8(1)))
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field of BACnetPriorityValueDateTime")
	}
	dateTimeValue := _dateTimeValue.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateTimeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueDateTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueDateTime{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		DateTimeValue: dateTimeValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueDateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueDateTime) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueDateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueDateTime")
		}

		// Simple Field (dateTimeValue)
		if pushErr := writeBuffer.PushContext("dateTimeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateTimeValue")
		}
		_dateTimeValueErr := writeBuffer.WriteSerializable(m.GetDateTimeValue())
		if popErr := writeBuffer.PopContext("dateTimeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateTimeValue")
		}
		if _dateTimeValueErr != nil {
			return errors.Wrap(_dateTimeValueErr, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueDateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueDateTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueDateTime) isBACnetPriorityValueDateTime() bool {
	return true
}

func (m *_BACnetPriorityValueDateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
