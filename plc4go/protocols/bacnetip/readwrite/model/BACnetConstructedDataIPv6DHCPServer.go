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

// BACnetConstructedDataIPv6DHCPServer is the corresponding interface of BACnetConstructedDataIPv6DHCPServer
type BACnetConstructedDataIPv6DHCPServer interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDhcpServer returns DhcpServer (property field)
	GetDhcpServer() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataIPv6DHCPServerExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPv6DHCPServer.
// This is useful for switch cases.
type BACnetConstructedDataIPv6DHCPServerExactly interface {
	BACnetConstructedDataIPv6DHCPServer
	isBACnetConstructedDataIPv6DHCPServer() bool
}

// _BACnetConstructedDataIPv6DHCPServer is the data-structure of this message
type _BACnetConstructedDataIPv6DHCPServer struct {
	*_BACnetConstructedData
	DhcpServer BACnetApplicationTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPServer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DHCP_SERVER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DHCPServer) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPServer) GetDhcpServer() BACnetApplicationTagOctetString {
	return m.DhcpServer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPServer) GetActualValue() BACnetApplicationTagOctetString {
	return CastBACnetApplicationTagOctetString(m.GetDhcpServer())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6DHCPServer factory function for _BACnetConstructedDataIPv6DHCPServer
func NewBACnetConstructedDataIPv6DHCPServer(dhcpServer BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DHCPServer {
	_result := &_BACnetConstructedDataIPv6DHCPServer{
		DhcpServer:             dhcpServer,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DHCPServer(structType interface{}) BACnetConstructedDataIPv6DHCPServer {
	if casted, ok := structType.(BACnetConstructedDataIPv6DHCPServer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DHCPServer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetTypeName() string {
	return "BACnetConstructedDataIPv6DHCPServer"
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dhcpServer)
	lengthInBits += m.DhcpServer.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPv6DHCPServerParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6DHCPServer, error) {
	return BACnetConstructedDataIPv6DHCPServerParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataIPv6DHCPServerParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6DHCPServer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DHCPServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DHCPServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dhcpServer)
	if pullErr := readBuffer.PullContext("dhcpServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dhcpServer")
	}
	_dhcpServer, _dhcpServerErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _dhcpServerErr != nil {
		return nil, errors.Wrap(_dhcpServerErr, "Error parsing 'dhcpServer' field of BACnetConstructedDataIPv6DHCPServer")
	}
	dhcpServer := _dhcpServer.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("dhcpServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dhcpServer")
	}

	// Virtual field
	_actualValue := dhcpServer
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DHCPServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DHCPServer")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPv6DHCPServer{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DhcpServer: dhcpServer,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPv6DHCPServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6DHCPServer) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DHCPServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DHCPServer")
		}

		// Simple Field (dhcpServer)
		if pushErr := writeBuffer.PushContext("dhcpServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dhcpServer")
		}
		_dhcpServerErr := writeBuffer.WriteSerializable(m.GetDhcpServer())
		if popErr := writeBuffer.PopContext("dhcpServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dhcpServer")
		}
		if _dhcpServerErr != nil {
			return errors.Wrap(_dhcpServerErr, "Error serializing 'dhcpServer' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DHCPServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DHCPServer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DHCPServer) isBACnetConstructedDataIPv6DHCPServer() bool {
	return true
}

func (m *_BACnetConstructedDataIPv6DHCPServer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
