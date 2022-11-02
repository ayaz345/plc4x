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

// BACnetConstructedDataMaskedAlarmValues is the corresponding interface of BACnetConstructedDataMaskedAlarmValues
type BACnetConstructedDataMaskedAlarmValues interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaskedAlarmValues returns MaskedAlarmValues (property field)
	GetMaskedAlarmValues() []BACnetDoorAlarmStateTagged
}

// BACnetConstructedDataMaskedAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaskedAlarmValues.
// This is useful for switch cases.
type BACnetConstructedDataMaskedAlarmValuesExactly interface {
	BACnetConstructedDataMaskedAlarmValues
	isBACnetConstructedDataMaskedAlarmValues() bool
}

// _BACnetConstructedDataMaskedAlarmValues is the data-structure of this message
type _BACnetConstructedDataMaskedAlarmValues struct {
	*_BACnetConstructedData
	MaskedAlarmValues []BACnetDoorAlarmStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaskedAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MASKED_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaskedAlarmValues) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaskedAlarmValues) GetMaskedAlarmValues() []BACnetDoorAlarmStateTagged {
	return m.MaskedAlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaskedAlarmValues factory function for _BACnetConstructedDataMaskedAlarmValues
func NewBACnetConstructedDataMaskedAlarmValues(maskedAlarmValues []BACnetDoorAlarmStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaskedAlarmValues {
	_result := &_BACnetConstructedDataMaskedAlarmValues{
		MaskedAlarmValues:      maskedAlarmValues,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaskedAlarmValues(structType interface{}) BACnetConstructedDataMaskedAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataMaskedAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaskedAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataMaskedAlarmValues"
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.MaskedAlarmValues) > 0 {
		for _, element := range m.MaskedAlarmValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataMaskedAlarmValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaskedAlarmValuesParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaskedAlarmValues, error) {
	return BACnetConstructedDataMaskedAlarmValuesParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataMaskedAlarmValuesParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaskedAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaskedAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaskedAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (maskedAlarmValues)
	if pullErr := readBuffer.PullContext("maskedAlarmValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maskedAlarmValues")
	}
	// Terminated array
	var maskedAlarmValues []BACnetDoorAlarmStateTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDoorAlarmStateTaggedParseWithBuffer(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'maskedAlarmValues' field of BACnetConstructedDataMaskedAlarmValues")
			}
			maskedAlarmValues = append(maskedAlarmValues, _item.(BACnetDoorAlarmStateTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("maskedAlarmValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maskedAlarmValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaskedAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaskedAlarmValues")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaskedAlarmValues{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaskedAlarmValues: maskedAlarmValues,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaskedAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaskedAlarmValues) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaskedAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaskedAlarmValues")
		}

		// Array Field (maskedAlarmValues)
		if pushErr := writeBuffer.PushContext("maskedAlarmValues", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maskedAlarmValues")
		}
		for _, _element := range m.GetMaskedAlarmValues() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'maskedAlarmValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("maskedAlarmValues", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maskedAlarmValues")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaskedAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaskedAlarmValues")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaskedAlarmValues) isBACnetConstructedDataMaskedAlarmValues() bool {
	return true
}

func (m *_BACnetConstructedDataMaskedAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
