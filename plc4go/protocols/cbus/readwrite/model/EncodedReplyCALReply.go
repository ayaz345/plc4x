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

// EncodedReplyCALReply is the corresponding interface of EncodedReplyCALReply
type EncodedReplyCALReply interface {
	utils.LengthAware
	utils.Serializable
	EncodedReply
	// GetCalReply returns CalReply (property field)
	GetCalReply() CALReply
}

// EncodedReplyCALReplyExactly can be used when we want exactly this type and not a type which fulfills EncodedReplyCALReply.
// This is useful for switch cases.
type EncodedReplyCALReplyExactly interface {
	EncodedReplyCALReply
	isEncodedReplyCALReply() bool
}

// _EncodedReplyCALReply is the data-structure of this message
type _EncodedReplyCALReply struct {
	*_EncodedReply
	CalReply CALReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EncodedReplyCALReply) InitializeParent(parent EncodedReply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_EncodedReplyCALReply) GetParent() EncodedReply {
	return m._EncodedReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EncodedReplyCALReply) GetCalReply() CALReply {
	return m.CalReply
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEncodedReplyCALReply factory function for _EncodedReplyCALReply
func NewEncodedReplyCALReply(calReply CALReply, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_EncodedReplyCALReply {
	_result := &_EncodedReplyCALReply{
		CalReply:      calReply,
		_EncodedReply: NewEncodedReply(peekedByte, cBusOptions, requestContext),
	}
	_result._EncodedReply._EncodedReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEncodedReplyCALReply(structType interface{}) EncodedReplyCALReply {
	if casted, ok := structType.(EncodedReplyCALReply); ok {
		return casted
	}
	if casted, ok := structType.(*EncodedReplyCALReply); ok {
		return *casted
	}
	return nil
}

func (m *_EncodedReplyCALReply) GetTypeName() string {
	return "EncodedReplyCALReply"
}

func (m *_EncodedReplyCALReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_EncodedReplyCALReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (calReply)
	lengthInBits += m.CalReply.GetLengthInBits()

	return lengthInBits
}

func (m *_EncodedReplyCALReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func EncodedReplyCALReplyParse(theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (EncodedReplyCALReply, error) {
	return EncodedReplyCALReplyParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), cBusOptions, requestContext) // TODO: get endianness from mspec
}

func EncodedReplyCALReplyParseWithBuffer(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (EncodedReplyCALReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EncodedReplyCALReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EncodedReplyCALReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (calReply)
	if pullErr := readBuffer.PullContext("calReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calReply")
	}
	_calReply, _calReplyErr := CALReplyParseWithBuffer(readBuffer, cBusOptions, requestContext)
	if _calReplyErr != nil {
		return nil, errors.Wrap(_calReplyErr, "Error parsing 'calReply' field of EncodedReplyCALReply")
	}
	calReply := _calReply.(CALReply)
	if closeErr := readBuffer.CloseContext("calReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calReply")
	}

	if closeErr := readBuffer.CloseContext("EncodedReplyCALReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EncodedReplyCALReply")
	}

	// Create a partially initialized instance
	_child := &_EncodedReplyCALReply{
		_EncodedReply: &_EncodedReply{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		CalReply: calReply,
	}
	_child._EncodedReply._EncodedReplyChildRequirements = _child
	return _child, nil
}

func (m *_EncodedReplyCALReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EncodedReplyCALReply) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EncodedReplyCALReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EncodedReplyCALReply")
		}

		// Simple Field (calReply)
		if pushErr := writeBuffer.PushContext("calReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for calReply")
		}
		_calReplyErr := writeBuffer.WriteSerializable(m.GetCalReply())
		if popErr := writeBuffer.PopContext("calReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for calReply")
		}
		if _calReplyErr != nil {
			return errors.Wrap(_calReplyErr, "Error serializing 'calReply' field")
		}

		if popErr := writeBuffer.PopContext("EncodedReplyCALReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EncodedReplyCALReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_EncodedReplyCALReply) isEncodedReplyCALReply() bool {
	return true
}

func (m *_EncodedReplyCALReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
