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

// BACnetApplicationTagCharacterString is the corresponding interface of BACnetApplicationTagCharacterString
type BACnetApplicationTagCharacterString interface {
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadCharacterString
	// GetValue returns Value (virtual field)
	GetValue() string
}

// BACnetApplicationTagCharacterStringExactly can be used when we want exactly this type and not a type which fulfills BACnetApplicationTagCharacterString.
// This is useful for switch cases.
type BACnetApplicationTagCharacterStringExactly interface {
	BACnetApplicationTagCharacterString
	isBACnetApplicationTagCharacterString() bool
}

// _BACnetApplicationTagCharacterString is the data-structure of this message
type _BACnetApplicationTagCharacterString struct {
	*_BACnetApplicationTag
	Payload BACnetTagPayloadCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagCharacterString) InitializeParent(parent BACnetApplicationTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetApplicationTagCharacterString) GetParent() BACnetApplicationTag {
	return m._BACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagCharacterString) GetPayload() BACnetTagPayloadCharacterString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagCharacterString) GetValue() string {
	return string(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagCharacterString factory function for _BACnetApplicationTagCharacterString
func NewBACnetApplicationTagCharacterString(payload BACnetTagPayloadCharacterString, header BACnetTagHeader) *_BACnetApplicationTagCharacterString {
	_result := &_BACnetApplicationTagCharacterString{
		Payload:               payload,
		_BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagCharacterString(structType interface{}) BACnetApplicationTagCharacterString {
	if casted, ok := structType.(BACnetApplicationTagCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagCharacterString) GetTypeName() string {
	return "BACnetApplicationTagCharacterString"
}

func (m *_BACnetApplicationTagCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetApplicationTagCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagCharacterStringParse(theBytes []byte, header BACnetTagHeader) (BACnetApplicationTagCharacterString, error) {
	return BACnetApplicationTagCharacterStringParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), header) // TODO: get endianness from mspec
}

func BACnetApplicationTagCharacterStringParseWithBuffer(readBuffer utils.ReadBuffer, header BACnetTagHeader) (BACnetApplicationTagCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadCharacterStringParseWithBuffer(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetApplicationTagCharacterString")
	}
	payload := _payload.(BACnetTagPayloadCharacterString)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_value := payload.GetValue()
	value := string(_value)
	_ = value

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagCharacterString")
	}

	// Create a partially initialized instance
	_child := &_BACnetApplicationTagCharacterString{
		_BACnetApplicationTag: &_BACnetApplicationTag{},
		Payload:               payload,
	}
	_child._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetApplicationTagCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagCharacterString) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagCharacterString")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _valueErr := writeBuffer.WriteVirtual("value", m.GetValue()); _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagCharacterString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagCharacterString) isBACnetApplicationTagCharacterString() bool {
	return true
}

func (m *_BACnetApplicationTagCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
