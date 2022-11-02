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

// BACnetConstructedDataControlledVariableReference is the corresponding interface of BACnetConstructedDataControlledVariableReference
type BACnetConstructedDataControlledVariableReference interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetControlledVariableReference returns ControlledVariableReference (property field)
	GetControlledVariableReference() BACnetObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectPropertyReference
}

// BACnetConstructedDataControlledVariableReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataControlledVariableReference.
// This is useful for switch cases.
type BACnetConstructedDataControlledVariableReferenceExactly interface {
	BACnetConstructedDataControlledVariableReference
	isBACnetConstructedDataControlledVariableReference() bool
}

// _BACnetConstructedDataControlledVariableReference is the data-structure of this message
type _BACnetConstructedDataControlledVariableReference struct {
	*_BACnetConstructedData
	ControlledVariableReference BACnetObjectPropertyReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataControlledVariableReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROLLED_VARIABLE_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataControlledVariableReference) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetControlledVariableReference() BACnetObjectPropertyReference {
	return m.ControlledVariableReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetActualValue() BACnetObjectPropertyReference {
	return CastBACnetObjectPropertyReference(m.GetControlledVariableReference())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataControlledVariableReference factory function for _BACnetConstructedDataControlledVariableReference
func NewBACnetConstructedDataControlledVariableReference(controlledVariableReference BACnetObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataControlledVariableReference {
	_result := &_BACnetConstructedDataControlledVariableReference{
		ControlledVariableReference: controlledVariableReference,
		_BACnetConstructedData:      NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataControlledVariableReference(structType interface{}) BACnetConstructedDataControlledVariableReference {
	if casted, ok := structType.(BACnetConstructedDataControlledVariableReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlledVariableReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataControlledVariableReference) GetTypeName() string {
	return "BACnetConstructedDataControlledVariableReference"
}

func (m *_BACnetConstructedDataControlledVariableReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataControlledVariableReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (controlledVariableReference)
	lengthInBits += m.ControlledVariableReference.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataControlledVariableReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataControlledVariableReferenceParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataControlledVariableReference, error) {
	return BACnetConstructedDataControlledVariableReferenceParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataControlledVariableReferenceParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataControlledVariableReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlledVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlledVariableReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (controlledVariableReference)
	if pullErr := readBuffer.PullContext("controlledVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for controlledVariableReference")
	}
	_controlledVariableReference, _controlledVariableReferenceErr := BACnetObjectPropertyReferenceParseWithBuffer(readBuffer)
	if _controlledVariableReferenceErr != nil {
		return nil, errors.Wrap(_controlledVariableReferenceErr, "Error parsing 'controlledVariableReference' field of BACnetConstructedDataControlledVariableReference")
	}
	controlledVariableReference := _controlledVariableReference.(BACnetObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("controlledVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for controlledVariableReference")
	}

	// Virtual field
	_actualValue := controlledVariableReference
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlledVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlledVariableReference")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataControlledVariableReference{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ControlledVariableReference: controlledVariableReference,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataControlledVariableReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataControlledVariableReference) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlledVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlledVariableReference")
		}

		// Simple Field (controlledVariableReference)
		if pushErr := writeBuffer.PushContext("controlledVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for controlledVariableReference")
		}
		_controlledVariableReferenceErr := writeBuffer.WriteSerializable(m.GetControlledVariableReference())
		if popErr := writeBuffer.PopContext("controlledVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for controlledVariableReference")
		}
		if _controlledVariableReferenceErr != nil {
			return errors.Wrap(_controlledVariableReferenceErr, "Error serializing 'controlledVariableReference' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlledVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlledVariableReference")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataControlledVariableReference) isBACnetConstructedDataControlledVariableReference() bool {
	return true
}

func (m *_BACnetConstructedDataControlledVariableReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
