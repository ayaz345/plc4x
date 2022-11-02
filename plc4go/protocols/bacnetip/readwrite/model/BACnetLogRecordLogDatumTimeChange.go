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

// BACnetLogRecordLogDatumTimeChange is the corresponding interface of BACnetLogRecordLogDatumTimeChange
type BACnetLogRecordLogDatumTimeChange interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() BACnetContextTagReal
}

// BACnetLogRecordLogDatumTimeChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumTimeChange.
// This is useful for switch cases.
type BACnetLogRecordLogDatumTimeChangeExactly interface {
	BACnetLogRecordLogDatumTimeChange
	isBACnetLogRecordLogDatumTimeChange() bool
}

// _BACnetLogRecordLogDatumTimeChange is the data-structure of this message
type _BACnetLogRecordLogDatumTimeChange struct {
	*_BACnetLogRecordLogDatum
	TimeChange BACnetContextTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumTimeChange) InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumTimeChange) GetTimeChange() BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumTimeChange factory function for _BACnetLogRecordLogDatumTimeChange
func NewBACnetLogRecordLogDatumTimeChange(timeChange BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumTimeChange {
	_result := &_BACnetLogRecordLogDatumTimeChange{
		TimeChange:               timeChange,
		_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumTimeChange(structType interface{}) BACnetLogRecordLogDatumTimeChange {
	if casted, ok := structType.(BACnetLogRecordLogDatumTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumTimeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetTypeName() string {
	return "BACnetLogRecordLogDatumTimeChange"
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumTimeChangeParse(theBytes []byte, tagNumber uint8) (BACnetLogRecordLogDatumTimeChange, error) {
	return BACnetLogRecordLogDatumTimeChangeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber) // TODO: get endianness from mspec
}

func BACnetLogRecordLogDatumTimeChangeParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumTimeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumTimeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumTimeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeChange)
	if pullErr := readBuffer.PullContext("timeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeChange")
	}
	_timeChange, _timeChangeErr := BACnetContextTagParseWithBuffer(readBuffer, uint8(uint8(9)), BACnetDataType(BACnetDataType_REAL))
	if _timeChangeErr != nil {
		return nil, errors.Wrap(_timeChangeErr, "Error parsing 'timeChange' field of BACnetLogRecordLogDatumTimeChange")
	}
	timeChange := _timeChange.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("timeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeChange")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumTimeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumTimeChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumTimeChange{
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		TimeChange: timeChange,
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumTimeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumTimeChange) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumTimeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumTimeChange")
		}

		// Simple Field (timeChange)
		if pushErr := writeBuffer.PushContext("timeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeChange")
		}
		_timeChangeErr := writeBuffer.WriteSerializable(m.GetTimeChange())
		if popErr := writeBuffer.PopContext("timeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeChange")
		}
		if _timeChangeErr != nil {
			return errors.Wrap(_timeChangeErr, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumTimeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumTimeChange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumTimeChange) isBACnetLogRecordLogDatumTimeChange() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumTimeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
