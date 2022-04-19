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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestRemovedRequestKey is the data-structure of this message
type BACnetConfirmedServiceRequestRemovedRequestKey struct {
	*BACnetConfirmedServiceRequest

	// Arguments.
	Len uint16
}

// IBACnetConfirmedServiceRequestRemovedRequestKey is the corresponding interface of BACnetConfirmedServiceRequestRemovedRequestKey
type IBACnetConfirmedServiceRequestRemovedRequestKey interface {
	IBACnetConfirmedServiceRequest
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

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) GetServiceChoice() uint8 {
	return 0x19
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

// NewBACnetConfirmedServiceRequestRemovedRequestKey factory function for BACnetConfirmedServiceRequestRemovedRequestKey
func NewBACnetConfirmedServiceRequestRemovedRequestKey(len uint16) *BACnetConfirmedServiceRequestRemovedRequestKey {
	_result := &BACnetConfirmedServiceRequestRemovedRequestKey{
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(len),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestRemovedRequestKey(structType interface{}) *BACnetConfirmedServiceRequestRemovedRequestKey {
	if casted, ok := structType.(BACnetConfirmedServiceRequestRemovedRequestKey); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestRemovedRequestKey); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestRemovedRequestKey(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestRemovedRequestKey(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) GetTypeName() string {
	return "BACnetConfirmedServiceRequestRemovedRequestKey"
}

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestRemovedRequestKeyParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequestRemovedRequestKey, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestRemovedRequestKey"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestRemovedRequestKey"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestRemovedRequestKey{
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestRemovedRequestKey"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestRemovedRequestKey"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestRemovedRequestKey) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
