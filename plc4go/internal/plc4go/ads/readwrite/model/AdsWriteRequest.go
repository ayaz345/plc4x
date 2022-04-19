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

// AdsWriteRequest is the data-structure of this message
type AdsWriteRequest struct {
	*AdsData
	IndexGroup  uint32
	IndexOffset uint32
	Data        []byte
}

// IAdsWriteRequest is the corresponding interface of AdsWriteRequest
type IAdsWriteRequest interface {
	IAdsData
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetData returns Data (property field)
	GetData() []byte
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

func (m *AdsWriteRequest) GetCommandId() CommandId {
	return CommandId_ADS_WRITE
}

func (m *AdsWriteRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *AdsWriteRequest) InitializeParent(parent *AdsData) {}

func (m *AdsWriteRequest) GetParent() *AdsData {
	return m.AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AdsWriteRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *AdsWriteRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *AdsWriteRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsWriteRequest factory function for AdsWriteRequest
func NewAdsWriteRequest(indexGroup uint32, indexOffset uint32, data []byte) *AdsWriteRequest {
	_result := &AdsWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Data:        data,
		AdsData:     NewAdsData(),
	}
	_result.Child = _result
	return _result
}

func CastAdsWriteRequest(structType interface{}) *AdsWriteRequest {
	if casted, ok := structType.(AdsWriteRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsWriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(AdsData); ok {
		return CastAdsWriteRequest(casted.Child)
	}
	if casted, ok := structType.(*AdsData); ok {
		return CastAdsWriteRequest(casted.Child)
	}
	return nil
}

func (m *AdsWriteRequest) GetTypeName() string {
	return "AdsWriteRequest"
}

func (m *AdsWriteRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsWriteRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsWriteRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsWriteRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsWriteRequest, error) {
	if pullErr := readBuffer.PullContext("AdsWriteRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (indexGroup)
	_indexGroup, _indexGroupErr := readBuffer.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field")
	}
	indexGroup := _indexGroup

	// Simple Field (indexOffset)
	_indexOffset, _indexOffsetErr := readBuffer.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field")
	}
	indexOffset := _indexOffset

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(length)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("AdsWriteRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Data:        data,
		AdsData:     &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child, nil
}

func (m *AdsWriteRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (indexGroup)
		indexGroup := uint32(m.IndexGroup)
		_indexGroupErr := writeBuffer.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.IndexOffset)
		_indexOffsetErr := writeBuffer.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.GetData())))
		_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("AdsWriteRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
