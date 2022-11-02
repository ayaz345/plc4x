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

// BACnetConstructedDataProtocolObjectTypesSupported is the corresponding interface of BACnetConstructedDataProtocolObjectTypesSupported
type BACnetConstructedDataProtocolObjectTypesSupported interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProtocolObjectTypesSupported returns ProtocolObjectTypesSupported (property field)
	GetProtocolObjectTypesSupported() BACnetObjectTypesSupportedTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectTypesSupportedTagged
}

// BACnetConstructedDataProtocolObjectTypesSupportedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProtocolObjectTypesSupported.
// This is useful for switch cases.
type BACnetConstructedDataProtocolObjectTypesSupportedExactly interface {
	BACnetConstructedDataProtocolObjectTypesSupported
	isBACnetConstructedDataProtocolObjectTypesSupported() bool
}

// _BACnetConstructedDataProtocolObjectTypesSupported is the data-structure of this message
type _BACnetConstructedDataProtocolObjectTypesSupported struct {
	*_BACnetConstructedData
	ProtocolObjectTypesSupported BACnetObjectTypesSupportedTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_OBJECT_TYPES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetProtocolObjectTypesSupported() BACnetObjectTypesSupportedTagged {
	return m.ProtocolObjectTypesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetActualValue() BACnetObjectTypesSupportedTagged {
	return CastBACnetObjectTypesSupportedTagged(m.GetProtocolObjectTypesSupported())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProtocolObjectTypesSupported factory function for _BACnetConstructedDataProtocolObjectTypesSupported
func NewBACnetConstructedDataProtocolObjectTypesSupported(protocolObjectTypesSupported BACnetObjectTypesSupportedTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProtocolObjectTypesSupported {
	_result := &_BACnetConstructedDataProtocolObjectTypesSupported{
		ProtocolObjectTypesSupported: protocolObjectTypesSupported,
		_BACnetConstructedData:       NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProtocolObjectTypesSupported(structType interface{}) BACnetConstructedDataProtocolObjectTypesSupported {
	if casted, ok := structType.(BACnetConstructedDataProtocolObjectTypesSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolObjectTypesSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetTypeName() string {
	return "BACnetConstructedDataProtocolObjectTypesSupported"
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (protocolObjectTypesSupported)
	lengthInBits += m.ProtocolObjectTypesSupported.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProtocolObjectTypesSupportedParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProtocolObjectTypesSupported, error) {
	return BACnetConstructedDataProtocolObjectTypesSupportedParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataProtocolObjectTypesSupportedParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProtocolObjectTypesSupported, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolObjectTypesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProtocolObjectTypesSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (protocolObjectTypesSupported)
	if pullErr := readBuffer.PullContext("protocolObjectTypesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for protocolObjectTypesSupported")
	}
	_protocolObjectTypesSupported, _protocolObjectTypesSupportedErr := BACnetObjectTypesSupportedTaggedParseWithBuffer(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _protocolObjectTypesSupportedErr != nil {
		return nil, errors.Wrap(_protocolObjectTypesSupportedErr, "Error parsing 'protocolObjectTypesSupported' field of BACnetConstructedDataProtocolObjectTypesSupported")
	}
	protocolObjectTypesSupported := _protocolObjectTypesSupported.(BACnetObjectTypesSupportedTagged)
	if closeErr := readBuffer.CloseContext("protocolObjectTypesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for protocolObjectTypesSupported")
	}

	// Virtual field
	_actualValue := protocolObjectTypesSupported
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolObjectTypesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProtocolObjectTypesSupported")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProtocolObjectTypesSupported{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProtocolObjectTypesSupported: protocolObjectTypesSupported,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolObjectTypesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProtocolObjectTypesSupported")
		}

		// Simple Field (protocolObjectTypesSupported)
		if pushErr := writeBuffer.PushContext("protocolObjectTypesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for protocolObjectTypesSupported")
		}
		_protocolObjectTypesSupportedErr := writeBuffer.WriteSerializable(m.GetProtocolObjectTypesSupported())
		if popErr := writeBuffer.PopContext("protocolObjectTypesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for protocolObjectTypesSupported")
		}
		if _protocolObjectTypesSupportedErr != nil {
			return errors.Wrap(_protocolObjectTypesSupportedErr, "Error serializing 'protocolObjectTypesSupported' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolObjectTypesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProtocolObjectTypesSupported")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) isBACnetConstructedDataProtocolObjectTypesSupported() bool {
	return true
}

func (m *_BACnetConstructedDataProtocolObjectTypesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
