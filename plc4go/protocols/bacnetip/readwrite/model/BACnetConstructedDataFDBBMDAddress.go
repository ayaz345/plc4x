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

// BACnetConstructedDataFDBBMDAddress is the corresponding interface of BACnetConstructedDataFDBBMDAddress
type BACnetConstructedDataFDBBMDAddress interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFDBBMDAddress returns FDBBMDAddress (property field)
	GetFDBBMDAddress() BACnetHostNPort
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetHostNPort
}

// BACnetConstructedDataFDBBMDAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFDBBMDAddress.
// This is useful for switch cases.
type BACnetConstructedDataFDBBMDAddressExactly interface {
	BACnetConstructedDataFDBBMDAddress
	isBACnetConstructedDataFDBBMDAddress() bool
}

// _BACnetConstructedDataFDBBMDAddress is the data-structure of this message
type _BACnetConstructedDataFDBBMDAddress struct {
	*_BACnetConstructedData
	FDBBMDAddress BACnetHostNPort
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFDBBMDAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FD_BBMD_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFDBBMDAddress) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFDBBMDAddress) GetFDBBMDAddress() BACnetHostNPort {
	return m.FDBBMDAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFDBBMDAddress) GetActualValue() BACnetHostNPort {
	return CastBACnetHostNPort(m.GetFDBBMDAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFDBBMDAddress factory function for _BACnetConstructedDataFDBBMDAddress
func NewBACnetConstructedDataFDBBMDAddress(fDBBMDAddress BACnetHostNPort, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFDBBMDAddress {
	_result := &_BACnetConstructedDataFDBBMDAddress{
		FDBBMDAddress:          fDBBMDAddress,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFDBBMDAddress(structType interface{}) BACnetConstructedDataFDBBMDAddress {
	if casted, ok := structType.(BACnetConstructedDataFDBBMDAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFDBBMDAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetTypeName() string {
	return "BACnetConstructedDataFDBBMDAddress"
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fDBBMDAddress)
	lengthInBits += m.FDBBMDAddress.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFDBBMDAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFDBBMDAddressParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFDBBMDAddress, error) {
	return BACnetConstructedDataFDBBMDAddressParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataFDBBMDAddressParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFDBBMDAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFDBBMDAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFDBBMDAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fDBBMDAddress)
	if pullErr := readBuffer.PullContext("fDBBMDAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fDBBMDAddress")
	}
	_fDBBMDAddress, _fDBBMDAddressErr := BACnetHostNPortParseWithBuffer(readBuffer)
	if _fDBBMDAddressErr != nil {
		return nil, errors.Wrap(_fDBBMDAddressErr, "Error parsing 'fDBBMDAddress' field of BACnetConstructedDataFDBBMDAddress")
	}
	fDBBMDAddress := _fDBBMDAddress.(BACnetHostNPort)
	if closeErr := readBuffer.CloseContext("fDBBMDAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fDBBMDAddress")
	}

	// Virtual field
	_actualValue := fDBBMDAddress
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFDBBMDAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFDBBMDAddress")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFDBBMDAddress{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FDBBMDAddress: fDBBMDAddress,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFDBBMDAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFDBBMDAddress) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFDBBMDAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFDBBMDAddress")
		}

		// Simple Field (fDBBMDAddress)
		if pushErr := writeBuffer.PushContext("fDBBMDAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fDBBMDAddress")
		}
		_fDBBMDAddressErr := writeBuffer.WriteSerializable(m.GetFDBBMDAddress())
		if popErr := writeBuffer.PopContext("fDBBMDAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fDBBMDAddress")
		}
		if _fDBBMDAddressErr != nil {
			return errors.Wrap(_fDBBMDAddressErr, "Error serializing 'fDBBMDAddress' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFDBBMDAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFDBBMDAddress")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFDBBMDAddress) isBACnetConstructedDataFDBBMDAddress() bool {
	return true
}

func (m *_BACnetConstructedDataFDBBMDAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
