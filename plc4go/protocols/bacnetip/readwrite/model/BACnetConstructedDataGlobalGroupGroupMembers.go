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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataGlobalGroupGroupMembers is the data-structure of this message
type BACnetConstructedDataGlobalGroupGroupMembers struct {
	*BACnetConstructedData
	GroupMembers []*BACnetDeviceObjectPropertyReference

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataGlobalGroupGroupMembers is the corresponding interface of BACnetConstructedDataGlobalGroupGroupMembers
type IBACnetConstructedDataGlobalGroupGroupMembers interface {
	IBACnetConstructedData
	// GetGroupMembers returns GroupMembers (property field)
	GetGroupMembers() []*BACnetDeviceObjectPropertyReference
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

func (m *BACnetConstructedDataGlobalGroupGroupMembers) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_GLOBAL_GROUP
}

func (m *BACnetConstructedDataGlobalGroupGroupMembers) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_MEMBERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataGlobalGroupGroupMembers) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataGlobalGroupGroupMembers) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataGlobalGroupGroupMembers) GetGroupMembers() []*BACnetDeviceObjectPropertyReference {
	return m.GroupMembers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataGlobalGroupGroupMembers factory function for BACnetConstructedDataGlobalGroupGroupMembers
func NewBACnetConstructedDataGlobalGroupGroupMembers(groupMembers []*BACnetDeviceObjectPropertyReference, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataGlobalGroupGroupMembers {
	_result := &BACnetConstructedDataGlobalGroupGroupMembers{
		GroupMembers:          groupMembers,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataGlobalGroupGroupMembers(structType interface{}) *BACnetConstructedDataGlobalGroupGroupMembers {
	if casted, ok := structType.(BACnetConstructedDataGlobalGroupGroupMembers); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGlobalGroupGroupMembers); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataGlobalGroupGroupMembers(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataGlobalGroupGroupMembers(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataGlobalGroupGroupMembers) GetTypeName() string {
	return "BACnetConstructedDataGlobalGroupGroupMembers"
}

func (m *BACnetConstructedDataGlobalGroupGroupMembers) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataGlobalGroupGroupMembers) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.GroupMembers) > 0 {
		for _, element := range m.GroupMembers {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataGlobalGroupGroupMembers) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataGlobalGroupGroupMembersParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataGlobalGroupGroupMembers, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGlobalGroupGroupMembers"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (groupMembers)
	if pullErr := readBuffer.PullContext("groupMembers", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	groupMembers := make([]*BACnetDeviceObjectPropertyReference, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'groupMembers' field")
			}
			groupMembers = append(groupMembers, CastBACnetDeviceObjectPropertyReference(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("groupMembers", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGlobalGroupGroupMembers"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataGlobalGroupGroupMembers{
		GroupMembers:          groupMembers,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataGlobalGroupGroupMembers) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGlobalGroupGroupMembers"); pushErr != nil {
			return pushErr
		}

		// Array Field (groupMembers)
		if m.GroupMembers != nil {
			if pushErr := writeBuffer.PushContext("groupMembers", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.GroupMembers {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'groupMembers' field")
				}
			}
			if popErr := writeBuffer.PopContext("groupMembers", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGlobalGroupGroupMembers"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataGlobalGroupGroupMembers) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
