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

// BACnetConstructedDataBinaryInputInterfaceValue is the data-structure of this message
type BACnetConstructedDataBinaryInputInterfaceValue struct {
	*BACnetConstructedData
	InterfaceValue *BACnetOptionalBinaryPV

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataBinaryInputInterfaceValue is the corresponding interface of BACnetConstructedDataBinaryInputInterfaceValue
type IBACnetConstructedDataBinaryInputInterfaceValue interface {
	IBACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() *BACnetOptionalBinaryPV
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

func (m *BACnetConstructedDataBinaryInputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_INPUT
}

func (m *BACnetConstructedDataBinaryInputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBinaryInputInterfaceValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBinaryInputInterfaceValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBinaryInputInterfaceValue) GetInterfaceValue() *BACnetOptionalBinaryPV {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBinaryInputInterfaceValue factory function for BACnetConstructedDataBinaryInputInterfaceValue
func NewBACnetConstructedDataBinaryInputInterfaceValue(interfaceValue *BACnetOptionalBinaryPV, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataBinaryInputInterfaceValue {
	_result := &BACnetConstructedDataBinaryInputInterfaceValue{
		InterfaceValue:        interfaceValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBinaryInputInterfaceValue(structType interface{}) *BACnetConstructedDataBinaryInputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryInputInterfaceValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryInputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBinaryInputInterfaceValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBinaryInputInterfaceValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBinaryInputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryInputInterfaceValue"
}

func (m *BACnetConstructedDataBinaryInputInterfaceValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBinaryInputInterfaceValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBinaryInputInterfaceValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBinaryInputInterfaceValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataBinaryInputInterfaceValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryInputInterfaceValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (interfaceValue)
	if pullErr := readBuffer.PullContext("interfaceValue"); pullErr != nil {
		return nil, pullErr
	}
	_interfaceValue, _interfaceValueErr := BACnetOptionalBinaryPVParse(readBuffer)
	if _interfaceValueErr != nil {
		return nil, errors.Wrap(_interfaceValueErr, "Error parsing 'interfaceValue' field")
	}
	interfaceValue := CastBACnetOptionalBinaryPV(_interfaceValue)
	if closeErr := readBuffer.CloseContext("interfaceValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryInputInterfaceValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBinaryInputInterfaceValue{
		InterfaceValue:        CastBACnetOptionalBinaryPV(interfaceValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBinaryInputInterfaceValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryInputInterfaceValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (interfaceValue)
		if pushErr := writeBuffer.PushContext("interfaceValue"); pushErr != nil {
			return pushErr
		}
		_interfaceValueErr := m.InterfaceValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("interfaceValue"); popErr != nil {
			return popErr
		}
		if _interfaceValueErr != nil {
			return errors.Wrap(_interfaceValueErr, "Error serializing 'interfaceValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryInputInterfaceValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBinaryInputInterfaceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
