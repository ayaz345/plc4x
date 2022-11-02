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

// MediaTransportControlDataPauseResume is the corresponding interface of MediaTransportControlDataPauseResume
type MediaTransportControlDataPauseResume interface {
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetOperation returns Operation (property field)
	GetOperation() byte
	// GetIsPause returns IsPause (virtual field)
	GetIsPause() bool
	// GetIsResume returns IsResume (virtual field)
	GetIsResume() bool
}

// MediaTransportControlDataPauseResumeExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataPauseResume.
// This is useful for switch cases.
type MediaTransportControlDataPauseResumeExactly interface {
	MediaTransportControlDataPauseResume
	isMediaTransportControlDataPauseResume() bool
}

// _MediaTransportControlDataPauseResume is the data-structure of this message
type _MediaTransportControlDataPauseResume struct {
	*_MediaTransportControlData
	Operation byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataPauseResume) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataPauseResume) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataPauseResume) GetOperation() byte {
	return m.Operation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataPauseResume) GetIsPause() bool {
	return bool(bool((m.GetOperation()) == (0x00)))
}

func (m *_MediaTransportControlDataPauseResume) GetIsResume() bool {
	return bool(bool((m.GetOperation()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataPauseResume factory function for _MediaTransportControlDataPauseResume
func NewMediaTransportControlDataPauseResume(operation byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataPauseResume {
	_result := &_MediaTransportControlDataPauseResume{
		Operation:                  operation,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataPauseResume(structType interface{}) MediaTransportControlDataPauseResume {
	if casted, ok := structType.(MediaTransportControlDataPauseResume); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataPauseResume); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataPauseResume) GetTypeName() string {
	return "MediaTransportControlDataPauseResume"
}

func (m *_MediaTransportControlDataPauseResume) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MediaTransportControlDataPauseResume) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (operation)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataPauseResume) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MediaTransportControlDataPauseResumeParse(theBytes []byte) (MediaTransportControlDataPauseResume, error) {
	return MediaTransportControlDataPauseResumeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian))) // TODO: get endianness from mspec
}

func MediaTransportControlDataPauseResumeParseWithBuffer(readBuffer utils.ReadBuffer) (MediaTransportControlDataPauseResume, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataPauseResume"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataPauseResume")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (operation)
	_operation, _operationErr := readBuffer.ReadByte("operation")
	if _operationErr != nil {
		return nil, errors.Wrap(_operationErr, "Error parsing 'operation' field of MediaTransportControlDataPauseResume")
	}
	operation := _operation

	// Virtual field
	_isPause := bool((operation) == (0x00))
	isPause := bool(_isPause)
	_ = isPause

	// Virtual field
	_isResume := bool((operation) > (0xFE))
	isResume := bool(_isResume)
	_ = isResume

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataPauseResume"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataPauseResume")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataPauseResume{
		_MediaTransportControlData: &_MediaTransportControlData{},
		Operation:                  operation,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataPauseResume) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataPauseResume) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataPauseResume"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataPauseResume")
		}

		// Simple Field (operation)
		operation := byte(m.GetOperation())
		_operationErr := writeBuffer.WriteByte("operation", (operation))
		if _operationErr != nil {
			return errors.Wrap(_operationErr, "Error serializing 'operation' field")
		}
		// Virtual field
		if _isPauseErr := writeBuffer.WriteVirtual("isPause", m.GetIsPause()); _isPauseErr != nil {
			return errors.Wrap(_isPauseErr, "Error serializing 'isPause' field")
		}
		// Virtual field
		if _isResumeErr := writeBuffer.WriteVirtual("isResume", m.GetIsResume()); _isResumeErr != nil {
			return errors.Wrap(_isResumeErr, "Error serializing 'isResume' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataPauseResume"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataPauseResume")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataPauseResume) isMediaTransportControlDataPauseResume() bool {
	return true
}

func (m *_MediaTransportControlDataPauseResume) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
