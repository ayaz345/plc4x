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

// BACnetPropertyStatesWriteStatus is the corresponding interface of BACnetPropertyStatesWriteStatus
type BACnetPropertyStatesWriteStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetWriteStatus returns WriteStatus (property field)
	GetWriteStatus() BACnetWriteStatusTagged
}

// BACnetPropertyStatesWriteStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesWriteStatus.
// This is useful for switch cases.
type BACnetPropertyStatesWriteStatusExactly interface {
	BACnetPropertyStatesWriteStatus
	isBACnetPropertyStatesWriteStatus() bool
}

// _BACnetPropertyStatesWriteStatus is the data-structure of this message
type _BACnetPropertyStatesWriteStatus struct {
	*_BACnetPropertyStates
	WriteStatus BACnetWriteStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesWriteStatus) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesWriteStatus) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesWriteStatus) GetWriteStatus() BACnetWriteStatusTagged {
	return m.WriteStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesWriteStatus factory function for _BACnetPropertyStatesWriteStatus
func NewBACnetPropertyStatesWriteStatus(writeStatus BACnetWriteStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesWriteStatus {
	_result := &_BACnetPropertyStatesWriteStatus{
		WriteStatus:           writeStatus,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesWriteStatus(structType interface{}) BACnetPropertyStatesWriteStatus {
	if casted, ok := structType.(BACnetPropertyStatesWriteStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesWriteStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesWriteStatus) GetTypeName() string {
	return "BACnetPropertyStatesWriteStatus"
}

func (m *_BACnetPropertyStatesWriteStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesWriteStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (writeStatus)
	lengthInBits += m.WriteStatus.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesWriteStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesWriteStatusParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesWriteStatus, error) {
	return BACnetPropertyStatesWriteStatusParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), peekedTagNumber) // TODO: get endianness from mspec
}

func BACnetPropertyStatesWriteStatusParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesWriteStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesWriteStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesWriteStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (writeStatus)
	if pullErr := readBuffer.PullContext("writeStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for writeStatus")
	}
	_writeStatus, _writeStatusErr := BACnetWriteStatusTaggedParseWithBuffer(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _writeStatusErr != nil {
		return nil, errors.Wrap(_writeStatusErr, "Error parsing 'writeStatus' field of BACnetPropertyStatesWriteStatus")
	}
	writeStatus := _writeStatus.(BACnetWriteStatusTagged)
	if closeErr := readBuffer.CloseContext("writeStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for writeStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesWriteStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesWriteStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesWriteStatus{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		WriteStatus:           writeStatus,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesWriteStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesWriteStatus) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesWriteStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesWriteStatus")
		}

		// Simple Field (writeStatus)
		if pushErr := writeBuffer.PushContext("writeStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for writeStatus")
		}
		_writeStatusErr := writeBuffer.WriteSerializable(m.GetWriteStatus())
		if popErr := writeBuffer.PopContext("writeStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for writeStatus")
		}
		if _writeStatusErr != nil {
			return errors.Wrap(_writeStatusErr, "Error serializing 'writeStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesWriteStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesWriteStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesWriteStatus) isBACnetPropertyStatesWriteStatus() bool {
	return true
}

func (m *_BACnetPropertyStatesWriteStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
